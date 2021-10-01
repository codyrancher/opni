package opnicluster

import (
	"fmt"
	"net/url"
	"strings"

	"emperror.dev/errors"
	"github.com/banzaicloud/operator-tools/pkg/reconciler"
	"github.com/go-logr/logr"
	"github.com/rancher/opni/apis/v1beta1"
	"github.com/rancher/opni/pkg/features"
	"github.com/rancher/opni/pkg/resources"
	"github.com/rancher/opni/pkg/resources/hyperparameters"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/pointer"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	natsNkeyDir = "/etc/nkey"
)

func (r *Reconciler) opniServices() ([]resources.Resource, error) {
	return []resources.Resource{
		r.nulogHyperparameters,
		r.inferenceDeployment,
		r.drainDeployment,
		r.payloadReceiverDeployment,
		r.payloadReceiverService,
		r.preprocessingDeployment,
		r.gpuCtrlDeployment,
		r.metricsDeployment,
	}, nil
}

func (r *Reconciler) pretrainedModels() (resourceList []resources.Resource, retError error) {
	resourceList = []resources.Resource{}
	lg := logr.FromContext(r.ctx)
	requirement, err := labels.NewRequirement(
		resources.PretrainedModelLabel,
		selection.Exists,
		[]string{},
	)
	if err != nil {
		panic(err)
	}
	deployments := appsv1.DeploymentList{}
	err = r.client.List(r.ctx, &deployments, &client.ListOptions{
		Namespace:     r.opniCluster.Namespace,
		LabelSelector: labels.NewSelector().Add(*requirement),
	})
	if err != nil {
		retError = err
		lg.Error(err, "failed to look up existing deployments")
		return
	}

	// Given a list of desired models, ensure the corresponding deployments exist.
	// If there are existing deployments that have been removed from the
	// list of desired models, mark them as deleted.
	models := map[string]corev1.LocalObjectReference{}
	desiredModels := r.opniCluster.Spec.Services.Inference.PretrainedModels
	for _, model := range desiredModels {
		models[model.Name] = model
	}
	existing := map[string]struct{}{}

	for _, deployment := range deployments.Items {
		// If the deployment is not in the list of desired models, mark it as deleted.
		// Otherwise, simply append the deployment to the list of resources which
		// should be present.
		modelName := deployment.Labels[resources.PretrainedModelLabel]
		existing[modelName] = struct{}{}
		if _, ok := models[modelName]; !ok {
			lg.Info("deleting pretrained model deployment", "model", modelName)
			resourceList = append(resourceList, resources.Absent(deployment.DeepCopy()))
		} else {
			deployment, err := r.pretrainedModelDeployment(models[modelName])
			if err != nil {
				lg.Error(err, "failed to reconcile existing model")
				retError = err
				continue
			}
			resourceList = append(resourceList, deployment)
		}
	}

	// Create new deployments for any that don't exist yet
	for k, v := range models {
		if _, ok := existing[k]; !ok {
			deployment, err := r.pretrainedModelDeployment(v)
			if err != nil {
				lg.Error(err, "failed to create pretrained model resource")
				retError = err
				continue
			}
			lg.Info("creating pretrained model deployment", "model", k)
			resourceList = append(resourceList, deployment)
		}
	}
	return
}

func (r *Reconciler) pretrainedModelDeployment(
	modelRef corev1.LocalObjectReference,
) (resources.Resource, error) {
	model, err := r.findPretrainedModel(modelRef)
	if err != nil {
		return nil, err
	}
	// Create a sidecar container either for downloading the model or copying it
	// directly from an image.
	var sidecar corev1.Container
	switch {
	case model.Spec.HTTP != nil:
		sidecar = httpSidecar(model)
	case model.Spec.Container != nil:
		sidecar = containerSidecar(model)
	default:
		return nil, fmt.Errorf(
			"model %q is invalid. Must specify either HTTP or Container", modelRef.Name)
	}

	return func() (runtime.Object, reconciler.DesiredState, error) {
		labels := resources.CombineLabels(
			r.serviceLabels(v1beta1.InferenceService),
			r.pretrainedModelLabels(model.Name),
		)
		imageSpec := r.serviceImageSpec(v1beta1.InferenceService)
		lg := logr.FromContext(r.ctx)
		lg.V(1).Info("generating pretrained model deployment", "name", model.Name)
		envVars, volumeMounts, volumes := r.genericEnvAndVolumes()
		s3EnvVars := r.s3EnvVars()
		envVars = append(envVars, s3EnvVars...)
		volumes = append(volumes, corev1.Volume{
			Name: "model-volume",
			VolumeSource: corev1.VolumeSource{
				EmptyDir: &corev1.EmptyDirVolumeSource{},
			},
		})
		volumeMounts = append(volumeMounts, corev1.VolumeMount{
			Name:      "model-volume",
			MountPath: "/model/",
		})
		deployment := &appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("opni-inference-%s", model.Name),
				Namespace: r.opniCluster.Namespace,
				OwnerReferences: []metav1.OwnerReference{
					{
						APIVersion: r.opniCluster.APIVersion,
						Kind:       r.opniCluster.Kind,
						Name:       r.opniCluster.Name,
						UID:        r.opniCluster.UID,
					},
				},
				Labels: labels,
			},
			Spec: appsv1.DeploymentSpec{
				Selector: &metav1.LabelSelector{
					MatchLabels: labels,
				},
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: labels,
						Annotations: map[string]string{
							"opni.io/hyperparameters": hyperparameters.GenerateHyperParametersHash(model.Spec.Hyperparameters),
						},
					},
					Spec: corev1.PodSpec{
						InitContainers: []corev1.Container{sidecar},
						Volumes:        volumes,
						Containers: []corev1.Container{
							{
								Name:            "inference-service",
								Image:           imageSpec.GetImage(),
								ImagePullPolicy: imageSpec.GetImagePullPolicy(),
								VolumeMounts:    volumeMounts,
								Env:             envVars,
							},
						},
						ImagePullSecrets: maybeImagePullSecrets(model),
						Tolerations:      r.serviceTolerations(v1beta1.InferenceService),
						NodeSelector:     r.serviceNodeSelector(v1beta1.InferenceService),
					},
				},
			},
		}
		insertHyperparametersVolume(deployment, model.Name)
		return deployment, reconciler.StatePresent, nil
	}, nil
}

func maybeImagePullSecrets(model v1beta1.PretrainedModel) []corev1.LocalObjectReference {
	if model.Spec.Container != nil {
		return model.Spec.Container.ImagePullSecrets
	}
	return nil
}

func httpSidecar(model v1beta1.PretrainedModel) corev1.Container {
	return corev1.Container{
		Name:  "download-model",
		Image: "docker.io/curlimages/curl:latest",
		Args: []string{
			"--silent",
			"--location",
			"--remote-name",
			"--output-dir", "/model/",
			"--url", model.Spec.HTTP.URL,
		},
		VolumeMounts: []corev1.VolumeMount{
			{
				Name:      "model-volume",
				MountPath: "/model/",
			},
		},
	}
}

func containerSidecar(model v1beta1.PretrainedModel) corev1.Container {
	return corev1.Container{
		Name:    "copy-model",
		Image:   model.Spec.Container.Image,
		Command: []string{"/bin/sh"},
		Args:    strings.Fields(`-c cp /staging/* /model/ && chmod -R a+r /model/`),
		VolumeMounts: []corev1.VolumeMount{
			{
				Name:      "model-volume",
				MountPath: "/model/model.tar.gz",
				SubPath:   "model.tar.gz",
			},
		},
	}
}

func (r *Reconciler) gpuWorkerContainer() corev1.Container {
	imageSpec := r.serviceImageSpec(v1beta1.InferenceService)
	envVars, volumeMounts, _ := r.genericEnvAndVolumes()
	envVars = append(envVars, r.s3EnvVars()...)
	envVars = append(envVars, []corev1.EnvVar{
		{
			Name:  "MODEL_THRESHOLD",
			Value: "0.5",
		},
		{
			Name:  "MIN_LOG_TOKENS",
			Value: "5",
		},
		{
			Name:  "IS_GPU_SERVICE",
			Value: "True",
		},
		{
			Name:  "NVIDIA_VISIBLE_DEVICES",
			Value: "all",
		},
		{
			Name:  "NVIDIA_DRIVER_CAPABILITIES",
			Value: "compute,utility",
		},
	}...)
	if r.opniCluster.Spec.S3.NulogS3Bucket != "" {
		envVars = append(envVars,
			corev1.EnvVar{
				Name:  "S3_BUCKET",
				Value: r.opniCluster.Spec.S3.NulogS3Bucket,
			})
	}
	return corev1.Container{
		Name:            "gpu-service-worker",
		Image:           *imageSpec.Image,
		ImagePullPolicy: imageSpec.GetImagePullPolicy(),
		Env:             envVars,
		VolumeMounts:    volumeMounts,
		Resources: corev1.ResourceRequirements{
			Limits: corev1.ResourceList{
				"nvidia.com/gpu": resource.MustParse("1"),
			},
		},
	}
}

func (r *Reconciler) findPretrainedModel(
	modelRef corev1.LocalObjectReference,
) (v1beta1.PretrainedModel, error) {
	model := v1beta1.PretrainedModel{}
	err := r.client.Get(r.ctx, types.NamespacedName{
		Name:      modelRef.Name,
		Namespace: r.opniCluster.Namespace,
	}, &model)
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *Reconciler) genericDeployment(service v1beta1.ServiceKind) *appsv1.Deployment {
	labels := r.serviceLabels(service)
	imageSpec := r.serviceImageSpec(service)
	envVars, volumeMounts, volumes := r.genericEnvAndVolumes()

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      labels[resources.AppNameLabel],
			Namespace: r.opniCluster.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				{
					APIVersion: r.opniCluster.APIVersion,
					Kind:       r.opniCluster.Kind,
					Name:       r.opniCluster.Name,
					UID:        r.opniCluster.UID,
				},
			},
			Labels: labels,
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            labels[resources.AppNameLabel],
							Image:           *imageSpec.Image,
							ImagePullPolicy: imageSpec.GetImagePullPolicy(),
							Env:             envVars,
							VolumeMounts:    volumeMounts,
						},
					},
					ImagePullSecrets: imageSpec.ImagePullSecrets,
					Volumes:          volumes,
					NodeSelector:     r.serviceNodeSelector(service),
					Tolerations:      r.serviceTolerations(service),
				},
			},
		},
	}
	return deployment
}

func (r *Reconciler) genericEnvAndVolumes() (
	envVars []corev1.EnvVar,
	volumeMounts []corev1.VolumeMount,
	volumes []corev1.Volume,
) {
	envVars = append(envVars, corev1.EnvVar{
		Name: "NATS_SERVER_URL",
		Value: fmt.Sprintf("nats://%s-nats-client.%s.svc:%d",
			r.opniCluster.Name, r.opniCluster.Namespace, natsDefaultClientPort),
	})
	switch r.opniCluster.Spec.Nats.AuthMethod {
	case v1beta1.NatsAuthUsername:
		var username string
		if r.opniCluster.Spec.Nats.Username == "" {
			username = "nats-user"
		} else {
			username = r.opniCluster.Spec.Nats.Username
		}
		newEnvVars := []corev1.EnvVar{
			{
				Name:  "NATS_USERNAME",
				Value: username,
			},
			{
				Name: "NATS_PASSWORD",
				ValueFrom: &corev1.EnvVarSource{
					SecretKeyRef: r.opniCluster.Status.Auth.AuthSecretKeyRef,
				},
			},
		}
		envVars = append(envVars, newEnvVars...)
	case v1beta1.NatsAuthNkey:
		newVolumes := []corev1.Volume{
			{
				Name: "nkey",
				VolumeSource: corev1.VolumeSource{
					Secret: &corev1.SecretVolumeSource{
						SecretName: r.opniCluster.Status.Auth.AuthSecretKeyRef.Name,
					},
				},
			},
		}
		volumes = append(volumes, newVolumes...)
		newVolumeMounts := []corev1.VolumeMount{
			{
				Name:      "nkey",
				ReadOnly:  true,
				MountPath: natsNkeyDir,
			},
		}
		volumeMounts = append(volumeMounts, newVolumeMounts...)
		newEnvVars := []corev1.EnvVar{
			{
				Name:  "NKEY_SEED_FILENAME",
				Value: fmt.Sprintf("%s/seed", natsNkeyDir),
			},
		}
		envVars = append(envVars, newEnvVars...)
	}
	envVars = append(envVars, corev1.EnvVar{
		Name:  "ES_ENDPOINT",
		Value: fmt.Sprintf("http://opni-es-client.%s.svc:9200", r.opniCluster.Namespace),
	}, corev1.EnvVar{
		Name:  "ES_USERNAME",
		Value: "admin",
	}, corev1.EnvVar{
		Name:  "ES_PASSWORD",
		Value: "admin",
	})
	return
}

func (r *Reconciler) s3EnvVars() (envVars []corev1.EnvVar) {
	// lg := logr.FromContext(r.ctx)
	if r.opniCluster.Status.Auth.S3AccessKey != nil &&
		r.opniCluster.Status.Auth.S3SecretKey != nil &&
		r.opniCluster.Status.Auth.S3Endpoint != "" {
		envVars = append(envVars, corev1.EnvVar{
			Name: "S3_ACCESS_KEY",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: r.opniCluster.Status.Auth.S3AccessKey,
			},
		}, corev1.EnvVar{
			Name: "S3_SECRET_KEY",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: r.opniCluster.Status.Auth.S3SecretKey,
			},
		}, corev1.EnvVar{
			Name:  "S3_ENDPOINT",
			Value: r.opniCluster.Status.Auth.S3Endpoint,
		})
	}
	return envVars
}

func deploymentState(enabled *bool) reconciler.DesiredState {
	if enabled == nil || *enabled {
		return reconciler.StatePresent
	}
	return reconciler.StateAbsent
}

func (r *Reconciler) nulogHyperparameters() (runtime.Object, reconciler.DesiredState, error) {
	var data map[string]intstr.IntOrString
	if len(r.opniCluster.Spec.NulogHyperparameters) > 0 {
		data = r.opniCluster.Spec.NulogHyperparameters
	} else {
		data = map[string]intstr.IntOrString{
			"modelThreshold": intstr.FromString("0.5"),
			"minLogTokens":   intstr.FromInt(5),
		}
	}
	cm, err := hyperparameters.GenerateHyperparametersConfigMap("nulog", r.opniCluster.Namespace, data)
	if err != nil {
		return nil, nil, err
	}
	cm.Labels["opni-service"] = "nulog"
	err = ctrl.SetControllerReference(r.opniCluster, &cm, r.client.Scheme())
	return &cm, reconciler.StatePresent, err
}

func (r *Reconciler) inferenceDeployment() (runtime.Object, reconciler.DesiredState, error) {
	deployment := r.genericDeployment(v1beta1.InferenceService)
	s3EnvVars := r.s3EnvVars()
	deployment.Spec.Template.Spec.Containers[0].Env = append(deployment.Spec.Template.Spec.Containers[0].Env, s3EnvVars...)
	if r.opniCluster.Spec.S3.NulogS3Bucket != "" {
		deployment.Spec.Template.Spec.Containers[0].Env = append(deployment.Spec.Template.Spec.Containers[0].Env,
			corev1.EnvVar{
				Name:  "S3_BUCKET",
				Value: r.opniCluster.Spec.S3.NulogS3Bucket,
			})
	}
	insertHyperparametersVolume(deployment, "nulog")
	return deployment, deploymentState(r.opniCluster.Spec.Services.Inference.Enabled), nil
}

func (r *Reconciler) drainDeployment() (runtime.Object, reconciler.DesiredState, error) {
	deployment := r.genericDeployment(v1beta1.DrainService)
	// temporary
	deployment.Spec.Template.Spec.Containers[0].Env = append(deployment.Spec.Template.Spec.Containers[0].Env,
		corev1.EnvVar{
			Name:  "FAIL_KEYWORDS",
			Value: "fail,error,missing,unable",
		})
	s3EnvVars := r.s3EnvVars()
	deployment.Spec.Template.Spec.Containers[0].Env = append(deployment.Spec.Template.Spec.Containers[0].Env, s3EnvVars...)
	if r.opniCluster.Spec.S3.DrainS3Bucket != "" {
		deployment.Spec.Template.Spec.Containers[0].Env = append(deployment.Spec.Template.Spec.Containers[0].Env,
			corev1.EnvVar{
				Name:  "S3_BUCKET",
				Value: r.opniCluster.Spec.S3.DrainS3Bucket,
			})
	}
	return deployment, deploymentState(r.opniCluster.Spec.Services.Drain.Enabled), nil
}

func (r *Reconciler) payloadReceiverDeployment() (runtime.Object, reconciler.DesiredState, error) {
	deployment := r.genericDeployment(v1beta1.PayloadReceiverService)
	return deployment, deploymentState(r.opniCluster.Spec.Services.PayloadReceiver.Enabled), nil
}

func (r *Reconciler) payloadReceiverService() (runtime.Object, reconciler.DesiredState, error) {
	labels := r.serviceLabels(v1beta1.PayloadReceiverService)
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      labels[resources.AppNameLabel],
			Namespace: r.opniCluster.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				{
					APIVersion: r.opniCluster.APIVersion,
					Kind:       r.opniCluster.Kind,
					Name:       r.opniCluster.Name,
					UID:        r.opniCluster.UID,
				},
			},
			Labels: labels,
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
			Ports: []corev1.ServicePort{
				{
					Port: 80,
				},
			},
		},
	}
	return service, deploymentState(r.opniCluster.Spec.Services.PayloadReceiver.Enabled), nil
}

func (r *Reconciler) preprocessingDeployment() (runtime.Object, reconciler.DesiredState, error) {
	deployment := r.genericDeployment(v1beta1.PreprocessingService)
	return deployment, deploymentState(r.opniCluster.Spec.Services.Preprocessing.Enabled), nil
}

func (r *Reconciler) gpuCtrlDeployment() (runtime.Object, reconciler.DesiredState, error) {
	deployment := r.genericDeployment(v1beta1.GPUControllerService)
	dataVolume := corev1.Volume{
		Name: "data",
		VolumeSource: corev1.VolumeSource{
			EmptyDir: &corev1.EmptyDirVolumeSource{},
		},
	}
	dataVolumeMount := corev1.VolumeMount{
		Name:      "data",
		MountPath: "/var/opni-data",
	}
	deployment.Spec.Template.Spec.RuntimeClassName = r.opniCluster.Spec.Services.GPUController.RuntimeClass
	if features.DefaultMutableFeatureGate.Enabled(features.GPUOperator) && r.opniCluster.Spec.Services.GPUController.RuntimeClass == nil {
		deployment.Spec.Template.Spec.RuntimeClassName = pointer.String("nvidia")
	}
	deployment.Spec.Strategy.Type = appsv1.RecreateDeploymentStrategyType
	deployment.Spec.Template.Spec.Containers = append(deployment.Spec.Template.Spec.Containers, r.gpuWorkerContainer())
	deployment.Spec.Template.Spec.Volumes = append(deployment.Spec.Template.Spec.Volumes, dataVolume)
	for i, container := range deployment.Spec.Template.Spec.Containers {
		container.VolumeMounts = append(container.VolumeMounts, dataVolumeMount)
		deployment.Spec.Template.Spec.Containers[i] = container
	}
	insertHyperparametersVolume(deployment, "nulog")
	return deployment, deploymentState(r.opniCluster.Spec.Services.GPUController.Enabled), nil
}

func (r *Reconciler) metricsDeployment() (runtime.Object, reconciler.DesiredState, error) {
	deployment := r.genericDeployment(v1beta1.MetricsService)
	_, err := url.ParseRequestURI(r.opniCluster.Spec.Services.Metrics.PrometheusEndpoint)
	if err != nil && (r.opniCluster.Spec.Services.Metrics.Enabled == nil || *r.opniCluster.Spec.Services.Metrics.Enabled) {
		return deployment, deploymentState(r.opniCluster.Spec.Services.Metrics.Enabled), errors.New("prometheus endpoint is not a valid URL")
	}
	deployment.Spec.Template.Spec.Containers[0].Env = append(deployment.Spec.Template.Spec.Containers[0].Env, corev1.EnvVar{
		Name:  "PROMETHEUS_ENDPOINT",
		Value: r.opniCluster.Spec.Services.Metrics.PrometheusEndpoint,
	})
	return deployment, deploymentState(r.opniCluster.Spec.Services.Metrics.Enabled), nil
}

func insertHyperparametersVolume(deployment *appsv1.Deployment, modelName string) {
	volume := corev1.Volume{
		Name: "hyperparameters",
		VolumeSource: corev1.VolumeSource{
			ConfigMap: &corev1.ConfigMapVolumeSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: fmt.Sprintf("opni-%s-hyperparameters", modelName),
				},
				Items: []corev1.KeyToPath{
					{
						Key:  "hyperparameters.json",
						Path: "hyperparameters.json",
					},
				},
			},
		},
	}
	deployment.Spec.Template.Spec.Volumes = append(deployment.Spec.Template.Spec.Volumes, volume)
	for i, container := range deployment.Spec.Template.Spec.Containers {
		container.VolumeMounts = append(container.VolumeMounts, corev1.VolumeMount{
			Name:      "hyperparameters",
			MountPath: "/etc/opni/hyperparameters.json",
			SubPath:   "hyperparameters.json",
			ReadOnly:  true,
		})
		deployment.Spec.Template.Spec.Containers[i] = container
	}
}
