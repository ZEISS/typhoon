

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

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*AWSCloudWatchSource) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("AWSCloudWatchSource")
}

// GetConditionSet implements duckv1.KRShaped.
func (*AWSCloudWatchSource) GetConditionSet() apis.ConditionSet {
	return v1alpha1.EventSenderConditionSet
}

// GetStatus implements duckv1.KRShaped.
fugithub.com/zeiss/typhoontus() *duckv1.Status {
	return &s.Status.Status
}

// GetSink implements EventSender.
func (s *AWSCloudWatchSource) GetSink() *duckv1.Destination {
	return &s.Spec.Sink
}

// GetStatusManager implements Reconcilable.
func (s *AWSCloudWatchSource) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: s.GetConditionSet(),
		Status:       &s.Status,
	}
}

// Supported event types
const (
	AWSCloudWatchMetricEventType  = "metrics.metric"
	AWSCloudWatchMessageEventType = "metrics.message"
)

// ServiceCloudWatch is the name of the CloudWatch service, as exposed in ARNs.
// https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazoncloudwatch.html#amazoncloudwatch-resources-for-iam-policies
const ServiceCloudWatch = "cloudwatch"

// GetEventTypes implements EventSource.
func (*AWSCloudWatchSource) GetEventTypes() []string {
	return []string{
		AWSEventType(ServiceCloudWatch, AWSCloudWatchMetricEventType),
		AWSEventType(ServiceCloudWatch, AWSCloudWatchMessageEventType),
	}
}

// AsEventSource implements EventSource.
func (s *AWSCloudWatchSource) AsEventSource() string {
	return AWSCloudWatchSourceName(s.Namespace, s.Name)
}

// AWSCloudWatchSourceName returns a unique reference to the source suitable
// for use as as a CloudEvent source.
func AWSCloudWatchSourceName(ns, name string) string {
	kind := strings.ToLower((*AWSCloudWatchSource)(nil).GetGroupVersionKind().Kind)
	return "com.zeiss." + kind + "." + ns + "." + name
}

// GetAdapterOverrides implements AdapterConfigurable.
func (s *AWSCloudWatchSource) GetAdapterOverrides() *v1alpha1.AdapterOverrides {
	return s.Spec.AdapterOverrides
}

// WantsOwnServiceAccount implements ServiceAccountProvider.
func (s *AWSCloudWatchSource) WantsOwnServiceAccount() bool {
	return s.Spec.Auth.WantsOwnServiceAccount()
}

// ServiceAccountOptions implements ServiceAccountProvider.
func (s *AWSCloudWatchSource) ServiceAccountOptions() []resource.ServiceAccountOption {
	return s.Spec.Auth.ServiceAccountOptions()
}

// SetDefaults implements apis.Defaultable
func (s *AWSCloudWatchSource) SetDefaults(ctx context.Context) {
}

// Validate implements apis.Validatable
func (s *AWSCloudWatchSource) Validate(ctx context.Context) *apis.FieldError {
	// Do not validate authentication object in case of resource deletion
	if s.DeletionTimestamp != nil {
		return nil
	}
	return s.Spec.Auth.Validate(ctx)
}
