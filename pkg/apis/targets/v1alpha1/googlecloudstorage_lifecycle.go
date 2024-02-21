

package v1alpha1

import (
	"context"
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	"github.com/zeiss/typhoon/pkg/reconciler/resource"
)

// Managed event types
const (
	EventTypeGoogleCloudStorageObjectInsert = "com.google.cloud.storage.object.insert"

	EventTypeGoogleCloudStorageResponse = "com.google.cloud.storage.object.insert.response"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*GoogleCloudStorageTarget) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("GoogleCloudStorageTarget")
}
github.com/zeiss/typhoon
// GetConditionSet implements duckv1.KRShaped.
func (*GoogleCloudStorageTarget) GetConditionSet() apis.ConditionSet {
	return v1alpha1.DefaultConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (t *GoogleCloudStorageTarget) GetStatus() *duckv1.Status {
	return &t.Status.Status
}

// GetStatusManager implements Reconcilable.
func (t *GoogleCloudStorageTarget) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: t.GetConditionSet(),
		Status:       &t.Status,
	}
}

// AcceptedEventTypes implements IntegrationTarget.
func (*GoogleCloudStorageTarget) AcceptedEventTypes() []string {
	return []string{
		EventTypeGoogleCloudStorageObjectInsert,
		EventTypeWildcard,
	}
}

// GetEventTypes implements EventSource.
func (*GoogleCloudStorageTarget) GetEventTypes() []string {
	return []string{
		EventTypeGoogleCloudStorageResponse,
	}
}

// AsEventSource implements EventSource.
func (t *GoogleCloudStorageTarget) AsEventSource() string {
	kind := strings.ToLower(t.GetGroupVersionKind().Kind)
	return "com.zeiss." + kind + "." + t.Namespace + "." + t.Name
}

// GetAdapterOverrides implements AdapterConfigurable.
func (t *GoogleCloudStorageTarget) GetAdapterOverrides() *v1alpha1.AdapterOverrides {
	return t.Spec.AdapterOverrides
}

// WantsOwnServiceAccount implements ServiceAccountProvider.
func (t *GoogleCloudStorageTarget) WantsOwnServiceAccount() bool {
	return t.Spec.Auth != nil && t.Spec.Auth.WantsOwnServiceAccount()
}

// ServiceAccountOptions implements ServiceAccountProvider.
func (t *GoogleCloudStorageTarget) ServiceAccountOptions() []resource.ServiceAccountOption {
	if t.Spec.Auth == nil {
		return []resource.ServiceAccountOption{}
	}
	return t.Spec.Auth.ServiceAccountOptions()
}

// SetDefaults implements apis.Defaultable
func (t *GoogleCloudStorageTarget) SetDefaults(ctx context.Context) {
}

// Validate implements apis.Validatable
func (t *GoogleCloudStorageTarget) Validate(ctx context.Context) *apis.FieldError {
	return nil
}
