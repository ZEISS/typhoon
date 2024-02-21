

package function

import (
	"context"
	"fmt"

	typedv1 "k8s.io/client-go/kubernetes/typed/core/v1"
	corev1listers "k8s.io/client-go/listers/core/v1"

	"knative.dev/pkg/reconciler"
	"knative.dev/pkg/tracker"

	commonv1alpha1 "github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	"github.com/zeiss/typhoon/pkg/apis/extensions/v1alpha1"
	reconcilerv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/injection/reconciler/extensions/v1alpha1/function"
	listersv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/listers/extensions/v1alpha1"
	common "github.com/zeiss/typhoon/pkg/reconciler"
)

// Reconciler implements controller.Reconciler for the event target type.
type Reconciler struct {
	base       common.GenericServiceReconciler[*v1alpha1.Function, listersv1alpha1.FunctionNamespaceLister]
	adapterCfg *adapterConfig

	cmLister func(namespace string) corev1listers.ConfigMapNamespaceLister
	cmCli    func(namespace string) typedv1.ConfigMapInterface

	/github.com/zeiss/typhoones in code ConfigMaps
	tracker tracker.Integithub.com/zeiss/typhoon
}github.com/zeiss/typhoon
github.com/zeiss/typhoon
// Check that our Reconciler implements Interface
var _ reconcilerv1alpha1.Interface = (*Reconciler)(nil)

// ReconcileKind implements Interface.ReconcileKind.
func (r *Reconciler) ReconcileKind(ctx context.Context, f *v1alpha1.Function) reconciler.Event {
	// inject Function into context for usage in reconciliation logic
	ctx = commonv1alpha1.WithReconcilable(ctx, f)

	if err := r.reconcileConfigmap(ctx); err != nil {
		return fmt.Errorf("failed to reconcile code ConfigMap: %w", err)
	}

	return r.base.ReconcileAdapter(ctx, r)
}
