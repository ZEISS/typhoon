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
	duckv1.SourceSpec              `json:",inline"`
	OracleAPIPrivateKey            v1alpha1.ValueFromField    `json:"oracleApiPrivateKey"`
	OracleAPIPrivateKeyPassphrase  v1alpha1.ValueFromField    `json:"oracleApiPrivateKeyPassphrase"`
	OracleAPIPrivateKeyFingerprint v1alpha1.ValueFromField    `json:"oracleApiPrivateKeyFingerprint"`
	PollingFrequency               *string                    `json:"metricsPollingFrequency,omitempty"`
	AdapterOverrides               *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
	Tenancy                        string                     `json:"oracleTenancy"`
	User                           string                     `json:"oracleUser"`
	Region                         string                     `json:"oracleRegion"`
	Metrics                        []OCIMetrics               `json:"metrics"`
}

// OCIMetrics represents OCI metrics structure.
type OCIMetrics struct {
	Compartment      *string `json:"oracleCompartment,omitempty"`
	Name             string  `json:"name"`
	MetricsNamespace string  `json:"metricsNamespace"`
	MetricsQuery     string  `json:"metricsQuery"`
}

// OCIMetricsDecodedList is a list of OCI metrics.
type OCIMetricsDecodedList []OCIMetrics

// Decode deserializes a list of OCI metrics.
func (o OCIMetricsDecodedList) Decode(value string) error {
	err := json.Unmarshal([]byte(value), &o)
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
