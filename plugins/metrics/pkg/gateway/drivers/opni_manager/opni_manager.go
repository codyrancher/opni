package drivers

import (
	"context"
	"fmt"
	"os"
	"strconv"

	opnicorev1beta1 "github.com/rancher/opni/apis/core/v1beta1"
	corev1 "github.com/rancher/opni/pkg/apis/core/v1"
	storagev1 "github.com/rancher/opni/pkg/apis/storage/v1"
	"github.com/rancher/opni/pkg/plugins/driverutil"
	"github.com/rancher/opni/pkg/storage"
	"github.com/rancher/opni/pkg/storage/kvutil"
	"github.com/rancher/opni/pkg/util/flagutil"
	"github.com/rancher/opni/pkg/util/k8sutil"
	"github.com/rancher/opni/plugins/metrics/apis/cortexops"
	"github.com/rancher/opni/plugins/metrics/pkg/cortex/configutil"
	"github.com/rancher/opni/plugins/metrics/pkg/gateway/drivers"
	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

type OpniManagerClusterDriverOptions struct {
	K8sClient          client.Client                                               `option:"k8sClient"`
	MonitoringCluster  types.NamespacedName                                        `option:"monitoringCluster"`
	GatewayRef         types.NamespacedName                                        `option:"gatewayRef"`
	DefaultConfigStore storage.ValueStoreT[*cortexops.CapabilityBackendConfigSpec] `option:"defaultConfigStore"`
}

func (k OpniManagerClusterDriverOptions) newMonitoringCluster() *opnicorev1beta1.MonitoringCluster {
	return &opnicorev1beta1.MonitoringCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      k.MonitoringCluster.Name,
			Namespace: k.MonitoringCluster.Namespace,
		},
	}
}

func (k OpniManagerClusterDriverOptions) newGateway() *opnicorev1beta1.Gateway {
	return &opnicorev1beta1.Gateway{
		ObjectMeta: metav1.ObjectMeta{
			Name:      k.GatewayRef.Name,
			Namespace: k.GatewayRef.Namespace,
		},
	}
}

type OpniManager struct {
	cortexops.UnsafeCortexOpsServer
	driverutil.DefaultConfigurableServer[*cortexops.CapabilityBackendConfigSpec, *cortexops.GetRequest]
	OpniManagerClusterDriverOptions

	configTracker *driverutil.DefaultingConfigTracker[*cortexops.CapabilityBackendConfigSpec]
}

func NewOpniManagerClusterDriver(options OpniManagerClusterDriverOptions) (*OpniManager, error) {
	if options.K8sClient == nil {
		s := scheme.Scheme
		opnicorev1beta1.AddToScheme(s)
		c, err := k8sutil.NewK8sClient(k8sutil.ClientOptions{
			Scheme: s,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create kubernetes client: %w", err)
		}
		options.K8sClient = c
	}
	if options.DefaultConfigStore == nil {
		return nil, fmt.Errorf("missing required option: DefaultConfigStore")
	}
	activeStore := kvutil.ValueStoreAdapter[*cortexops.CapabilityBackendConfigSpec]{
		PutFunc: func(ctx context.Context, value *cortexops.CapabilityBackendConfigSpec, opts ...storage.PutOpt) error {
			putOptions := storage.PutOptions{}
			putOptions.Apply(opts...)

			cluster := options.newMonitoringCluster()
			err := options.K8sClient.Get(ctx, options.MonitoringCluster, cluster)
			exists := true
			if err != nil {
				if k8serrors.IsNotFound(err) {
					exists = false
				} else {
					return fmt.Errorf("failed to get monitoring cluster: %w", err)
				}
			}

			// look up the gateway so we can set it as an owner reference
			gateway := &opnicorev1beta1.Gateway{}
			err = options.K8sClient.Get(ctx, options.GatewayRef, gateway)
			if err != nil {
				return fmt.Errorf("failed to get gateway: %w", err)
			}

			if !exists {
				cluster.Spec.Gateway = v1.LocalObjectReference{
					Name: gateway.Name,
				}
				controllerutil.SetControllerReference(gateway, cluster, options.K8sClient.Scheme())
			}

			cluster.Spec.Cortex = opnicorev1beta1.CortexSpec{
				Enabled:         value.Enabled,
				CortexWorkloads: value.CortexWorkloads,
				CortexConfig:    value.CortexConfig,
			}
			cluster.Spec.Grafana = opnicorev1beta1.GrafanaSpec{
				GrafanaConfig: value.Grafana,
			}
			if !exists {
				err := options.K8sClient.Create(ctx, cluster)
				if err != nil {
					return fmt.Errorf("failed to create monitoring cluster: %w", err)
				}
			} else {
				err := options.K8sClient.Update(ctx, cluster)
				if err != nil {
					return fmt.Errorf("failed to update monitoring cluster: %w", err)
				}
			}
			if putOptions.RevisionOut != nil {
				*putOptions.RevisionOut, _ = strconv.ParseInt(cluster.ResourceVersion, 10, 64)
			}

			return nil
		},
		GetFunc: func(ctx context.Context, opts ...storage.GetOpt) (*cortexops.CapabilityBackendConfigSpec, error) {
			mc := options.newMonitoringCluster()
			err := options.K8sClient.Get(ctx, options.MonitoringCluster, mc)
			if err != nil {
				if k8serrors.IsNotFound(err) {
					return nil, storage.ErrNotFound
				}
				return nil, fmt.Errorf("failed to get monitoring cluster: %w", err)
			}
			return &cortexops.CapabilityBackendConfigSpec{
				Enabled:         mc.Spec.Cortex.Enabled,
				CortexWorkloads: mc.Spec.Cortex.CortexWorkloads,
				CortexConfig:    mc.Spec.Cortex.CortexConfig,
				Grafana:         mc.Spec.Grafana.GrafanaConfig,
			}, nil
		},
		DeleteFunc: func(ctx context.Context, opts ...storage.DeleteOpt) error {
			return options.K8sClient.Delete(ctx, options.newMonitoringCluster())
		},
	}

	configTracker := driverutil.NewDefaultingConfigTracker[*cortexops.CapabilityBackendConfigSpec](
		options.DefaultConfigStore, activeStore, flagutil.LoadDefaults[*cortexops.CapabilityBackendConfigSpec])

	return &OpniManager{
		OpniManagerClusterDriverOptions: options,
		configTracker:                   configTracker,
		DefaultConfigurableServer:       driverutil.NewDefaultConfigurableServer[cortexops.CortexOpsServer](configTracker),
	}, nil
}

// ListPresets implements cortexops.CortexOpsServer.
func (k *OpniManager) ListPresets(context.Context, *emptypb.Empty) (*cortexops.PresetList, error) {
	return &cortexops.PresetList{
		Items: []*cortexops.Preset{
			{
				Id: &corev1.Reference{Id: "all-in-one"},
				Metadata: &cortexops.PresetMetadata{
					DisplayName: "All In One",
					Description: "Minimal Cortex deployment with all components running in a single process",
					Notes: []string{
						"Warning: this configuration is not recommended for production use.",
					},
				},
				Spec: &cortexops.CapabilityBackendConfigSpec{
					CortexWorkloads: &cortexops.CortexWorkloadsConfig{
						Targets: map[string]*cortexops.CortexWorkloadSpec{
							"all": {Replicas: lo.ToPtr[int32](1)},
						},
					},
					CortexConfig: &cortexops.CortexApplicationConfig{
						LogLevel: lo.ToPtr("debug"),
						Storage: &storagev1.Config{
							Backend: lo.ToPtr(storagev1.Filesystem),
							Filesystem: &storagev1.FilesystemConfig{
								Dir: lo.ToPtr("/data"),
							},
						},
					},
				},
			},
			{
				Id: &corev1.Reference{Id: "highly-available"},
				Metadata: &cortexops.PresetMetadata{
					DisplayName: "Highly Available",
					Description: "Basic HA Cortex deployment with all components running in separate processes",
					Notes: []string{
						"Additional storage configuration is required. Note that filesystem storage cannot be used in HA mode.",
						"Not all components are scaled to multiple replicas by default. The replica count for each component can be modified at any time.",
					},
				},
				Spec: &cortexops.CapabilityBackendConfigSpec{
					CortexWorkloads: &cortexops.CortexWorkloadsConfig{
						Targets: map[string]*cortexops.CortexWorkloadSpec{
							"distributor":    {Replicas: lo.ToPtr[int32](1)},
							"query-frontend": {Replicas: lo.ToPtr[int32](1)},
							"purger":         {Replicas: lo.ToPtr[int32](1)},
							"ruler":          {Replicas: lo.ToPtr[int32](3)},
							"compactor":      {Replicas: lo.ToPtr[int32](3)},
							"store-gateway":  {Replicas: lo.ToPtr[int32](3)},
							"ingester":       {Replicas: lo.ToPtr[int32](3)},
							"alertmanager":   {Replicas: lo.ToPtr[int32](3)},
							"querier":        {Replicas: lo.ToPtr[int32](3)},
						},
					},
					CortexConfig: &cortexops.CortexApplicationConfig{
						LogLevel: lo.ToPtr("debug"),
						Storage: &storagev1.Config{
							Backend: lo.ToPtr(storagev1.S3),
						},
					},
				},
			},
		},
	}, nil
}

// Status implements cortexops.CortexOpsServer.
func (k *OpniManager) Status(ctx context.Context, _ *emptypb.Empty) (*cortexops.InstallStatus, error) {
	status := &cortexops.InstallStatus{
		ConfigState:  cortexops.ConfigurationState_NotConfigured,
		InstallState: cortexops.InstallState_NotInstalled,
		AppState:     cortexops.ApplicationState_NotRunning,
		Metadata: map[string]string{
			"driver": "opni-manager",
		},
	}

	cluster := k.newMonitoringCluster()
	err := k.K8sClient.Get(ctx, k.MonitoringCluster, cluster)
	if err != nil {
		if !k8serrors.IsNotFound(err) {
			return nil, fmt.Errorf("failed to get monitoring cluster: %w", err)
		}
	} else {
		status.ConfigState = cortexops.ConfigurationState_Configured
		if cluster.Spec.Cortex.Enabled != nil && *cluster.Spec.Cortex.Enabled {
			status.InstallState = cortexops.InstallState_Installed
		}
		mcStatus := cluster.Status.Cortex
		if err != nil {
			return nil, err
		}
		status.Version = mcStatus.Version
		if cluster.GetDeletionTimestamp() != nil {
			status.InstallState = cortexops.InstallState_Uninstalling
			status.AppState = cortexops.ApplicationState_Running
		} else {
			if mcStatus.WorkloadsReady {
				status.AppState = cortexops.ApplicationState_Running
			} else {
				status.AppState = cortexops.ApplicationState_Pending
				status.Warnings = append(status.Warnings, mcStatus.Conditions...)
			}
		}
	}

	return status, nil
}

func (k *OpniManager) ShouldDisableNode(_ *corev1.Reference) error {
	stat, err := k.Status(context.TODO(), &emptypb.Empty{})
	if err != nil {
		// can't determine cluster status, so don't disable the node
		return nil
	}
	switch stat.State {
	case cortexops.InstallState_NotInstalled, cortexops.InstallState_Uninstalling:
		return status.Error(codes.Unavailable, fmt.Sprintf("Cortex cluster is not installed"))
	case cortexops.InstallState_Updating, cortexops.InstallState_Installed:
		return nil
	case cortexops.InstallState_Unknown:
		fallthrough
	default:
		// can't determine cluster status, so don't disable the node
		return nil
	}
}

func (k *OpniManager) DryRun(ctx context.Context, req *cortexops.DryRunRequest) (*cortexops.DryRunResponse, error) {
	res, err := k.configTracker.DryRun(ctx, req.Target, req.Action, req.Spec)
	if err != nil {
		return nil, err
	}
	return &cortexops.DryRunResponse{
		Current:          res.Current,
		Modified:         res.Modified,
		ValidationErrors: configutil.CollectValidationErrorLogs(res.Modified.GetCortexConfig()),
	}, nil
}

func (k *OpniManager) ConfigurationHistory(ctx context.Context, req *cortexops.ConfigurationHistoryRequest) (*cortexops.ConfigurationHistoryResponse, error) {
	revisions, err := k.configTracker.History(ctx, req.GetTarget(), storage.IncludeValues(req.GetIncludeValues()))
	if err != nil {
		return nil, err
	}
	resp := &cortexops.ConfigurationHistoryResponse{
		Entries: make([]*cortexops.CapabilityBackendConfigSpec, len(revisions)),
	}
	for i, rev := range revisions {
		if req.IncludeValues {
			spec := rev.Value()
			spec.Revision = corev1.NewRevision(rev.Revision(), rev.Timestamp())
			resp.Entries[i] = spec
		} else {
			resp.Entries[i] = &cortexops.CapabilityBackendConfigSpec{
				Revision: corev1.NewRevision(rev.Revision(), rev.Timestamp()),
			}
		}
	}
	return resp, nil
}

func init() {
	drivers.ClusterDrivers.Register("opni-manager", func(_ context.Context, opts ...driverutil.Option) (drivers.ClusterDriver, error) {
		options := OpniManagerClusterDriverOptions{
			MonitoringCluster: types.NamespacedName{
				Namespace: os.Getenv("POD_NAMESPACE"),
				Name:      "opni",
			},
			GatewayRef: types.NamespacedName{
				Namespace: os.Getenv("POD_NAMESPACE"),
				Name:      os.Getenv("GATEWAY_NAME"),
			},
		}
		if err := driverutil.ApplyOptions(&options, opts...); err != nil {
			return nil, err
		}

		return NewOpniManagerClusterDriver(options)
	})
}
