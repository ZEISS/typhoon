package filter

import (
	"context"

	"knative.dev/pkg/reconciler"

	commonv1alpha1 "github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	"github.com/zeiss/typhoon/pkg/apis/routing/v1alpha1"
	reconcilerv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/injection/reconciler/routing/v1alpha1/filter"
	listersv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/listers/routing/v1alpha1"
	common "github.com/zeiss/typhoon/pkg/reconciler"
)

// Reconciler implements addressableservicereconciler.Interface for
// AddressableService resources.
type Reconciler struct {
	base       common.GenericServiceReconciler[*v1alpha1.Filter, listersv1alpha1.FilterNamespaceLister]
	adapterCfg *adapterConfig
}

// Check that our Reconciler implements Interface
var _ reconcilerv1alpha1.Interface = (*Reconciler)(nil)

// ReconcileKind implements Interface.ReconcileKind.
func (r *Reconciler) ReconcileKind(ctx context.Context, o *v1alpha1.Filter) reconciler.Event {
	// inject component instance into context for usage in reconciliation logic
	ctx = commonv1alpha1.WithReconcilable(ctx, o)

	return r.base.ReconcileAdapter(ctx, r)
}
