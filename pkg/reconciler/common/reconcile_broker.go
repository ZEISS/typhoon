package common

import (
	"context"
	"strconv"

	"go.uber.org/zap"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	appsv1listers "k8s.io/client-go/listers/apps/v1"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"knative.dev/eventing/pkg/apis/duck"
	k8sclient "knative.dev/pkg/client/injection/kube/client"
	"knative.dev/pkg/logging"
	pkgreconciler "knative.dev/pkg/reconciler"

	eventingv1alpha1 "github.com/zeiss/typhoon/pkg/apis/eventing/v1alpha1"
	"github.com/zeiss/typhoon/pkg/reconciler/resource"
	"github.com/zeiss/typhoon/pkg/reconciler/semantic"
)

const (
	brokerResourceSuffix           = "broker"
	brokerDeploymentComponentLabel = "broker-deployment"

	// container ports must be >1024 to be able to bind them
	// in unprivileged environments.
	brokerContainerPort = 8080

	defaultBrokerServicePort = 80
	metricsServicePort       = 9090
)

type BrokerReconciler interface {
	Reconcile(ctx context.Context, rb eventingv1alpha1.ReconcilableBroker, sa *corev1.ServiceAccount, secret *corev1.Secret, configMap *corev1.ConfigMap, do ...resource.ObjectOption) (*appsv1.Deployment, *corev1.Service, error)
}

type brokerReconciler struct {
	client           kubernetes.Interface
	deploymentLister appsv1listers.DeploymentLister
	serviceLister    corev1listers.ServiceLister
	endpointsLister  corev1listers.EndpointsLister
	image            string
	// TODO remove when using releases
	pullPolicy corev1.PullPolicy
}

func NewBrokerReconciler(ctx context.Context,
	deploymentLister appsv1listers.DeploymentLister,
	serviceLister corev1listers.ServiceLister,
	endpointsLister corev1listers.EndpointsLister,
	image string,
	pullPolicy corev1.PullPolicy,
) BrokerReconciler {
	return &brokerReconciler{
		client:           k8sclient.Get(ctx),
		deploymentLister: deploymentLister,
		serviceLister:    serviceLister,
		endpointsLister:  endpointsLister,
		image:            image,
		pullPolicy:       pullPolicy,
	}
}

func (r *brokerReconciler) Reconcile(ctx context.Context, rb eventingv1alpha1.ReconcilableBroker, sa *corev1.ServiceAccount, secret *corev1.Secret, configMap *corev1.ConfigMap, deploymentOptions ...resource.ObjectOption) (*appsv1.Deployment, *corev1.Service, error) {
	d, err := r.reconcileDeployment(ctx, rb, sa, secret, configMap, deploymentOptions)
	if err != nil {
		return nil, nil, err
	}

	svc, err := r.reconcileService(ctx, rb)
	if err != nil {
		return d, nil, err
	}

	_, err = r.reconcileEndpoints(ctx, svc, rb)
	if err != nil {
		return d, nil, err
	}

	return d, svc, nil
}

func buildBrokerDeployment(rb eventingv1alpha1.ReconcilableBroker, sa *corev1.ServiceAccount, secret *corev1.Secret, cm *corev1.ConfigMap, image string, pullPolicy corev1.PullPolicy, extraOptions ...resource.ObjectOption) *appsv1.Deployment {
	meta := rb.GetObjectMeta()
	ns, name := meta.GetNamespace(), meta.GetName()
	bs := rb.GetReconcilableBrokerSpec()

	copts := []resource.ObjectOption{
		resource.ContainerAddArgs("start"),
		resource.ContainerAddEnvFromValue("PORT", strconv.Itoa(int(brokerContainerPort))),
		resource.ContainerAddEnvFromFieldRef("BROKER_NAME", "metadata.name"),
		resource.ContainerAddEnvFromFieldRef("KUBERNETES_NAMESPACE", "metadata.namespace"),
		resource.ContainerAddEnvFromValue("KUBERNETES_BROKER_CONFIG_SECRET_NAME", secret.Name),
		resource.ContainerAddEnvFromValue("KUBERNETES_BROKER_CONFIG_SECRET_KEY", ConfigSecretKey),
		resource.ContainerAddEnvFromValue("KUBERNETES_STATUS_CONFIGMAP_NAME", cm.Name),
		resource.ContainerWithImagePullPolicy(pullPolicy),
		resource.ContainerAddPort("httpce", brokerContainerPort),
		resource.ContainerAddPort("metrics", metricsServicePort),
	}

	if bs.Observability != nil && bs.Observability.ValueFromConfigMap != "" {
		copts = append(copts, resource.ContainerAddEnvFromValue("KUBERNETES_OBSERVABILITY_CONFIGMAP_NAME", bs.Observability.ValueFromConfigMap))
	}

	dn := name + "-" + rb.GetOwnedObjectsSuffix() + "-" + brokerResourceSuffix
	d := resource.NewDeployment(ns, dn,
		resource.DeploymentWithMetaOptions(
			resource.MetaAddLabel(resource.AppNameLabel, AppAnnotationValue(rb)),
			resource.MetaAddLabel(resource.AppComponentLabel, brokerDeploymentComponentLabel),
			resource.MetaAddLabel(resource.AppPartOfLabel, resource.PartOf),
			resource.MetaAddLabel(resource.AppManagedByLabel, resource.ManagedBy),
			resource.MetaAddLabel(resource.AppInstanceLabel, dn),
			resource.MetaAddOwner(meta, rb.GetGroupVersionKind())),
		resource.DeploymentAddSelectorForTemplate(resource.AppComponentLabel, brokerDeploymentComponentLabel),
		resource.DeploymentAddSelectorForTemplate(resource.AppInstanceLabel, dn),
		resource.DeploymentSetReplicas(1),
		resource.DeploymentWithTemplateSpecOptions(
			// Needed for prometheus PodMonitor.
			resource.PodTemplateSpecWithMetaOptions(
				resource.MetaAddLabel(resource.AppPartOfLabel, resource.PartOf),
				resource.MetaAddLabel(resource.AppManagedByLabel, resource.ManagedBy),
			),
			resource.PodTemplateSpecWithPodSpecOptions(
				resource.PodSpecWithServiceAccountName(sa.Name),
				resource.PodSpecAddContainer(
					resource.NewContainer("broker", image, copts...)))))

	if len(extraOptions) != 0 {
		for _, o := range extraOptions {
			o(d)
		}
	}

	return d
}

func (r *brokerReconciler) reconcileDeployment(
	ctx context.Context,
	rb eventingv1alpha1.ReconcilableBroker,
	sa *corev1.ServiceAccount,
	secret *corev1.Secret,
	configMap *corev1.ConfigMap,
	deploymentOptions []resource.ObjectOption,
) (*appsv1.Deployment, error) {
	desired := buildBrokerDeployment(rb, sa, secret, configMap, r.image, r.pullPolicy, deploymentOptions...)
	current, err := r.deploymentLister.Deployments(desired.Namespace).Get(desired.Name)

	switch {
	case err == nil:
		// Compare current object with desired, update if needed.
		if !semantic.Semantic.DeepEqual(desired, current) {
			desired.Status = current.Status
			desired.ResourceVersion = current.ResourceVersion

			current, err = r.client.AppsV1().Deployments(desired.Namespace).Update(ctx, desired, metav1.UpdateOptions{})
			if err != nil {
				fullname := types.NamespacedName{Namespace: desired.Namespace, Name: desired.Name}
				logging.FromContext(ctx).Error("Unable to update broker deployment", zap.String("deployment", fullname.String()), zap.Error(err))
				rb.GetReconcilableBrokerStatus().MarkBrokerDeploymentFailed(ReasonFailedDeploymentUpdate, "Failed to update broker deployment")

				return nil, pkgreconciler.NewEvent(corev1.EventTypeWarning, ReasonFailedDeploymentUpdate,
					"Failed to get broker deployment %s: %w", fullname, err)
			}
		}

	case !apierrs.IsNotFound(err):
		// An error occurred retrieving current deployment.
		fullname := types.NamespacedName{Namespace: desired.Namespace, Name: desired.Name}
		logging.FromContext(ctx).Error("Unable to get broker deployment", zap.String("deployment", fullname.String()), zap.Error(err))
		rb.GetReconcilableBrokerStatus().MarkBrokerDeploymentFailed(ReasonFailedDeploymentGet, "Failed to get broker deployment")

		return nil, pkgreconciler.NewEvent(corev1.EventTypeWarning, ReasonFailedDeploymentGet,
			"Failed to get broker deployment %s: %w", fullname, err)

	default:
		// The deployment has not been found, create it.
		current, err = r.client.AppsV1().Deployments(desired.Namespace).Create(ctx, desired, metav1.CreateOptions{})
		if err != nil {
			fullname := types.NamespacedName{Namespace: desired.Namespace, Name: desired.Name}
			logging.FromContext(ctx).Error("Unable to create broker deployment", zap.String("deployment", fullname.String()), zap.Error(err))
			rb.GetReconcilableBrokerStatus().MarkBrokerDeploymentFailed(ReasonFailedDeploymentCreate, "Failed to create broker deployment")

			return nil, pkgreconciler.NewEvent(corev1.EventTypeWarning, ReasonFailedDeploymentCreate,
				"Failed to create broker deployment %s: %w", fullname, err)
		}
	}

	// Update status based on deployment
	rb.GetReconcilableBrokerStatus().PropagateBrokerDeploymentAvailability(ctx, &current.Status)

	return current, nil
}

func buildBrokerService(rb eventingv1alpha1.ReconcilableBroker) *corev1.Service {
	meta := rb.GetObjectMeta()
	ns, name := meta.GetNamespace(), meta.GetName()
	bs := rb.GetReconcilableBrokerSpec()

	brokerPort := defaultBrokerServicePort
	if bs.Port != nil {
		brokerPort = *bs.Port
	}

	sn := name + "-" + rb.GetOwnedObjectsSuffix() + "-" + brokerResourceSuffix
	return resource.NewService(ns, sn,
		resource.ServiceWithMetaOptions(
			resource.MetaAddLabel(resource.AppNameLabel, AppAnnotationValue(rb)),
			resource.MetaAddLabel(resource.AppComponentLabel, "broker-service"),
			resource.MetaAddLabel(resource.AppPartOfLabel, resource.PartOf),
			resource.MetaAddLabel(resource.AppManagedByLabel, resource.ManagedBy),
			resource.MetaAddLabel(resource.AppInstanceLabel, sn),
			resource.MetaAddOwner(meta, rb.GetGroupVersionKind())),
		resource.ServiceSetType(corev1.ServiceTypeClusterIP),
		resource.ServiceAddSelectorLabel(resource.AppComponentLabel, brokerDeploymentComponentLabel),
		resource.ServiceAddSelectorLabel(resource.AppInstanceLabel, sn),
		resource.ServiceAddPort("httpce", int32(brokerPort), brokerContainerPort))
}

func (r *brokerReconciler) reconcileService(ctx context.Context, rb eventingv1alpha1.ReconcilableBroker) (*corev1.Service, error) {
	desired := buildBrokerService(rb)
	current, err := r.serviceLister.Services(desired.Namespace).Get(desired.Name)

	switch {
	case err == nil:
		// Set Status
		// Compare current object with desired, update if needed.
		if !semantic.Semantic.DeepEqual(desired, current) {
			desired.Status = current.Status
			desired.ResourceVersion = current.ResourceVersion

			current, err = r.client.CoreV1().Services(desired.Namespace).Update(ctx, desired, metav1.UpdateOptions{})
			if err != nil {
				fullname := types.NamespacedName{Namespace: desired.Namespace, Name: desired.Name}
				logging.FromContext(ctx).Error("Unable to update broker service", zap.String("service", fullname.String()), zap.Error(err))
				rb.GetReconcilableBrokerStatus().MarkBrokerServiceFailed(ReasonFailedServiceUpdate, "Failed to update broker service")

				return nil, pkgreconciler.NewEvent(corev1.EventTypeWarning, ReasonFailedServiceUpdate,
					"Failed to get broker service %s: %w", fullname, err)
			}
		}

	case !apierrs.IsNotFound(err):
		// An error occurred retrieving current object.
		fullname := types.NamespacedName{Namespace: desired.Namespace, Name: desired.Name}
		logging.FromContext(ctx).Error("Unable to get the service", zap.String("service", fullname.String()), zap.Error(err))
		rb.GetReconcilableBrokerStatus().MarkBrokerServiceFailed(ReasonFailedServiceGet, "Failed to get broker service")

		return nil, pkgreconciler.NewEvent(corev1.EventTypeWarning, ReasonFailedServiceGet,
			"Failed to get broker service %s: %w", fullname, err)

	default:
		// The object has not been found, create it.
		current, err = r.client.CoreV1().Services(desired.Namespace).Create(ctx, desired, metav1.CreateOptions{})
		if err != nil {
			fullname := types.NamespacedName{Namespace: desired.Namespace, Name: desired.Name}
			logging.FromContext(ctx).Error("Unable to create the service", zap.String("service", fullname.String()), zap.Error(err))
			rb.GetReconcilableBrokerStatus().MarkBrokerServiceFailed(ReasonFailedServiceCreate, "Failed to create broker service")

			return nil, pkgreconciler.NewEvent(corev1.EventTypeWarning, ReasonFailedServiceCreate,
				"Failed to create broker service %s: %w", fullname, err)
		}
	}

	// Service exists and is up to date.
	rb.GetReconcilableBrokerStatus().MarkBrokerServiceReady()

	return current, nil
}

func (r *brokerReconciler) reconcileEndpoints(ctx context.Context, service *corev1.Service, rb eventingv1alpha1.ReconcilableBroker) (*corev1.Endpoints, error) {
	ep, err := r.endpointsLister.Endpoints(service.Namespace).Get(service.Name)
	switch {
	case err == nil:
		if duck.EndpointsAreAvailable(ep) {
			rb.GetReconcilableBrokerStatus().MarkBrokerEndpointsTrue()
			return ep, nil
		}

		rb.GetReconcilableBrokerStatus().MarkBrokerEndpointsFailed(ReasonUnavailableEndpoints, "Endpoints for broker service are not available")
		return nil, pkgreconciler.NewEvent(corev1.EventTypeWarning, ReasonUnavailableEndpoints,
			"Endpoints for broker service %q are not available",
			types.NamespacedName{Namespace: ep.Namespace, Name: ep.Name})

	case apierrs.IsNotFound(err):
		rb.GetReconcilableBrokerStatus().MarkBrokerEndpointsFailed(ReasonUnavailableEndpoints, "Endpoints for broker service do not exist")
		return nil, pkgreconciler.NewEvent(corev1.EventTypeWarning, ReasonUnavailableEndpoints,
			"Endpoints for broker service %q do not exist",
			types.NamespacedName{Namespace: service.Namespace, Name: service.Name})
	}

	fullname := types.NamespacedName{Namespace: service.Namespace, Name: service.Name}
	rb.GetReconcilableBrokerStatus().MarkBrokerEndpointsUnknown(ReasonFailedEndpointsGet, "Could not retrieve endpoints for broker service")
	logging.FromContext(ctx).Error("Unable to get the broker service endpoints", zap.String("endpoint", fullname.String()), zap.Error(err))
	return nil, pkgreconciler.NewEvent(corev1.EventTypeWarning, ReasonFailedEndpointsGet,
		"Failed to get broker service ednpoints %s: %w", fullname, err)
}
