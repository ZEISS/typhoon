package common

import (
	"context"

	"go.uber.org/zap"

	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	corev1listers "k8s.io/client-go/listers/core/v1"
	rbacv1listers "k8s.io/client-go/listers/rbac/v1"
	k8sclient "knative.dev/pkg/client/injection/kube/client"
	"knative.dev/pkg/logging"
	pkgreconciler "knative.dev/pkg/reconciler"

	eventingv1alpha1 "github.com/zeiss/typhoon/pkg/apis/eventing/v1alpha1"
	"github.com/zeiss/typhoon/pkg/reconciler/resource"
)

const (
	// Broker ClusterRole that was created as part of Typhoon core installation.
	BrokerDeploymentRole         = "typhoon-broker"
	serviceAccountResourceSuffix = "broker"
	roleBindingResourceSuffix    = "broker"
)

type ServiceAccountReconciler interface {
	Reconcile(ctx context.Context, rb eventingv1alpha1.ReconcilableBroker) (*corev1.ServiceAccount, *rbacv1.RoleBinding, error)
}

type serviceAccountReconciler struct {
	client               kubernetes.Interface
	serviceAccountLister corev1listers.ServiceAccountLister
	roleBindingLister    rbacv1listers.RoleBindingLister
}

var _ ServiceAccountReconciler = (*serviceAccountReconciler)(nil)

func NewServiceAccountReconciler(ctx context.Context, serviceAccountLister corev1listers.ServiceAccountLister, roleBindingLister rbacv1listers.RoleBindingLister) ServiceAccountReconciler {
	return &serviceAccountReconciler{
		client:               k8sclient.Get(ctx),
		serviceAccountLister: serviceAccountLister,
		roleBindingLister:    roleBindingLister,
	}
}

func (r *serviceAccountReconciler) Reconcile(ctx context.Context, rb eventingv1alpha1.ReconcilableBroker) (*corev1.ServiceAccount, *rbacv1.RoleBinding, error) {
	sa, err := r.reconcileServiceAccount(ctx, rb)
	if err != nil {
		return nil, nil, err
	}

	roleb, err := r.reconcileRoleBinding(ctx, rb, sa)
	if err != nil {
		return nil, nil, err
	}

	return sa, roleb, nil
}

func buildBrokerServiceAccount(rb eventingv1alpha1.ReconcilableBroker) *corev1.ServiceAccount {
	meta := rb.GetObjectMeta()
	ns, name := meta.GetNamespace(), meta.GetName()+"-"+rb.GetOwnedObjectsSuffix()+"-"+serviceAccountResourceSuffix

	return resource.NewServiceAccount(ns, name,
		resource.ServiceAccountWithMetaOptions(
			resource.MetaAddLabel(resource.AppNameLabel, AppAnnotationValue(rb)),
			resource.MetaAddLabel(resource.AppComponentLabel, "broker-serviceaccount"),
			resource.MetaAddLabel(resource.AppPartOfLabel, resource.PartOf),
			resource.MetaAddLabel(resource.AppManagedByLabel, resource.ManagedBy),
			resource.MetaAddLabel(resource.AppInstanceLabel, name),
			resource.MetaAddOwner(meta, rb.GetGroupVersionKind())))
}

func (r *serviceAccountReconciler) reconcileServiceAccount(ctx context.Context, rb eventingv1alpha1.ReconcilableBroker) (*corev1.ServiceAccount, error) {
	desired := buildBrokerServiceAccount(rb)
	current, err := r.serviceAccountLister.ServiceAccounts(desired.Namespace).Get(desired.Name)

	switch {
	case err == nil:
		// TODO compare

	case !apierrs.IsNotFound(err):
		// An error occurred retrieving current object.
		fullname := types.NamespacedName{Namespace: desired.Namespace, Name: desired.Name}
		logging.FromContext(ctx).Error("Unable to get broker ServiceAccount", zap.String("serviceAccount", fullname.String()), zap.Error(err))
		rb.GetReconcilableBrokerStatus().MarkBrokerServiceAccountFailed(ReasonFailedServiceAccountGet, "Failed to get broker ServiceAccount")

		return nil, pkgreconciler.NewEvent(corev1.EventTypeWarning, ReasonFailedServiceAccountGet,
			"Failed to get broker ServiceAccount %s: %w", fullname, err)

	default:
		// The ServiceAccount has not been found, create it.
		current, err = r.client.CoreV1().ServiceAccounts(desired.Namespace).Create(ctx, desired, metav1.CreateOptions{})
		if err != nil {
			fullname := types.NamespacedName{Namespace: desired.Namespace, Name: desired.Name}
			logging.FromContext(ctx).Error("Unable to create broker ServiceAccount", zap.String("serviceAccount", fullname.String()), zap.Error(err))
			rb.GetReconcilableBrokerStatus().MarkBrokerServiceAccountFailed(ReasonFailedServiceAccountCreate, "Failed to create broker ServiceAccount")

			return nil, pkgreconciler.NewEvent(corev1.EventTypeWarning, ReasonFailedServiceAccountCreate,
				"Failed to create broker ServiceAccount %s: %w", fullname, err)
		}
	}

	// Update status
	rb.GetReconcilableBrokerStatus().MarkBrokerServiceAccountReady()

	return current, nil
}

func buildBrokerRoleBinding(rb eventingv1alpha1.ReconcilableBroker, sa *corev1.ServiceAccount) *rbacv1.RoleBinding {
	meta := rb.GetObjectMeta()
	ns, name := meta.GetNamespace(), meta.GetName()+"-"+rb.GetOwnedObjectsSuffix()+"-"+roleBindingResourceSuffix

	return resource.NewRoleBinding(ns, name, BrokerDeploymentRole, sa.Name,
		resource.RoleBindingWithMetaOptions(
			resource.MetaAddLabel(resource.AppNameLabel, AppAnnotationValue(rb)),
			resource.MetaAddLabel(resource.AppComponentLabel, "broker-rolebinding"),
			resource.MetaAddLabel(resource.AppPartOfLabel, resource.PartOf),
			resource.MetaAddLabel(resource.AppManagedByLabel, resource.ManagedBy),
			resource.MetaAddLabel(resource.AppInstanceLabel, name),
			resource.MetaAddOwner(meta, rb.GetGroupVersionKind())))
}

func (r *serviceAccountReconciler) reconcileRoleBinding(ctx context.Context, rb eventingv1alpha1.ReconcilableBroker, sa *corev1.ServiceAccount) (*rbacv1.RoleBinding, error) {
	desired := buildBrokerRoleBinding(rb, sa)
	current, err := r.roleBindingLister.RoleBindings(desired.Namespace).Get(desired.Name)

	switch {
	case err == nil:
		// TODO compare

	case !apierrs.IsNotFound(err):
		// An error occurred retrieving current object.
		fullname := types.NamespacedName{Namespace: desired.Namespace, Name: desired.Name}
		logging.FromContext(ctx).Error("Unable to get broker RoleBinding", zap.String("roleBinding", fullname.String()), zap.Error(err))
		rb.GetReconcilableBrokerStatus().MarkBrokerRoleBindingFailed(ReasonFailedRoleBindingGet, "Failed to get broker RoleBinding")

		return nil, pkgreconciler.NewEvent(corev1.EventTypeWarning, ReasonFailedRoleBindingGet,
			"Failed to get broker RoleBinding %s: %w", fullname, err)

	default:
		// The RoleBinding has not been found, create it.
		current, err = r.client.RbacV1().RoleBindings(desired.Namespace).Create(ctx, desired, metav1.CreateOptions{})
		if err != nil {
			fullname := types.NamespacedName{Namespace: desired.Namespace, Name: desired.Name}
			logging.FromContext(ctx).Error("Unable to create broker RoleBinding", zap.String("roleBinding", fullname.String()), zap.Error(err))
			rb.GetReconcilableBrokerStatus().MarkBrokerRoleBindingFailed(ReasonFailedRoleBindingCreate, "Failed to create broker RoleBinding")

			return nil, pkgreconciler.NewEvent(corev1.EventTypeWarning, ReasonFailedRoleBindingCreate,
				"Failed to create broker RoleBinding %s: %w", fullname, err)
		}
	}

	// Update status
	rb.GetReconcilableBrokerStatus().MarkBrokerRoleBindingReady()

	return current, nil
}
