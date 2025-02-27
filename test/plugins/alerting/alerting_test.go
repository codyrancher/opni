package alerting_test

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	alertingv1 "github.com/rancher/opni/pkg/apis/alerting/v1"
	corev1 "github.com/rancher/opni/pkg/apis/core/v1"
	managementv1 "github.com/rancher/opni/pkg/apis/management/v1"
	"github.com/rancher/opni/pkg/test"
	"github.com/rancher/opni/pkg/test/alerting"
	"github.com/rancher/opni/pkg/test/testruntime"
	"github.com/rancher/opni/plugins/alerting/apis/alertops"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	numServers             = 5
	numNotificationServers = 2
)

func init() {
	testruntime.IfIntegration(func() {
		BuildAlertingClusterIntegrationTests([]*alertops.ClusterConfiguration{
			{
				NumReplicas: 1,
				ResourceLimits: &alertops.ResourceLimitSpec{
					Cpu:     "100m",
					Memory:  "100Mi",
					Storage: "5Gi",
				},
				ClusterSettleTimeout:    "1m",
				ClusterGossipInterval:   "1m",
				ClusterPushPullInterval: "1m",
			},
			// HA Mode has inconsistent state results
			// {
			// 	NumReplicas: 3,
			// 	ResourceLimits: &alertops.ResourceLimitSpec{
			// 		Cpu:     "100m",
			// 		Memory:  "100Mi",
			// 		Storage: "5Gi",
			// 	},
			// 	ClusterSettleTimeout:    "1m",
			// 	ClusterGossipInterval:   "1m",
			// 	ClusterPushPullInterval: "1m",
			// },
		},
			func() alertops.AlertingAdminClient {
				return alertops.NewAlertingAdminClient(env.ManagementClientConn())
			},
			func() alertingv1.AlertConditionsClient {
				return env.NewAlertConditionsClient()
			},
			func() alertingv1.AlertEndpointsClient {
				return env.NewAlertEndpointsClient()
			},
			func() alertingv1.AlertNotificationsClient {
				return env.NewAlertNotificationsClient()
			},
			func() managementv1.ManagementClient {
				return env.NewManagementClient()
			},
		)
	})
}

type agentWithContext struct {
	id string
	context.Context
	context.CancelFunc
	port int
}

func BuildAlertingClusterIntegrationTests(
	clusterConfigurations []*alertops.ClusterConfiguration,
	alertingAdminConstructor func() alertops.AlertingAdminClient,
	alertingConditionsConstructor func() alertingv1.AlertConditionsClient,
	alertingEndpointsConstructor func() alertingv1.AlertEndpointsClient,
	alertingNotificationsConstructor func() alertingv1.AlertNotificationsClient,
	mgmtClientConstructor func() managementv1.ManagementClient,
) bool {
	return Describe("Alerting Cluster Integration tests", Ordered, Label("integration"), func() {
		var alertClusterClient alertops.AlertingAdminClient
		var alertEndpointsClient alertingv1.AlertEndpointsClient
		var alertConditionsClient alertingv1.AlertConditionsClient
		var alertNotificationsClient alertingv1.AlertNotificationsClient
		var mgmtClient managementv1.ManagementClient
		var numAgents int

		// contains agent id and other useful metadata and functions
		var agents []*agentWithContext
		// physical servers that receive opni alerting notifications
		var servers []*alerting.MockIntegrationWebhookServer
		// physical servers that receive all opni alerting notifications
		var notificationServers []*alerting.MockIntegrationWebhookServer
		// expected ways the conditions dispatch to endpoints
		expectedRouting := map[string][]string{}
		// maps condition ids where agents are disconnect to their webhook ids
		involvedDisconnects := map[string][]string{}
		When("Installing the Alerting Cluster", func() {
			BeforeAll(func() {
				alertClusterClient = alertingAdminConstructor()
				alertEndpointsClient = alertingEndpointsConstructor()
				alertConditionsClient = alertingConditionsConstructor()
				alertNotificationsClient = alertingNotificationsConstructor()
				mgmtClient = mgmtClientConstructor()
				numAgents = 5
			})
			for _, clusterConf := range clusterConfigurations {
				It("should install the alerting cluster", func() {
					_, err := alertClusterClient.InstallCluster(env.Context(), &emptypb.Empty{})
					Expect(err).To(BeNil())

					Eventually(func() error {
						status, err := alertClusterClient.GetClusterStatus(env.Context(), &emptypb.Empty{})
						if err != nil {
							return err
						}
						if status.State != alertops.InstallState_Installed {
							return fmt.Errorf("alerting cluster install state is %s", status.State.String())
						}
						return nil
					}, time.Second*30, time.Second*5).Should(Succeed())
				})

				It("should apply the configuration configuration", func() {
					_, err := alertClusterClient.ConfigureCluster(env.Context(), clusterConf)
					if err != nil {
						if s, ok := status.FromError(err); ok { // conflict is ok if using default config
							Expect(s.Code()).To(Equal(codes.FailedPrecondition))
						}
					}
					Expect(err).To(BeNil())
					Eventually(func() error {
						status, err := alertClusterClient.GetClusterStatus(env.Context(), &emptypb.Empty{})
						if err != nil {
							return err
						}
						if status.State != alertops.InstallState_Installed {
							return fmt.Errorf("alerting cluster install state is %s", status.State.String())
						}
						return nil
					}, time.Second*30, time.Second*5).Should(Succeed())

					Eventually(func() error {
						getConf, err := alertClusterClient.GetClusterConfiguration(env.Context(), &emptypb.Empty{})
						if !proto.Equal(getConf, clusterConf) {
							return fmt.Errorf("cluster config not equal : not applied")
						}
						return err
					}, time.Minute*30, time.Second*5)
				})

				It("should be able to create some endpoints", func() {
					servers = alerting.CreateWebhookServer(env, numServers)

					for _, server := range servers {
						ref, err := alertEndpointsClient.CreateAlertEndpoint(env.Context(), server.Endpoint())
						Expect(err).To(Succeed())
						server.EndpointId = ref.Id
					}
					endpList, err := alertEndpointsClient.ListAlertEndpoints(env.Context(), &alertingv1.ListAlertEndpointsRequest{})
					Expect(err).To(Succeed())
					Expect(endpList.Items).To(HaveLen(numServers))
				})

				It("should create some default conditions when bootstrapping agents", func() {
					By("expecting to have no initial conditions")
					condList, err := alertConditionsClient.ListAlertConditions(env.Context(), &alertingv1.ListAlertConditionRequest{})
					Expect(err).To(Succeed())
					Expect(condList.Items).To(HaveLen(0))

					By(fmt.Sprintf("bootstrapping %d agents", numAgents))
					certsInfo, err := mgmtClient.CertsInfo(context.Background(), &emptypb.Empty{})
					Expect(err).NotTo(HaveOccurred())
					fingerprint := certsInfo.Chain[len(certsInfo.Chain)-1].Fingerprint
					Expect(fingerprint).NotTo(BeEmpty())

					token, err := mgmtClient.CreateBootstrapToken(context.Background(), &managementv1.CreateBootstrapTokenRequest{
						Ttl: durationpb.New(1 * time.Hour),
					})
					Expect(err).NotTo(HaveOccurred())
					agentIdFunc := func(i int) string {
						return fmt.Sprintf("agent-%d-%s", i, uuid.New().String())
					}
					agents = []*agentWithContext{}
					for i := 0; i < numAgents; i++ {
						ctxCa, ca := context.WithCancel(env.Context())
						id := agentIdFunc(i)
						port, errC := env.StartAgent(id, token, []string{fingerprint}, test.WithContext(ctxCa))
						Eventually(errC).Should(Receive(BeNil()))
						agents = append(agents, &agentWithContext{
							port:       port,
							CancelFunc: ca,
							Context:    ctxCa,
							id:         id,
						})
					}
					By("verifying that there are default conditions")
					Eventually(func() error {
						condList, err := alertConditionsClient.ListAlertConditions(env.Context(), &alertingv1.ListAlertConditionRequest{})
						if err != nil {
							return err
						}
						if len(condList.Items) != numAgents*2 {
							return fmt.Errorf("expected %d conditions, got %d", numAgents*2, len(condList.Items))
						}
						return nil
					}, time.Second*30, time.Second*5).Should(Succeed())
				})

				It("shoud list conditions by given filters", func() {
					for _, agent := range agents {
						filteredByCluster, err := alertConditionsClient.ListAlertConditions(env.Context(), &alertingv1.ListAlertConditionRequest{
							Clusters: []string{agent.id},
						})
						Expect(err).To(Succeed())
						Expect(filteredByCluster.Items).To(HaveLen(2))
					}

					By("verifying all the agents conditions are critical")

					filterList, err := alertConditionsClient.ListAlertConditions(env.Context(), &alertingv1.ListAlertConditionRequest{
						Severities: []alertingv1.OpniSeverity{
							alertingv1.OpniSeverity_Warning,
							alertingv1.OpniSeverity_Error,
							alertingv1.OpniSeverity_Info,
						},
					})
					Expect(err).To(Succeed())
					Expect(filterList.Items).To(HaveLen(0))

					By("verifying we have an equal number of disconnect and capability unhealthy")

					disconnectList, err := alertConditionsClient.ListAlertConditions(env.Context(), &alertingv1.ListAlertConditionRequest{
						AlertTypes: []alertingv1.AlertType{
							alertingv1.AlertType_System,
						},
					})
					Expect(err).To(Succeed())

					capabilityList, err := alertConditionsClient.ListAlertConditions(env.Context(), &alertingv1.ListAlertConditionRequest{
						AlertTypes: []alertingv1.AlertType{
							alertingv1.AlertType_DownstreamCapability,
						},
					})
					Expect(err).To(Succeed())
					Expect(capabilityList.Items).To(HaveLen(len(disconnectList.Items)))
				})

				It("should be able to attach endpoints to conditions", func() {
					By("attaching a sample of random endpoints to default agent conditions")
					condList, err := alertConditionsClient.ListAlertConditions(env.Context(), &alertingv1.ListAlertConditionRequest{})
					Expect(err).To(Succeed())
					endpList, err := alertEndpointsClient.ListAlertEndpoints(env.Context(), &alertingv1.ListAlertEndpointsRequest{})
					for _, cond := range condList.Items {
						Expect(err).To(Succeed())
						endps := lo.Map(
							lo.Samples(endpList.Items, 1+rand.Intn(len(endpList.Items)-1)),
							func(a *alertingv1.AlertEndpointWithId, _ int) *alertingv1.AttachedEndpoint {
								return &alertingv1.AttachedEndpoint{
									EndpointId: a.GetId().Id,
								}
							})
						if cond.GetAlertCondition().GetAlertType().GetSystem() != nil {
							expectedRouting[cond.GetId().Id] = lo.Map(endps, func(a *alertingv1.AttachedEndpoint, _ int) string {
								return a.EndpointId
							})
							cond.AlertCondition.AttachedEndpoints = &alertingv1.AttachedEndpoints{
								Items:              endps,
								InitialDelay:       durationpb.New(time.Second * 1),
								ThrottlingDuration: durationpb.New(time.Second * 1),
								Details: &alertingv1.EndpointImplementation{
									Title: "disconnected agent",
									Body:  "agent %s is disconnected",
								},
							}
							cond.AlertCondition.AlertType.GetSystem().Timeout = durationpb.New(time.Second * 30)
							_, err = alertConditionsClient.UpdateAlertCondition(env.Context(), &alertingv1.UpdateAlertConditionRequest{
								Id:          cond.GetId(),
								UpdateAlert: cond.AlertCondition,
							})
							Expect(err).To(Succeed())
						}
					}

					By("creating some default webhook servers as endpoints")
					notificationServers = alerting.CreateWebhookServer(env, numNotificationServers)
					for _, server := range notificationServers {
						ref, err := alertEndpointsClient.CreateAlertEndpoint(env.Context(), server.Endpoint())
						Expect(err).To(Succeed())
						server.EndpointId = ref.Id
					}
					endpList, err = alertEndpointsClient.ListAlertEndpoints(env.Context(), &alertingv1.ListAlertEndpointsRequest{})
					Expect(err).To(Succeed())
					Expect(endpList.Items).To(HaveLen(numNotificationServers + numServers))

					By("setting the default servers as default endpoints")
					for _, server := range notificationServers {
						_, err = alertEndpointsClient.ToggleNotifications(env.Context(), &alertingv1.ToggleRequest{
							Id: &corev1.Reference{Id: server.EndpointId},
						})
						Expect(err).To(Succeed())
					}

					By("expecting the conditions to eventually move to the 'OK' state")
					Eventually(func() error {
						for _, cond := range condList.Items {
							status, err := alertConditionsClient.AlertConditionStatus(env.Context(), cond.Id)
							if err != nil {
								return err
							}
							if status.State != alertingv1.AlertConditionState_Ok {
								return fmt.Errorf("condition %s is not OK, instead in state %s, %s", cond.AlertCondition.Name, status.State.String(), status.Reason)
							}
						}
						return nil
					}, time.Second*90, time.Second*20).Should(Succeed())

					By("verifying the routing relationships are correctly loaded")
					relationships, err := alertNotificationsClient.ListRoutingRelationships(env.Context(), &emptypb.Empty{})
					Expect(err).To(Succeed())
					Expect(len(relationships.RoutingRelationships)).To(Equal(len(expectedRouting)))
					for conditionId, rel := range relationships.RoutingRelationships {
						Expect(lo.Map(rel.Items, func(c *corev1.Reference, _ int) string {
							return c.Id
						})).To(ConsistOf(expectedRouting[conditionId]))
					}
				})

				Specify("agent disconnect alarms should fire when agents are disconnected ", func() {

					// Disconnect a random 3 agents, and verify the servers have the messages
					By("disconnecting a random 3 agents")
					disconnectedIds := []string{}
					toDisconnect := lo.Samples(agents, 3)
					for _, disc := range toDisconnect {
						disc.CancelFunc()
						disconnectedIds = append(disconnectedIds, disc.id)
					}

					condList, err := alertConditionsClient.ListAlertConditions(env.Context(), &alertingv1.ListAlertConditionRequest{})
					Expect(err).To(Succeed())
					notInvolvedDisconnects := map[string]struct{}{}
					for _, cond := range condList.Items {
						if cond.GetAlertCondition().GetAlertType().GetSystem() != nil {
							if slices.Contains(disconnectedIds, cond.GetAlertCondition().GetAlertType().GetSystem().ClusterId.Id) {
								endps := lo.Map(cond.AlertCondition.AttachedEndpoints.Items,
									func(a *alertingv1.AttachedEndpoint, _ int) string {
										return a.EndpointId
									})
								involvedDisconnects[cond.GetAlertCondition().Id] = endps
							} else {
								notInvolvedDisconnects[cond.GetAlertCondition().Id] = struct{}{}
							}
						}
					}
					webhooks := lo.Uniq(lo.Flatten(lo.Values(involvedDisconnects)))
					Expect(len(webhooks)).To(BeNumerically(">", 0))

					By("verifying the agents are actually disconnected")
					Eventually(func() error {
						clusters, err := mgmtClient.ListClusters(env.Context(), &managementv1.ListClustersRequest{})
						if err != nil {
							return err
						}
						for _, cl := range clusters.Items {
							if slices.Contains(disconnectedIds, cl.GetId()) {
								healthStatus, err := mgmtClient.GetClusterHealthStatus(env.Context(), cl.Reference())
								if err != nil {
									return err
								}
								if !healthStatus.Status.Connected == false {
									return fmt.Errorf("expected disconnected health status for cluster %s: %s", cl.GetId(), healthStatus.Status.String())
								}
							}

						}
						return nil
					}, 90*time.Second, 5*time.Second).Should(Succeed())

					By("verifying the physical servers have received the disconnect messages")
					Eventually(func() error {
						servers := servers
						conditionIds := lo.Keys(involvedDisconnects)
						for _, id := range conditionIds {
							status, err := alertConditionsClient.AlertConditionStatus(env.Context(), &corev1.Reference{Id: id})
							if err != nil {
								return err
							}
							if status.GetState() != alertingv1.AlertConditionState_Firing {
								return fmt.Errorf("expected alerting condition %s to be firing, got %s", id, status.GetState().String())
							}
						}

						for id := range notInvolvedDisconnects {
							status, err := alertConditionsClient.AlertConditionStatus(env.Context(), &corev1.Reference{Id: id})
							if err != nil {
								return err
							}
							if status.GetState() != alertingv1.AlertConditionState_Ok {
								return fmt.Errorf("expected unaffected alerting condition %s to be ok, got %s", id, status.GetState().String())
							}
						}

						for _, server := range servers {
							if slices.Contains(webhooks, server.EndpointId) {
								// hard to map these excatly without recreating the internal routing logic from the routers
								// since we have dedicated routing integration tests, we can just check that the buffer is not empty
								if len(server.GetBuffer()) == 0 {
									return fmt.Errorf("expected webhook server %s to have messages, got %d", server.EndpointId, len(server.Buffer))
								}
							}
						}
						return nil
					}, time.Minute*2, time.Second*15).Should(Succeed())

					By("verifying the notification servers have not received any alarm disconnect messages")
					Eventually(func() error {
						for _, server := range notificationServers {
							if len(server.GetBuffer()) != 0 {
								return fmt.Errorf("expected webhook server %s to not have any notifications, got %d", server.EndpointId, len(server.Buffer))
							}
						}
						return nil
					}, time.Second*30, time.Second*7)
				})

				It("should be able to batch list status and filter by status", func() {
					condList, err := alertConditionsClient.ListAlertConditions(env.Context(), &alertingv1.ListAlertConditionRequest{})
					Expect(err).To(Succeed())

					statusCondList, err := alertConditionsClient.ListAlertConditionsWithStatus(env.Context(), &alertingv1.ListStatusRequest{})
					Expect(err).To(Succeed())
					Expect(statusCondList.AlertConditions).To(HaveLen(len(condList.Items)))
					for condId, cond := range statusCondList.AlertConditions {
						if slices.Contains(lo.Keys(involvedDisconnects), condId) {
							Expect(cond.Status.State).To(Equal(alertingv1.AlertConditionState_Firing))
						} else {
							Expect(cond.Status.State).To(Equal(alertingv1.AlertConditionState_Ok))
						}
					}
					firingOnlyStatusList, err := alertConditionsClient.ListAlertConditionsWithStatus(env.Context(), &alertingv1.ListStatusRequest{
						States: []alertingv1.AlertConditionState{
							alertingv1.AlertConditionState_Firing,
						},
					})
					Expect(err).To(Succeed())
					Expect(firingOnlyStatusList.AlertConditions).To(HaveLen(len(involvedDisconnects)))
					for _, cond := range firingOnlyStatusList.AlertConditions {
						Expect(cond.Status.State).To(Equal(alertingv1.AlertConditionState_Firing))
					}
				})

				It("should be able to push notifications to our notification endpoints", func() {
					Expect(len(notificationServers)).To(BeNumerically(">", 0))
					By("forwarding the message to AlertManager")
					_, err := alertNotificationsClient.PushNotification(env.Context(), &alertingv1.Notification{
						Title: "hello",
						Body:  "world",
						// set to critical in order to expedite the notification during testing
						Properties: map[string]string{
							alertingv1.NotificationPropertySeverity: alertingv1.OpniSeverity_Critical.String(),
						},
					})
					Expect(err).To(Succeed())

					By("verifying the endpoints have received the notification messages")
					Eventually(func() error {
						for _, server := range notificationServers {
							if len(server.GetBuffer()) == 0 {
								return fmt.Errorf("expected webhook server %s to have messages, got %d", server.EndpointId, len(server.Buffer))
							}
						}
						return nil
					}, time.Minute*2, time.Second*30)
				})

				It("should be able to list opni messages", func() {
					Eventually(func() error {
						list, err := alertNotificationsClient.ListNotifications(env.Context(), &alertingv1.ListNotificationRequest{})
						if err != nil {
							return err
						}
						if len(list.Items) == 0 {
							return fmt.Errorf("expected to find at least one notification, got 0")
						}
						return nil
					}, time.Minute*2, time.Second*15).Should(BeNil())

					By("verifying we enforce limits")
					list, err := alertNotificationsClient.ListNotifications(env.Context(), &alertingv1.ListNotificationRequest{
						Limit: lo.ToPtr(int32(1)),
					})
					Expect(err).To(Succeed())
					Expect(len(list.Items)).To(Equal(1))
				})

				It("should return warnings when trying to edit/delete alert endpoints that are involved in conditions", func() {
					webhooks := lo.Uniq(lo.Flatten(lo.Values(involvedDisconnects)))
					Expect(len(webhooks)).To(BeNumerically(">", 0))

					for _, webhook := range webhooks {
						involvedConditions, err := alertEndpointsClient.UpdateAlertEndpoint(env.Context(), &alertingv1.UpdateAlertEndpointRequest{
							Id: &corev1.Reference{
								Id: webhook,
							},
							UpdateAlert: &alertingv1.AlertEndpoint{
								Name:        "update",
								Description: "update",
								Endpoint: &alertingv1.AlertEndpoint_Webhook{
									Webhook: &alertingv1.WebhookEndpoint{
										Url: "http://example.com",
									},
								},
								Id: "id",
							},
							ForceUpdate: false,
						})
						Expect(err).NotTo(HaveOccurred())
						Expect(involvedConditions.Items).NotTo(HaveLen(0))
						involvedConditions, err = alertEndpointsClient.DeleteAlertEndpoint(env.Context(), &alertingv1.DeleteAlertEndpointRequest{
							Id: &corev1.Reference{
								Id: webhook,
							},
							ForceDelete: false,
						})
						Expect(err).NotTo(HaveOccurred())
						Expect(involvedConditions.Items).NotTo(HaveLen(0))
					}
				})

				It("should have a functional timeline", func() {
					condList, err := alertConditionsClient.ListAlertConditions(env.Context(), &alertingv1.ListAlertConditionRequest{})
					Expect(err).To(Succeed())

					// By("verifying the timeline shows only the firing conditions ")
					Eventually(func() error {
						timeline, err := alertConditionsClient.Timeline(env.Context(), &alertingv1.TimelineRequest{
							LookbackWindow: durationpb.New(time.Minute * 5),
						})
						if err != nil {
							return err
						}

						By("verifying the timeline matches the conditions")
						if len(timeline.Items) != len(condList.Items) {
							return fmt.Errorf("expected timeline to have %d items, got %d", len(condList.Items), len(timeline.Items))
						}

						for id, item := range timeline.GetItems() {
							if slices.Contains(lo.Keys(involvedDisconnects), id) {
								if len(item.Windows) == 0 {
									return fmt.Errorf("firing condition should show up on timeline, but does not")
								}
								if len(item.Windows) != 1 {
									return fmt.Errorf("condition evaluation is flaky, should only have one window, but has %d", len(item.Windows))
								}
								messages, err := alertNotificationsClient.ListAlarmMessages(env.Context(), &alertingv1.ListAlarmMessageRequest{
									ConditionId: id,
									Start:       item.Windows[0].Start,
									End:         timestamppb.Now(),
								})
								if err != nil {
									return err
								}
								Expect(len(messages.Items)).To(BeNumerically(">", 0))

							} else {
								if len(item.Windows) != 0 {
									return fmt.Errorf("conditions that have not fired should not show up on timeline, but do")
								}
							}
						}
						return nil
					}, time.Minute*1, time.Second*15)
				})

				It("should force update/delete alert endpoints involved in conditions", func() {
					By("verifying we can edit Alert Endpoints in use by Alert Conditions")
					endpList, err := alertEndpointsClient.ListAlertEndpoints(env.Context(), &alertingv1.ListAlertEndpointsRequest{})
					Expect(err).NotTo(HaveOccurred())
					Expect(len(endpList.Items)).To(BeNumerically(">", 0))
					for _, endp := range endpList.Items {
						_, err := alertEndpointsClient.UpdateAlertEndpoint(env.Context(), &alertingv1.UpdateAlertEndpointRequest{
							Id: &corev1.Reference{
								Id: endp.Id.Id,
							},
							UpdateAlert: &alertingv1.AlertEndpoint{
								Name:        "update",
								Description: "update",
								Endpoint: &alertingv1.AlertEndpoint_Webhook{
									Webhook: &alertingv1.WebhookEndpoint{
										Url: "http://example.com",
									},
								},
								Id: "id",
							},
							ForceUpdate: true,
						})
						Expect(err).NotTo(HaveOccurred())
					}
					endpList, err = alertEndpointsClient.ListAlertEndpoints(env.Context(), &alertingv1.ListAlertEndpointsRequest{})
					Expect(err).NotTo(HaveOccurred())
					Expect(endpList.Items).To(HaveLen(numServers + numNotificationServers))
					updatedList := lo.Filter(endpList.Items, func(item *alertingv1.AlertEndpointWithId, _ int) bool {
						if item.Endpoint.GetWebhook() != nil {
							return item.Endpoint.GetWebhook().Url == "http://example.com" && item.Endpoint.GetName() == "update" && item.Endpoint.GetDescription() == "update"
						}
						return false
					})
					Expect(updatedList).To(HaveLen(len(endpList.Items)))

					By("verifying we can delete Alert Endpoint in use by Alert Conditions")
					for _, endp := range endpList.Items {
						_, err := alertEndpointsClient.DeleteAlertEndpoint(env.Context(), &alertingv1.DeleteAlertEndpointRequest{
							Id: &corev1.Reference{
								Id: endp.Id.Id,
							},
							ForceDelete: true,
						})
						Expect(err).NotTo(HaveOccurred())
					}
					endpList, err = alertEndpointsClient.ListAlertEndpoints(env.Context(), &alertingv1.ListAlertEndpointsRequest{})
					Expect(err).NotTo(HaveOccurred())
					Expect(endpList.Items).To(HaveLen(0))

					condList, err := alertConditionsClient.ListAlertConditions(env.Context(), &alertingv1.ListAlertConditionRequest{})
					Expect(err).NotTo(HaveOccurred())
					Expect(condList.Items).NotTo(HaveLen(0))
					hasEndpoints := lo.Filter(condList.Items, func(item *alertingv1.AlertConditionWithId, _ int) bool {
						if item.AlertCondition.AttachedEndpoints != nil {
							return len(item.AlertCondition.AttachedEndpoints.Items) != 0
						}
						return false
					})
					Expect(hasEndpoints).To(HaveLen(0))
				})

				It("should delete the downstream agents", func() {
					client := env.NewManagementClient()
					agents, err := client.ListClusters(env.Context(), &managementv1.ListClustersRequest{})
					Expect(err).NotTo(HaveOccurred())
					for _, agent := range agents.Items {
						_, err := client.DeleteCluster(env.Context(), agent.Reference())
						Expect(err).NotTo(HaveOccurred())
					}
				})

				It("should uninstall the alerting cluster", func() {
					_, err := alertClusterClient.UninstallCluster(env.Context(), &alertops.UninstallRequest{
						DeleteData: true,
					})
					Expect(err).To(BeNil())
				})
			}
		})
	})
}
