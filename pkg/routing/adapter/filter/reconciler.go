package filter

import (
	"context"

	"knative.dev/pkg/reconciler"

	"github.com/zeiss/typhoon/pkg/apis/routing/v1alpha1"
	reconcilerv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/injection/reconciler/routing/v1alpha1/filter"
)

// Reconciler implements controller.Reconciler for the event source type.
type Reconciler struct{}

// Check the interfaces Reconciler should implement.
var _ reconcilerv1alpha1.Interface = (*Reconciler)(nil)

// ReconcileKind implements reconcilerv1alpha1.Interface.
func (r *Reconciler) ReconcileKind(ctx context.Context, s *v1alpha1.Filter) reconciler.Event {
	return nil
}
