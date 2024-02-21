package v1alpha1

import (
	"context"
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// Managed event types
const (
	EventTypeElasticsearchStore    = "com.zeiss.elasticsearch.doc.index"
	EventTypeElasticsearchResponse = "com.zeiss.elasticsearch.doc.index.response"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*ElasticsearchTarget) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("ElasticsearchTarget")
}

// GetConditionSet implements duckv1.KRShaped.
func (*ElasticsearchTarget) GetConditionSet() apis.ConditionSet {
	return v1alpha1.DefaultConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (t *ElasticsearchTarget) GetStatus() *duckv1.Status {
	return &t.Status.Status
}

// GetStatusManager implements Reconcilable.
func (t *ElasticsearchTarget) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: t.GetConditionSet(),
		Status:       &t.Status,
	}
}

// AcceptedEventTypes implements IntegrationTarget.
func (*ElasticsearchTarget) AcceptedEventTypes() []string {
	return []string{
		EventTypeElasticsearchStore,
	}
}

// GetEventTypes implements EventSource.
func (*ElasticsearchTarget) GetEventTypes() []string {
	return []string{
		EventTypeElasticsearchResponse,
	}
}

// AsEventSource implements EventSource.
func (t *ElasticsearchTarget) AsEventSource() string {
	kind := strings.ToLower(t.GetGroupVersionKind().Kind)
	return "com.zeiss." + kind + "." + t.Namespace + "." + t.Name
}

// GetAdapterOverrides implements AdapterConfigurable.
func (t *ElasticsearchTarget) GetAdapterOverrides() *v1alpha1.AdapterOverrides {
	return t.Spec.AdapterOverrides
}

// SetDefaults implements apis.Defaultable
func (t *ElasticsearchTarget) SetDefaults(ctx context.Context) {
}

// Validate implements apis.Validatable
func (t *ElasticsearchTarget) Validate(ctx context.Context) *apis.FieldError {
	return nil
}
