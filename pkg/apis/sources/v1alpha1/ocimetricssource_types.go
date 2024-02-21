package v1alpha1

import (
	"encoding/json"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OCIMetricsSource is the schema for the event source.
type OCIMetricsSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OCIMetricsSourceSpec `json:"spec,omitempty"`
	Status v1alpha1.Status      `json:"status,omitempty"`
}

// Check the interfaces the event source should be implementing.
var (
	_ v1alpha1.Reconcilable        = (*OCIMetricsSource)(nil)
	_ v1alpha1.AdapterConfigurable = (*OCIMetricsSource)(nil)
	_ v1alpha1.EventSource         = (*OCIMetricsSource)(nil)
	_ v1alpha1.EventSender         = (*OCIMetricsSource)(nil)
)

// OCIMetricsSourceSpec defines the desired state of the event source.
type OCIMetricsSourceSpec struct {
	// inherits duck/v1 SourceSpec, which currently provides:
	// * Sink - a reference to an object that will resolve to a domain name or
	//   a URI directly to use as the sink.
	// * CloudEventOverrides - defines overrides to control the output format
	//   and modifications of the event sent to the sink.
	duckv1.SourceSpec `json:",inline"`

	// Oracle User API private key
	OracleAPIPrivateKey v1alpha1.ValueFromField `json:"oracleApiPrivateKey"`

	// Oracle User API private key passphrase
	OracleAPIPrivateKeyPassphrase v1alpha1.ValueFromField `json:"oracleApiPrivateKeyPassphrase"`

	// Oracle User API cert fingerprint
	OracleAPIPrivateKeyFingerprint v1alpha1.ValueFromField `json:"oracleApiPrivateKeyFingerprint"`

	// Oracle Tenancy OCID
	Tenancy string `json:"oracleTenancy"`

	// Oracle User OCID associated with the API key
	User string `json:"oracleUser"`

	// Oracle Cloud Region
	Region string `json:"oracleRegion"`

	// OCI Metrics Polling Frequency
	// +optional
	PollingFrequency *string `json:"metricsPollingFrequency,omitempty"`

	// Array of metrics
	Metrics []OCIMetrics `json:"metrics"`

	// Adapter spec overrides parameters.
	// +optional
	AdapterOverrides *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
}

// OCIMetrics represents OCI metrics structure.
type OCIMetrics struct {
	// Human description for the metrics entry
	Name string `json:"name"`

	// Namespace for the query metric to use
	MetricsNamespace string `json:"metricsNamespace"`

	// OCI Metrics Query See https://docs.cloud.oracle.com/en-us/iaas/api/#/en/monitoring/20180401/MetricData
	MetricsQuery string `json:"metricsQuery"`

	// Oracle Compartment OCID
	Compartment *string `json:"oracleCompartment,omitempty"`
}

// OCIMetricsDecodedList is a list of OCI metrics.
type OCIMetricsDecodedList []OCIMetrics

// Decode deserializes a list of OCI metrics.
func (o *OCIMetricsDecodedList) Decode(value string) error {
	err := json.Unmarshal([]byte(value), o)
	if err != nil {
		return fmt.Errorf("unable to deserialize metrics: %w", err)
	}

	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OCIMetricsSourceList contains a list of event sources.
type OCIMetricsSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OCIMetricsSource `json:"items"`
}
