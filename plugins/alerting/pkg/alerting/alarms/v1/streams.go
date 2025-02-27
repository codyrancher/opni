package alarms

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	alertingv1 "github.com/rancher/opni/pkg/apis/alerting/v1"
	corev1 "github.com/rancher/opni/pkg/apis/core/v1"

	"github.com/nats-io/nats.go"
	"github.com/rancher/opni/pkg/alerting/shared"
	"github.com/rancher/opni/pkg/alerting/storage"
	"github.com/rancher/opni/plugins/metrics/apis/cortexadmin"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewAgentStream() *nats.StreamConfig {
	return &nats.StreamConfig{
		Name:      shared.AgentClusterHealthStatusStream,
		Subjects:  []string{shared.AgentClusterHealthStatusSubjects},
		Retention: nats.LimitsPolicy,
		MaxAge:    1 * time.Hour,
		MaxBytes:  1 * 1024 * 50, //50KB
	}
}

func NewAgentDurableReplayConsumer(clusterId string) *nats.ConsumerConfig {
	return &nats.ConsumerConfig{
		Durable:        NewDurableAgentReplaySubject(clusterId),
		DeliverSubject: NewDurableAgentReplaySubject(clusterId),
		DeliverPolicy:  nats.DeliverNewPolicy,
		FilterSubject:  NewAgentStreamSubject(clusterId),
		AckPolicy:      nats.AckExplicitPolicy,
		ReplayPolicy:   nats.ReplayInstantPolicy,
	}
}

func NewDurableAgentReplaySubject(clusterId string) string {
	return fmt.Sprintf("%s-%s", shared.AgentClusterHealthStatusDurableReplay, clusterId)
}

func NewAgentStreamSubject(clusterId string) string {
	return fmt.Sprintf("%s.%s", shared.AgentClusterHealthStatusStream, clusterId)
}

func NewCortexStatusStream() *nats.StreamConfig {
	return &nats.StreamConfig{
		Name:      shared.CortexStatusStream,
		Subjects:  []string{shared.CortexStatusStreamSubjects},
		Retention: nats.LimitsPolicy,
		MaxAge:    1 * time.Hour,
		MaxBytes:  1 * 1024 * 50, //50KB
	}
}

func NewCortexStatusSubject() string {
	return fmt.Sprintf("%s.%s", shared.CortexStatusStream, "cortex")
}

func (p *AlarmServerComponent) onSystemConditionCreate(conditionId, conditionName, namespace string, condition *alertingv1.AlertCondition) error {
	lg := p.logger.With("onSystemConditionCreate", conditionId)
	lg.Debugf("received condition update: %v", condition)
	disconnect := condition.GetAlertType().GetSystem()
	jsCtx, cancel := context.WithCancel(p.ctx)
	lg.Debugf("Creating agent disconnect with timeout %s", disconnect.GetTimeout().AsDuration())
	agentId := condition.GetClusterId().Id
	evaluator := NewInternalConditionEvaluator(
		&internalConditionMetadata{
			conditionId:        conditionId,
			conditionName:      conditionName,
			lg:                 lg,
			clusterId:          agentId,
			alertmanagerlabels: map[string]string{},
		},
		&internalConditionContext{
			parentCtx:        p.ctx,
			evaluationCtx:    jsCtx,
			evaluateInterval: time.Second * 10,
			cancelEvaluation: cancel,
			evaluateDuration: disconnect.GetTimeout().AsDuration(),
		},
		&internalConditionStorage{
			js:              p.js.Get(),
			durableConsumer: NewAgentDurableReplayConsumer(agentId),
			streamSubject:   NewAgentStreamSubject(agentId),
			stateStorage:    p.stateStorage.Get(),
			incidentStorage: p.incidentStorage.Get(),
			msgCh:           make(chan *nats.Msg, 32),
		},
		&internalConditionState{},
		&internalConditionHooks[*corev1.ClusterHealthStatus]{
			healthOnMessage: func(h *corev1.ClusterHealthStatus) (health bool, ts *timestamppb.Timestamp) {
				lg.Debugf("received agent health update connected %v : %s", h.HealthStatus.Status.Connected, h.HealthStatus.Status.Timestamp.String())
				return h.HealthStatus.Status.Connected, h.HealthStatus.Status.Timestamp
			},
			triggerHook: func(ctx context.Context, conditionId string, labels, annotations map[string]string) {
				p.notifications.TriggerAlerts(ctx, &alertingv1.TriggerAlertsRequest{
					ConditionId:   &corev1.Reference{Id: conditionId},
					ConditionName: conditionName,
					Namespace:     namespace,
					Labels:        lo.Assign(condition.GetRoutingLabels(), labels),
					Annotations:   lo.Assign(condition.GetRoutingAnnotations(), annotations),
				})
			},
			resolveHook: func(ctx context.Context, conditionId string, labels, annotations map[string]string) {
				_, _ = p.notifications.ResolveAlerts(ctx, &alertingv1.ResolveAlertsRequest{
					ConditionId:   &corev1.Reference{Id: conditionId},
					ConditionName: conditionName,
					Namespace:     namespace,
					Labels:        lo.Assign(condition.GetRoutingLabels(), labels),
					Annotations:   lo.Assign(condition.GetRoutingAnnotations(), annotations),
				})
			},
		},
	)
	// handles re-entrant conditions
	evaluator.CalculateInitialState()
	go func() {
		defer cancel() // cancel parent context, if we return (non-recoverable)
		evaluator.SubscriberLoop()
	}()
	// spawn a watcher for the incidents
	go func() {
		defer cancel() // cancel parent context, if we return (non-recoverable)
		evaluator.EvaluateLoop()
	}()
	p.runner.AddSystemConfigListener(conditionId, EvaluatorContext{
		Ctx:    evaluator.evaluationCtx,
		Cancel: evaluator.cancelEvaluation,
	})
	return nil
}

func (p *AlarmServerComponent) onDownstreamCapabilityConditionCreate(conditionId, conditionName, namespace string, condition *alertingv1.AlertCondition) error {
	lg := p.logger.With("onCapabilityStatusCreate", conditionId)
	capability := condition.GetAlertType().GetDownstreamCapability()
	lg.Debugf("received condition update: %v", condition)
	jsCtx, cancel := context.WithCancel(p.ctx)
	lg.Debugf("Creating agent capability unhealthy with timeout %s", capability.GetFor().AsDuration())
	agentId := condition.GetClusterId().Id
	evaluator := NewInternalConditionEvaluator(
		&internalConditionMetadata{
			conditionId:        conditionId,
			conditionName:      conditionName,
			lg:                 lg,
			clusterId:          agentId,
			alertmanagerlabels: map[string]string{},
		},
		&internalConditionContext{
			parentCtx:        p.ctx,
			evaluationCtx:    jsCtx,
			evaluateInterval: time.Second * 10,
			cancelEvaluation: cancel,
			evaluateDuration: capability.GetFor().AsDuration(),
		},
		&internalConditionStorage{
			js:              p.js.Get(),
			durableConsumer: NewAgentDurableReplayConsumer(agentId),
			streamSubject:   NewAgentStreamSubject(agentId),
			stateStorage:    p.stateStorage.Get(),
			incidentStorage: p.incidentStorage.Get(),
			msgCh:           make(chan *nats.Msg, 32),
		},
		&internalConditionState{},
		&internalConditionHooks[*corev1.ClusterHealthStatus]{
			healthOnMessage: func(h *corev1.ClusterHealthStatus) (healthy bool, ts *timestamppb.Timestamp) {
				healthy = true
				if h.HealthStatus.Health == nil {
					return false, h.HealthStatus.Status.Timestamp
				}
				lg.Debugf("found health conditions %v", h.HealthStatus.Health.Conditions)
				for _, s := range h.HealthStatus.Health.Conditions {
					for _, badState := range capability.GetCapabilityState() {
						if strings.Contains(s, badState) {
							healthy = false
							break
						}
					}
				}
				return healthy, h.HealthStatus.Status.Timestamp
			},
			triggerHook: func(ctx context.Context, conditionId string, labels, annotations map[string]string) {
				_, _ = p.notifications.TriggerAlerts(ctx, &alertingv1.TriggerAlertsRequest{
					ConditionId:   &corev1.Reference{Id: conditionId},
					ConditionName: conditionName,
					Namespace:     namespace,
					Labels:        lo.Assign(condition.GetRoutingLabels(), labels),
					Annotations:   lo.Assign(condition.GetRoutingAnnotations(), annotations),
				})
			},
			resolveHook: func(ctx context.Context, conditionId string, labels, annotations map[string]string) {
				_, _ = p.notifications.ResolveAlerts(ctx, &alertingv1.ResolveAlertsRequest{
					ConditionId:   &corev1.Reference{Id: conditionId},
					ConditionName: conditionName,
					Namespace:     namespace,
					Labels:        lo.Assign(condition.GetRoutingLabels(), labels),
					Annotations:   lo.Assign(condition.GetRoutingAnnotations(), annotations),
				})
			},
		},
	)
	// handles re-entrant conditions
	evaluator.CalculateInitialState()
	go func() {
		defer cancel() // cancel parent context, if we return (non-recoverable)
		evaluator.SubscriberLoop()
	}()
	// spawn a watcher for the incidents
	go func() {
		defer cancel() // cancel parent context, if we return (non-recoverable)
		evaluator.EvaluateLoop()
	}()
	p.runner.AddSystemConfigListener(conditionId, EvaluatorContext{
		Ctx:    evaluator.evaluationCtx,
		Cancel: evaluator.cancelEvaluation,
	})
	return nil
}

func reduceCortexAdminStates(componentsToTrack []string, cStatus *cortexadmin.CortexStatus) (healthy bool, ts *timestamppb.Timestamp) {
	if cStatus == nil {
		return false, timestamppb.Now()
	}
	ts = cStatus.GetTimestamp()
	// helps track status errors to particular components, like having 3 expected replicas, but only 1-2 are running
	memberReports := map[string]bool{}
	for _, cmp := range componentsToTrack {
		switch cmp {
		case shared.CortexDistributor:
			if cStatus.Distributor == nil {
				return false, ts
			}
			for _, svc := range cStatus.Distributor.GetServices().GetServices() {
				memberReports[svc.GetName()] = true
				if svc.GetStatus() != "Running" {
					return false, ts
				}
			}
		case shared.CortexIngester:
			if cStatus.Ingester == nil {
				return false, ts
			}
			for _, member := range cStatus.Ingester.Memberlist.Members.Items {
				if _, ok := memberReports[member.Name]; !ok {
					memberReports[member.Name] = true
				}
			}
			for _, svc := range cStatus.Ingester.GetServices().GetServices() {
				memberReports[svc.GetName()] = true
				if svc.GetStatus() != "Running" {
					return false, ts
				}
			}
		case shared.CortexRuler:
			if cStatus.Ruler == nil {
				return false, ts
			}
			for _, member := range cStatus.Ruler.Memberlist.Members.Items {
				if _, ok := memberReports[member.Name]; !ok {
					memberReports[member.Name] = true
				}
			}
			for _, svc := range cStatus.Ruler.GetServices().GetServices() {
				memberReports[svc.GetName()] = true
				if svc.GetStatus() != "Running" {
					return false, ts
				}
			}
		case shared.CortexPurger:
			if cStatus.Purger == nil {
				return false, ts
			}
			for _, svc := range cStatus.Purger.GetServices().GetServices() {
				memberReports[svc.GetName()] = true
				if svc.GetStatus() != "Running" {
					return false, ts
				}
			}
		case shared.CortexCompactor:
			if cStatus.Compactor == nil {
				return false, ts
			}
			for _, member := range cStatus.Compactor.Memberlist.Members.Items {
				if _, ok := memberReports[member.Name]; !ok {
					memberReports[member.Name] = true
				}
			}
			for _, svc := range cStatus.Compactor.GetServices().GetServices() {
				memberReports[svc.GetName()] = true
				if svc.GetStatus() != "Running" {
					return false, ts
				}
			}
		case shared.CortexStoreGateway:
			if cStatus.StoreGateway == nil {
				return false, ts
			}
			for _, svc := range cStatus.StoreGateway.GetServices().GetServices() {
				memberReports[svc.GetName()] = true
				if svc.GetStatus() != "Running" {
					return false, ts
				}
			}
		case shared.CortexQueryFrontend:
			if cStatus.QueryFrontend == nil {
				return false, ts
			}
			for _, svc := range cStatus.QueryFrontend.GetServices().GetServices() {
				memberReports[svc.GetName()] = true
				if svc.GetStatus() != "Running" {
					return false, ts
				}
			}
		case shared.CortexQuerier:
			if cStatus.Querier == nil {
				return false, ts
			}
			for _, svc := range cStatus.Querier.GetServices().GetServices() {
				memberReports[svc.GetName()] = true
				if svc.GetStatus() != "Running" {
					return false, ts
				}
			}
		}
	}
	// on cortex-status error, if a specific component is not reported, we assume it is unhealthy
	for _, component := range componentsToTrack {
		for member, reportedOn := range memberReports {
			if strings.Contains(member, component) && !reportedOn {
				return false, ts
			}
		}
	}
	return true, ts
}

func (p *AlarmServerComponent) onCortexClusterStatusCreate(conditionId, conditionName, namespace string, condition *alertingv1.AlertCondition) error {
	lg := p.logger.With("onCortexClusterStatusCreate", conditionId)
	cortex := condition.GetAlertType().GetMonitoringBackend()
	lg.Debugf("received condition update: %v", condition)
	jsCtx, cancel := context.WithCancel(p.ctx)
	lg.Debugf("Creating cortex status with timeout %s", cortex.GetFor().AsDuration())

	evaluator := NewInternalConditionEvaluator(
		&internalConditionMetadata{
			conditionId:        conditionId,
			conditionName:      conditionName,
			lg:                 lg,
			clusterId:          "", // unused here
			alertmanagerlabels: map[string]string{},
		},
		&internalConditionContext{
			parentCtx:        p.ctx,
			evaluationCtx:    jsCtx,
			evaluateInterval: time.Minute,
			cancelEvaluation: cancel,
			evaluateDuration: cortex.GetFor().AsDuration(),
		},
		&internalConditionStorage{
			js:              p.js.Get(),
			durableConsumer: nil,
			streamSubject:   NewCortexStatusSubject(),
			stateStorage:    p.stateStorage.Get(),
			incidentStorage: p.incidentStorage.Get(),
			msgCh:           make(chan *nats.Msg, 32),
		},
		&internalConditionState{},
		&internalConditionHooks[*cortexadmin.CortexStatus]{
			healthOnMessage: func(h *cortexadmin.CortexStatus) (healthy bool, ts *timestamppb.Timestamp) {
				return reduceCortexAdminStates(cortex.GetBackendComponents(), h)
			},
			triggerHook: func(ctx context.Context, conditionId string, labels, annotations map[string]string) {
				_, _ = p.notifications.TriggerAlerts(ctx, &alertingv1.TriggerAlertsRequest{
					ConditionId:   &corev1.Reference{Id: conditionId},
					ConditionName: conditionName,
					Namespace:     namespace,
					Labels:        lo.Assign(condition.GetRoutingLabels(), labels),
					Annotations:   lo.Assign(condition.GetRoutingAnnotations(), annotations),
				})
			},
			resolveHook: func(ctx context.Context, conditionId string, labels, annotations map[string]string) {
				lg.Debug("resolve cortex status condition")
				_, _ = p.notifications.ResolveAlerts(ctx, &alertingv1.ResolveAlertsRequest{
					ConditionId:   &corev1.Reference{Id: conditionId},
					ConditionName: conditionName,
					Namespace:     namespace,
					Labels:        lo.Assign(condition.GetRoutingLabels(), labels),
					Annotations:   lo.Assign(condition.GetRoutingAnnotations(), annotations),
				})
			},
		},
	)
	// handles re-entrant conditions
	evaluator.CalculateInitialState()
	go func() {
		defer cancel() // cancel parent context, if we return (non-recoverable)
		evaluator.SubscriberLoop()
	}()
	// spawn a watcher for the incidents
	go func() {
		defer cancel() // cancel parent context, if we return (non-recoverable)
		evaluator.EvaluateLoop()
	}()
	p.runner.AddSystemConfigListener(conditionId, EvaluatorContext{
		Ctx:    evaluator.evaluationCtx,
		Cancel: evaluator.cancelEvaluation,
	})
	return nil
}

type internalConditionMetadata struct {
	lg                 *zap.SugaredLogger
	conditionName      string
	conditionId        string
	clusterId          string
	alertmanagerlabels map[string]string
}

type internalConditionContext struct {
	parentCtx        context.Context
	cancelEvaluation context.CancelFunc
	evaluateDuration time.Duration
	evaluationCtx    context.Context
	evaluateInterval time.Duration
}

type internalConditionStorage struct {
	js              nats.JetStreamContext
	streamSubject   string
	durableConsumer *nats.ConsumerConfig
	incidentStorage storage.IncidentStorage
	stateStorage    storage.StateStorage

	msgCh chan *nats.Msg
}

type internalConditionState struct {
	inMemoryFiring bool
	stateLock      sync.Mutex
	firingLock     sync.RWMutex
}

type internalConditionHooks[T proto.Message] struct {
	healthOnMessage func(h T) (healthy bool, ts *timestamppb.Timestamp)
	triggerHook     func(ctx context.Context, conditionId string, labels, annotations map[string]string)
	resolveHook     func(ctx context.Context, conditionId string, labels, annotations map[string]string)
}

func NewInternalConditionEvaluator[T proto.Message](
	metadata *internalConditionMetadata,
	context *internalConditionContext,
	storage *internalConditionStorage,
	state *internalConditionState,
	hooks *internalConditionHooks[T],
) *InternalConditionEvaluator[T] {
	return &InternalConditionEvaluator[T]{
		metadata,
		context,
		storage,
		state,
		hooks,
		"",
	}
}

// --------------------------------
type InternalConditionEvaluator[T proto.Message] struct {
	*internalConditionMetadata
	*internalConditionContext
	*internalConditionStorage
	*internalConditionState
	*internalConditionHooks[T]
	fingerprint string
}

// infinite & blocking : must be run in a goroutine
func (c *InternalConditionEvaluator[T]) SubscriberLoop() {
	defer c.cancelEvaluation()
	//replay consumer if it exists
	t := time.NewTicker(c.evaluateInterval)
	defer t.Stop()
	for {
		shouldExit := false
		select {
		case <-c.evaluationCtx.Done():
			return
		case <-t.C:
			subStream, err := c.js.ChanSubscribe(c.streamSubject, c.msgCh)
			if err != nil {
				c.lg.Warn("failed to subscribe to stream %s", err)
				continue
			}
			defer subStream.Unsubscribe()
			if err != nil {
				continue
			}
			shouldExit = true
		}
		if shouldExit {
			break
		}
	}
	t.Stop()
	for {
		select {
		case <-c.parentCtx.Done():
			c.lg.Info("parent context is exiting, exiting evaluation loop")
			return
		case <-c.evaluationCtx.Done():
			c.lg.Info("evaluation context is exiting, exiting evaluation loop")
			return
		case msg := <-c.msgCh:
			var status T
			err := json.Unmarshal(msg.Data, &status)
			if err != nil {
				c.lg.Error(err)
			}
			healthy, ts := c.healthOnMessage(status)
			incomingState := alertingv1.CachedState{
				Healthy:   healthy,
				Firing:    c.IsFiring(),
				Timestamp: ts,
			}
			c.UpdateState(c.evaluationCtx, &incomingState)
			msg.Ack()
		}
	}
}

// infinite & blocking : must be run in a goroutine
func (c *InternalConditionEvaluator[T]) EvaluateLoop() {
	defer c.cancelEvaluation() // cancel parent context, if we return (non-recoverable)
	ticker := time.NewTicker(c.evaluateInterval)
	defer ticker.Stop()
	for {
		select {
		case <-c.parentCtx.Done():
			c.lg.Info("parent context is exiting, exiting evaluation loop")
			return
		case <-c.evaluationCtx.Done():
			c.lg.Info("evaluation context is exiting, exiting evaluation loop")
			return
		case <-ticker.C:
			lastKnownState, err := c.stateStorage.Get(c.evaluationCtx, c.conditionId)
			if err != nil {
				continue
			}
			if !lastKnownState.Healthy {
				c.lg.Debugf("condition %s is unhealthy", c.conditionName)
				interval := timestamppb.Now().AsTime().Sub(lastKnownState.Timestamp.AsTime())
				if interval > c.evaluateDuration { // then we must fire an alert
					if !c.IsFiring() {
						c.fingerprint = strconv.Itoa(int(time.Now().Unix()))
						c.SetFiring(true)
						err = c.UpdateState(c.evaluationCtx, &alertingv1.CachedState{
							Healthy:   lastKnownState.Healthy,
							Firing:    c.IsFiring(),
							Timestamp: timestamppb.Now(),
						})
						if err != nil {
							c.lg.Error(err)
						}
						err = c.incidentStorage.OpenInterval(c.evaluationCtx, c.conditionId, c.fingerprint, timestamppb.Now())
						if err != nil {
							c.lg.Error(err)
						}
					}
					c.lg.Debugf("triggering alert for condition %s", c.conditionName)
					c.triggerHook(c.evaluationCtx, c.conditionId, map[string]string{
						alertingv1.NotificationPropertyFingerprint: c.fingerprint,
					}, map[string]string{
						alertingv1.NotificationPropertyFingerprint: c.fingerprint,
					})
				}
			} else if lastKnownState.Healthy && c.IsFiring() &&
				// avoid potential noise from api streams & replays
				lastKnownState.Timestamp.AsTime().Add(-c.evaluateInterval).Before(time.Now()) {
				c.lg.Debugf("condition %s is now healthy again after having fired", c.conditionName)
				c.SetFiring(false)
				err = c.incidentStorage.CloseInterval(c.evaluationCtx, c.conditionId, c.fingerprint, timestamppb.Now())
				if err != nil {
					c.lg.Error(err)
				}
				c.resolveHook(c.evaluationCtx, c.conditionId, map[string]string{
					alertingv1.NotificationPropertyFingerprint: c.fingerprint,
				}, map[string]string{
					alertingv1.NotificationPropertyFingerprint: c.fingerprint,
				})
				c.fingerprint = ""
			}
		}
	}
}

func (c *InternalConditionEvaluator[T]) SetFiring(firing bool) {
	c.firingLock.Lock()
	defer c.firingLock.Unlock()
	c.inMemoryFiring = firing
}

func (c *InternalConditionEvaluator[T]) IsFiring() bool {
	c.firingLock.RLock()
	defer c.firingLock.RUnlock()
	return c.inMemoryFiring
}

func (c *InternalConditionEvaluator[T]) UpdateState(ctx context.Context, s *alertingv1.CachedState) error {
	c.stateLock.Lock()
	defer c.stateLock.Unlock()
	if c.stateStorage.IsDiff(ctx, c.conditionId, s) {
		return c.stateStorage.Put(ctx, c.conditionId, s)
	}
	return nil
}

func (c *InternalConditionEvaluator[T]) CalculateInitialState() {
	incomingState := alertingv1.DefaultCachedState()
	if _, getErr := c.incidentStorage.Get(c.evaluationCtx, c.conditionId); errors.Is(nats.ErrKeyNotFound, getErr) {
		err := c.incidentStorage.Put(c.evaluationCtx, c.conditionId, alertingv1.NewIncidentIntervals())
		if err != nil {
			c.lg.Error(err)
			c.cancelEvaluation()
			return
		}
	} else if getErr != nil {
		c.lg.Error(getErr)
	}
	if st, getErr := c.stateStorage.Get(c.evaluationCtx, c.conditionId); errors.Is(nats.ErrKeyNotFound, getErr) {
		if err := c.stateStorage.Put(c.evaluationCtx, c.conditionId, incomingState); err != nil {
			c.cancelEvaluation()
			return
		}
	} else if getErr == nil {
		incomingState = st
	}
	if incomingState.Firing { // need to update this in memory value
		c.SetFiring(true)
	}
	_ = c.UpdateState(c.evaluationCtx, incomingState)
}
