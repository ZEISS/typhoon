package common

import (
	"context"

	eventingv1alpha1 "github.com/zeiss/typhoon/pkg/apis/eventing/v1alpha1"
	"github.com/zeiss/typhoon/pkg/reconciler/resource"
	corev1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	corev1listers "k8s.io/client-go/listers/core/v1"
	k8sclient "knative.dev/pkg/client/injection/kube/client"
	pkgreconciler "knative.dev/pkg/reconciler"
	// import the other required packages
)

type ConfigMapReconciler interface {
	Reconcile(ctx context.Context, rb eventingv1alpha1.ReconcilableBroker) (*corev1.ConfigMap, error)
}

type configMapReconciler struct {
	client          kubernetes.Interface
	configMapLister corev1listers.ConfigMapLister
}

var _ ConfigMapReconciler = (*configMapReconciler)(nil)

func NewConfigMapReconciler(ctx context.Context, configMapLister corev1listers.ConfigMapLister) ConfigMapReconciler {
	return &configMapReconciler{
		client:          k8sclient.Get(ctx),
		configMapLister: configMapLister,
	}
}

func (r *configMapReconciler) Reconcile(ctx context.Context, rb eventingv1alpha1.ReconcilableBroker) (*corev1.ConfigMap, error) {
	meta := rb.GetObjectMeta()
	ns := meta.GetNamespace()

	configMapName := GetBrokerConfigMapName(rb)

	desired := resource.NewConfigMap(ns, configMapName,
		resource.ConfigMapWithMetaOptions(
			resource.MetaAddLabel(resource.AppNameLabel, AppAnnotationValue(rb)),
			resource.MetaAddLabel(resource.AppComponentLabel, "broker-status"),
			resource.MetaAddLabel(resource.AppPartOfLabel, resource.PartOf),
			resource.MetaAddLabel(resource.AppManagedByLabel, resource.ManagedBy),
			resource.MetaAddLabel(resource.AppInstanceLabel, configMapName),
			resource.MetaAddOwner(meta, rb.GetGroupVersionKind())),
	)

	_, err := r.configMapLister.ConfigMaps(desired.Namespace).Get(desired.Name)
	switch {
	case err == nil:
		// We only require the ConfigMap to exist, no action needed.

	case apierrs.IsNotFound(err):
		// The configMap has not been found, create it.
		_, err = r.client.CoreV1().ConfigMaps(desired.Namespace).Create(ctx, desired, metav1.CreateOptions{})
		if err != nil {
			rb.GetReconcilableBrokerStatus().MarkStatusConfigFailed(ReasonStatusConfigMapCreateFailed, "Failed to create configMap for status reporting")
			return nil, pkgreconciler.NewEvent(corev1.EventTypeWarning, ReasonStatusConfigMapCreateFailed,
				"Failed to create configMap for status reporting %s: %w", desired.Name, err)
		}

	default:
		rb.GetReconcilableBrokerStatus().MarkStatusConfigFailed(ReasonStatusConfigMapGetFailed, "Failed to get configMap for status reporting")
		return nil, pkgreconciler.NewEvent(corev1.EventTypeWarning, ReasonStatusConfigMapGetFailed,
			"Failed to get configMap for status reporting %s: %w", desired.Name, err)
	}

	rb.GetReconcilableBrokerStatus().MarkStatusConfigReady()

	return desired, nil
}
