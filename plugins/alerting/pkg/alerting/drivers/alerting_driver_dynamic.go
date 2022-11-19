package drivers

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/banzaicloud/k8s-objectmatcher/patch"
	backoffv2 "github.com/lestrrat-go/backoff/v2"
	corev1beta1 "github.com/rancher/opni/apis/core/v1beta1"
	"github.com/rancher/opni/apis/v1beta2"
	"github.com/rancher/opni/pkg/alerting/backend"
	"github.com/rancher/opni/pkg/alerting/routing"
	"github.com/rancher/opni/pkg/alerting/shared"
	"github.com/rancher/opni/plugins/alerting/pkg/apis/alertops"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	appsv1 "k8s.io/api/apps/v1"
	k8scorev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const opniReloadAnnotation = "opni.io/reload-configmap"

// Implementation of the AlertingDynamic Server for
// AlertingManager

var mu = sync.Mutex{}

func (a *AlertingManager) ConfigFromBackend(ctx context.Context) (*routing.RoutingTree, *routing.OpniInternalRouting, error) {
	mu.Lock()
	defer mu.Unlock()
	rawConfig, err := a.Fetch(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, nil, err
	}
	config, err := routing.NewRoutingTreeFrom(rawConfig.RawAlertManagerConfig)
	if err != nil {
		return nil, nil, err
	}
	internal, err := routing.NewOpniInternalRoutingFrom(rawConfig.RawInternalRouting)
	if err != nil {
		return nil, nil, err
	}
	return config, internal, nil
}

func (a *AlertingManager) ApplyConfigToBackend(
	ctx context.Context,
	config *routing.RoutingTree,
	internal *routing.OpniInternalRouting,
) error {
	mu.Lock()
	defer mu.Unlock()
	rawAlertManagerData, err := config.Marshal()
	if err != nil {
		return err
	}
	rawInternalRoutingData, err := internal.Marshal()
	if err != nil {
		return err
	}
	a.Logger.Debug("updating config map & pod annotations...")
	_, err = a.Update(ctx, &alertops.AlertingConfig{
		RawAlertManagerConfig: string(rawAlertManagerData),
		RawInternalRouting:    string(rawInternalRoutingData),
	})
	if err != nil {
		return err
	}
	a.Logger.Debug("triggering alertmanager reload + injected server hooks...")
	_, err = a.Reload(ctx, &alertops.ReloadInfo{
		UpdatedConfig: string(rawAlertManagerData),
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *AlertingManager) Fetch(ctx context.Context, _ *emptypb.Empty) (*alertops.AlertingConfig, error) {
	lg := a.Logger.With("action", "Fetch")
	name := a.AlertingOptions.ConfigMap
	namespace := a.gatewayRef.Namespace
	cfgMap := &k8scorev1.ConfigMap{}
	err := a.k8sClient.Get(
		ctx,
		client.ObjectKey{
			Name: name, Namespace: namespace},
		cfgMap)

	if err != nil || cfgMap == nil {
		msg := fmt.Sprintf("K8s runtime error, config map: %s/%s not found: %s",
			namespace,
			name,
			err)
		lg.Error(msg)
		returnErr := shared.WithInternalServerError(
			msg,
		)
		return nil, returnErr
	}
	if _, ok := cfgMap.Data[a.configKey]; !ok {
		msg := fmt.Sprintf("K8s runtime error, config map : %s key : %s not found",
			name,
			a.configKey)
		lg.Error(msg)
		return nil, shared.WithInternalServerError(
			msg,
		)
	}
	if _, ok := cfgMap.Data[a.internalRoutingKey]; !ok {
		msg := fmt.Sprintf("K8s runtime error, config map : %s key : %s not found",
			name,
			a.internalRoutingKey)
		lg.Error(msg)
		return nil, shared.WithInternalServerError(
			msg,
		)
	}
	return &alertops.AlertingConfig{
		RawAlertManagerConfig: cfgMap.Data[a.configKey],
		RawInternalRouting:    cfgMap.Data[a.internalRoutingKey],
	}, nil
}

func (a *AlertingManager) Update(ctx context.Context, conf *alertops.AlertingConfig) (*emptypb.Empty, error) {
	lg := a.Logger.With("action", "Update")
	a.configPersistMu.Lock()
	defer a.configPersistMu.Unlock()
	cfgStruct := &routing.RoutingTree{}
	err := cfgStruct.Parse(conf.RawAlertManagerConfig)
	if err != nil {
		return nil, err
	}
	loopError := backend.ReconcileInvalidStateLoop(
		time.Duration(time.Second*10),
		cfgStruct,
		lg)
	if loopError != nil {
		return nil, shared.WithInternalServerError(fmt.Sprintf("failed to reconcile config : %s", loopError))
	}

	mutator := func(object client.Object) error {
		switch gateway := object.(type) {
		case *corev1beta1.Gateway:
			gateway.Spec.Alerting.RawAlertManagerConfig = conf.RawAlertManagerConfig
			gateway.Spec.Alerting.RawInternalRouting = conf.RawInternalRouting
			return nil
		case *v1beta2.Gateway:
			gateway.Spec.Alerting.RawAlertManagerConfig = conf.RawAlertManagerConfig
			gateway.Spec.Alerting.RawInternalRouting = conf.RawInternalRouting
			return nil
		default:
			return fmt.Errorf("unkown gateway type %T", gateway)
		}
	}

	err = retry.OnError(retry.DefaultBackoff, k8serrors.IsConflict, func() error {
		existing, err := a.newOpniGateway()
		if err != nil {
			return err
		}
		err = a.k8sClient.Get(ctx, client.ObjectKeyFromObject(existing), existing)
		if err != nil {
			return err
		}
		var clone client.Object
		switch gateway := existing.(type) {
		case *corev1beta1.Gateway:
			clone = gateway.DeepCopyObject().(client.Object)
		case *v1beta2.Gateway:
			clone = gateway.DeepCopyObject().(client.Object)
		default:
			return fmt.Errorf("unkown gateway type %T", gateway)
		}
		if err := mutator(clone); err != nil {
			return err
		}
		cmp, err := patch.DefaultPatchMaker.Calculate(existing, clone,
			patch.IgnoreStatusFields(),
			patch.IgnoreVolumeClaimTemplateTypeMetaAndStatus(),
			patch.IgnorePDBSelector(),
		)

		if err == nil {
			if cmp.IsEmpty() {
				return status.Error(codes.FailedPrecondition, "no changes to apply")
			}
		}
		return a.k8sClient.Update(ctx, clone)
	})
	if err != nil {
		return nil, err
	}
	lg.Debug("editing statefulsets...")
	// !! must edit statefulset pod annotations to trigger a SELF_DELETE from the
	// !! mounted config symlink inside the alertmanager pod
	controllerSvcData := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      a.AlertingOptions.ControllerNodeService + "-internal",
			Namespace: a.gatewayRef.Namespace,
		},
	}
	err = a.k8sClient.Get(ctx, client.ObjectKeyFromObject(controllerSvcData), controllerSvcData)
	if err != nil {
		lg.Error(err)
		return nil, err
	}
	numReplicas := controllerSvcData.Spec.Replicas
	if numReplicas == nil {
		lg.Error(err)
		return nil, err
	}
	workerSvcData := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      a.AlertingOptions.WorkerNodesService + "-internal",
			Namespace: a.gatewayRef.Namespace,
		},
	}
	var numWorkerReplicas *int32
	zero := int32(0)
	err = a.k8sClient.Get(ctx, client.ObjectKeyFromObject(workerSvcData), workerSvcData)
	if k8serrors.IsNotFound(err) {
		numWorkerReplicas = &zero
	} else if numWorkerReplicas == nil {
		numWorkerReplicas = &zero
	} else {
		numWorkerReplicas = workerSvcData.Spec.Replicas
	}
	var wg sync.WaitGroup
	annotationMutator := func(object client.Object) {
		ann := object.GetAnnotations()
		ann[opniReloadAnnotation] = fmt.Sprintf("%d", time.Now().UnixNano())
		object.SetAnnotations(ann)
	}
	lg.Debugf("number of controller replicas : %d", *numReplicas)
	lg.Debugf("number of worker replicas : %d", *numWorkerReplicas)
	for i := 0; i < int(*numReplicas); i++ {
		i := i // capture loop variable in closure
		wg.Add(1)
		go func() {
			defer wg.Done()
			err = retry.OnError(retry.DefaultBackoff, k8serrors.IsConflict, func() error {
				pod := &k8scorev1.Pod{
					ObjectMeta: metav1.ObjectMeta{
						Name:      fmt.Sprintf("%s-%d", controllerSvcData.ObjectMeta.Name, i),
						Namespace: a.gatewayRef.Namespace,
					},
				}
				err := a.k8sClient.Get(ctx, client.ObjectKeyFromObject(pod), pod)
				if err != nil {
					lg.Error(err)
					return err
				}
				clone := pod.DeepCopyObject().(client.Object)
				annotationMutator(clone)

				cmp, err := patch.DefaultPatchMaker.Calculate(pod, clone,
					patch.IgnoreStatusFields(),
					patch.IgnoreVolumeClaimTemplateTypeMetaAndStatus(),
					patch.IgnorePDBSelector(),
				)
				if err == nil {
					if cmp.IsEmpty() {
						return status.Error(codes.FailedPrecondition, "no changes to apply")
					}
				}
				return a.k8sClient.Update(ctx, clone)
			})
			if err != nil {
				lg.Error(err)
			}
		}()
	}
	for j := 0; j < int(*numWorkerReplicas); j++ {
		j := j // capture loop variable in closure
		wg.Add(1)
		go func() {
			defer wg.Done()
			err = retry.OnError(retry.DefaultBackoff, k8serrors.IsConflict, func() error {
				pod := &k8scorev1.Pod{
					ObjectMeta: metav1.ObjectMeta{
						Name:      fmt.Sprintf("%s-%d", workerSvcData.ObjectMeta.Name, j),
						Namespace: a.gatewayRef.Namespace,
					},
				}
				err := a.k8sClient.Get(ctx, client.ObjectKeyFromObject(pod), pod)
				if err != nil {
					lg.Error(err)
					return err
				}
				clone := pod.DeepCopyObject().(client.Object)
				annotationMutator(clone)

				cmp, err := patch.DefaultPatchMaker.Calculate(pod, clone,
					patch.IgnoreStatusFields(),
					patch.IgnoreVolumeClaimTemplateTypeMetaAndStatus(),
					patch.IgnorePDBSelector(),
				)
				if err == nil {
					if cmp.IsEmpty() {
						return status.Error(codes.FailedPrecondition, "no changes to apply")
					}
				}
				return a.k8sClient.Update(ctx, clone)
			})
			if err != nil {
				lg.Error(err)
			}
		}()
	}
	wg.Wait()
	lg.Debug("updating annotations done")
	return &emptypb.Empty{}, nil
}

type apiConfigRequest struct {
	RawConfig string `json:"config"`
}

type endpoint struct {
	AlertManagerEndpoint string
	OpniEndpoint         string
}

func (a *AlertingManager) Reload(ctx context.Context, reloadInfo *alertops.ReloadInfo) (*emptypb.Empty, error) {
	lg := a.Logger.With("alerting-backend", "k8s", "action", "reload")

	reloadEndpoints := []endpoint{}
	// RELOAD the controller!!!
	reloadEndpoints = append(reloadEndpoints, endpoint{
		AlertManagerEndpoint: a.AlertingOptions.GetControllerEndpoint(),
		OpniEndpoint:         a.AlertingOptions.GetInternalControllerOpniEndpoint(),
	})
	// RELOAD the workers
	name := a.AlertingOptions.WorkerNodesService
	namespace := a.AlertingOptions.Namespace
	workersEndpoints := k8scorev1.Endpoints{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}
	err := a.k8sClient.Get(ctx, client.ObjectKeyFromObject(&workersEndpoints), &workersEndpoints)
	if err != nil {
		return nil, err
	}
	if len(workersEndpoints.Subsets) > 0 {
		addresses := workersEndpoints.Subsets[0].Addresses
		for _, address := range addresses {
			reloadEndpoints = append(reloadEndpoints, endpoint{
				AlertManagerEndpoint: fmt.Sprintf("%s:%d", address.IP, a.AlertingOptions.WorkerNodePort),
				OpniEndpoint:         fmt.Sprintf("%s:%d", address.IP, shared.AlertingDefaultHookPort),
			})
		}
	}
	wg := sync.WaitGroup{}
	errors := &sharedErrors{}
	updatedConfig, err := json.Marshal(apiConfigRequest{RawConfig: reloadInfo.UpdatedConfig})
	if err != nil {
		return nil, err
	}
	for _, endpoint := range reloadEndpoints {
		wg.Add(1)
		endpoint := endpoint //!! must capture in closure
		pipelineRetrier := backoffv2.Exponential(
			backoffv2.WithMinInterval(time.Second*2),
			backoffv2.WithMaxInterval(time.Second*5),
			backoffv2.WithMaxRetries(3),
			backoffv2.WithMultiplier(1.2),
		)
		configReloadRetrier := backoffv2.Exponential(
			backoffv2.WithMinInterval(time.Second*3),
			backoffv2.WithMaxInterval(time.Second*10),
			backoffv2.WithMaxRetries(12),
			backoffv2.WithMultiplier(1.5),
		)
		lg.With("alertmanager-endpoint", endpoint.AlertManagerEndpoint, "opni-endpoint", endpoint.OpniEndpoint).Debug("reloading...")
		go func() {
			defer wg.Done()
			pipelineErr := backend.NewApiPipline(
				ctx,
				[]*backend.AlertManagerAPI{
					backend.NewAlertManagerReadyClient(ctx, endpoint.AlertManagerEndpoint, backend.WithRetrier(pipelineRetrier), backend.WithExpectClosure(backend.NewExpectStatusOk())),
					backend.NewAlertManagerOpniConfigClient(
						ctx,
						endpoint.OpniEndpoint,
						backend.WithRetrier(configReloadRetrier),
						backend.WithRequestBody(updatedConfig),
						backend.WithExpectClosure(backend.NewExpectStatusOk())),
					backend.NewAlertManagerReloadClient(ctx, endpoint.AlertManagerEndpoint,
						backend.WithRetrier(pipelineRetrier), backend.WithExpectClosure(backend.NewExpectStatusOk())),
				},
				&pipelineRetrier,
			)
			if pipelineErr != nil {
				lg.Error(pipelineErr)
				appendError(errors, fmt.Errorf("pipeline error for %s : %s", endpoint, pipelineErr))
			}
		}()
	}
	wg.Wait()
	if len(errors.errors) > 0 {
		return nil, shared.WithInternalServerErrorf("alert backend reload failed %s", strings.Join(func() []string {
			res := []string{}
			for _, e := range errors.errors {
				res = append(res, e.Error())
			}
			return res
		}(), ","))
	}
	return &emptypb.Empty{}, nil
}
