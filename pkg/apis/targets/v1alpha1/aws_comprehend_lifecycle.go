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

// Returned event types
const (
	// EventTypeAWSComprehendResult contains the result of the processing of an S3 event.
	EventTypeAWSComprehendResult = "com.zeiss.targets.aws.comprehend.result"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*AWSComprehendTarget) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("AWSComprehendTarget")
}

// github.com/zeiss/typhoon.KRShaped.
func (*AWSComprehendTarget) GetConditionSet() apis.ConditionSet {
	return v1alpha1.DefaultConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (t *AWSComprehendTarget) GetStatus() *duckv1.Status {
	return &t.Status.Status
}

// GetStatusManager implements Reconcilable.
func (t *AWSComprehendTarget) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: t.GetConditionSet(),
		Status:       &t.Status,
	}
}

// GetEventTypes implements EventSource.
func (*AWSComprehendTarget) GetEventTypes() []string {
	return []string{
		EventTypeAWSComprehendResult,
	}
}

// AsEventSource implements EventSource.
func (t *AWSComprehendTarget) AsEventSource() string {
	kind := strings.ToLower(t.GetGroupVersionKind().Kind)
	return "com.zeiss." + kind + "." + t.Namespace + "." + t.Name
}

// GetAdapterOverrides implements AdapterConfigurable.
func (t *AWSComprehendTarget) GetAdapterOverrides() *v1alpha1.AdapterOverrides {
	return t.Spec.AdapterOverrides
}

// WantsOwnServiceAccount implements ServiceAccountProvider.
func (t *AWSComprehendTarget) WantsOwnServiceAccount() bool {
	return t.Spec.Auth.WantsOwnServiceAccount()
}

// ServiceAccountOptions implements ServiceAccountProvider.
func (t *AWSComprehendTarget) ServiceAccountOptions() []resource.ServiceAccountOption {
	return t.Spec.Auth.ServiceAccountOptions()
}

// SetDefaults implements apis.Defaultable
func (t *AWSComprehendTarget) SetDefaults(ctx context.Context) {
}

// Validate implements apis.Validatable
func (t *AWSComprehendTarget) Validate(ctx context.Context) *apis.FieldError {
	// Do not validate authentication object in case of resource deletion
	if t.DeletionTimestamp != nil {
		return nil
	}
	return t.Spec.Auth.Validate(ctx)
}
