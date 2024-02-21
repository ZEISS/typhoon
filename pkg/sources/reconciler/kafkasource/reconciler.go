package kafkasource

import (
	"context"

	"knative.dev/pkg/reconciler"

	commonv1alpha1 "github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	"github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	reconcilerv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/injection/reconciler/sources/v1alpha1/kafkasource"
	listersv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/listers/sources/v1alpha1"
	common "github.com/zeiss/typhoon/pkg/reconciler"
)

// Reconciler implements controller.Reconciler for the event source type.
type Reconciler struct {
	base       common.GenericDeploymentReconciler[*v1alpha1.KafkaSource, listersv1alpha1.KafkaSourceNamespaceLister]
	adapterCfg *adapterConfig
}

// Check that our Reconciler implements Interface
var _ reconcilerv1alpha1.Interface = (*Reconciler)(nil)

// ReconcileKind implements Interface.ReconcileKind.
func (r *Reconciler) ReconcileKind(ctx context.Context, src *v1alpha1.KafkaSource) reconciler.Event {
	// inject source into context for usage in reconciliation logic
	ctx = commonv1alpha1.WithReconcilable(ctx, src)
	return r.base.ReconcileAdapter(ctx, r)
}
