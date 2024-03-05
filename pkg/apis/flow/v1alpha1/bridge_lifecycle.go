package v1alpha1

import (
	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*Bridge) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("Bridge")
}

// GetConditionSet implements duckv1.KRShaped.
func (t *Bridge) GetConditionSet() apis.ConditionSet {
	return v1alpha1.DefaultConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (t *Bridge) GetStatus() *duckv1.Status {
	return &t.Status.Status
}

// GetStatusManager implements Reconcilable.
func (t *Bridge) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: t.GetConditionSet(),
		Status:       &t.Status,
	}
}
