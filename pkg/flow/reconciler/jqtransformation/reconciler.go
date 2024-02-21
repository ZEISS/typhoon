package jqtransformation

import (
	"context"

	"knative.dev/pkg/reconciler"

	commonv1alpha1 "github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	"github.com/zeiss/typhoon/pkg/apis/flow/v1alpha1"
	reconcilerv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/injection/reconciler/flow/v1alpha1/jqtransformation"
	listersv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/listers/flow/v1alpha1"
	common "github.com/zeiss/typhoon/pkg/reconciler"
)

// Reconciler implements controller.Reconciler for the event target type.
type Reconciler struct {
	base       common.GenericServiceReconciler[*v1alpha1.JQTransformation, listersv1alpha1.JQTransformationNamespaceLister]
	adapterCfg *adapterConfig
}

// Check that our Reconciler implements Interface
var _ reconcilerv1alpha1.Interface = (*Reconciler)(nil)

// ReconcileKind implements Interface.ReconcileKind.
func (r *Reconciler) ReconcileKind(ctx context.Context, trg *v1alpha1.JQTransformation) reconciler.Event {
	// inject target into context for usage in reconciliation logic
	ctx = commonv1alpha1.WithReconcilable(ctx, trg)

	return r.base.ReconcileAdapter(ctx, r)
}
