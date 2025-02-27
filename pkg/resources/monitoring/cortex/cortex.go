package cortex

import (
	"context"

	"github.com/cisco-open/operator-tools/pkg/reconciler"
	corev1beta1 "github.com/rancher/opni/apis/core/v1beta1"
	"github.com/rancher/opni/pkg/logger"
	"github.com/rancher/opni/pkg/resources"
	"github.com/rancher/opni/pkg/util/k8sutil"
	"go.uber.org/zap"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type Reconciler struct {
	reconciler.ResourceReconciler
	ctx    context.Context
	client client.Client
	logger *zap.SugaredLogger
	mc     *corev1beta1.MonitoringCluster
}

func NewReconciler(
	ctx context.Context,
	client client.Client,
	mc *corev1beta1.MonitoringCluster,
) *Reconciler {
	return &Reconciler{
		ResourceReconciler: reconciler.NewReconcilerWith(client,
			reconciler.WithEnableRecreateWorkload(),
			reconciler.WithRecreateErrorMessageCondition(reconciler.MatchImmutableErrorMessages),
			reconciler.WithLog(log.FromContext(ctx)),
			reconciler.WithScheme(client.Scheme()),
		),
		ctx:    ctx,
		client: client,
		mc:     mc,
		logger: logger.New().Named("controller").Named("cortex"),
	}

}

func (r *Reconciler) Reconcile() (*reconcile.Result, error) {
	allResources := []resources.Resource{}

	updated, err := r.updateCortexVersionStatus()
	if err != nil {
		return k8sutil.RequeueErr(err).ResultPtr()
	}
	if updated {
		return k8sutil.Requeue().ResultPtr()
	}

	config, err := r.config()
	if err != nil {
		return nil, err
	}
	allResources = append(allResources, config)

	runtimeConfig := r.runtimeConfig()
	allResources = append(allResources, runtimeConfig)

	fallbackConfig := r.alertmanagerFallbackConfig()
	allResources = append(allResources, fallbackConfig)

	serviceAccount := r.serviceAccount()
	allResources = append(allResources, serviceAccount)

	deployments := r.deployments()
	allResources = append(allResources, deployments...)

	statefulSets := r.statefulSets()
	allResources = append(allResources, statefulSets...)

	services := r.services()
	allResources = append(allResources, services...)

	if op := resources.ReconcileAll(r, allResources); op.ShouldRequeue() {
		return op.ResultPtr()
	}

	// watch cortex components until they are healthy
	if op := r.pollCortexHealth(append(deployments, statefulSets...)); op.ShouldRequeue() {
		return op.ResultPtr()
	}

	return nil, nil
}
