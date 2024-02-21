package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LogzMetricsTarget ...
type LogzMetricsTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LogzMetricsTargetSpec `json:"spec"`
	Status v1alpha1.Status       `json:"status,omitempty"`
}

// Check the interfaces the event target should be implementing.
var (
	_ v1alpha1.Reconcilable        = (*LogzMetricsTarget)(nil)
	_ v1alpha1.AdapterConfigurable = (*LogzMetricsTarget)(nil)
	_ v1alpha1.EventReceiver       = (*LogzMetricsTarget)(nil)
)

// LogzMetricsTargetSpec defines the desired state of the event target.
type LogzMetricsTargetSpec struct {
	// Connection information for LogzMetrics.
	Connection LogzMetricsConnection `json:"connection"`

	// Instruments configured for pushing metrics. It is mandatory that all metrics
	// pushed by using this target are pre-registered using this list.
	Instruments []Instrument `json:"instruments"`

	// EventOptions for targets
	// +optional
	EventOptions *EventOptions `json:"eventOptions,omitempty"`

	// Adapter spec overrides parameters.
	// +optional
	AdapterOverrides *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
}

// LogzMetricsConnection contains the information to connect to a Logz tenant to push metrics.
type LogzMetricsConnection struct {
	// Token for connecting to Logz metrics listener.
	Token SecretValueFromSource `json:"token"`

	// ListenerURL for pushing metrics.
	ListenerURL string `json:"listenerURL"`
}

// InstrumentKind as defined by OpenTelemetry.
type InstrumentKind string

// NumberKind as defined by OpenTelemetry.
type NumberKind string

const (
	// Instrument Kinds
	InstrumentKindHistogram     InstrumentKind = "Histogram"
	InstrumentKindCounter       InstrumentKind = "Counter"
	InstrumentKindUpDownCounter InstrumentKind = "UpDownCounter"

	// Number Kinds
	NumberKindInt64   NumberKind = "Int64"
	NumberKindFloat64 NumberKind = "Float64"
)

// Instrument push metrics for.
type Instrument struct {
	// Name for the Instrument.
	Name string `json:"name"`

	// Description for the Instrument
	// +optional
	Description *string `json:"description,omitempty"`

	// Instrument Kind as defined by OpenTelemetry. Supported values are:
	//
	// - Histogram: for absolute values that can be aggregated.
	// - Counter: for delta values that increase monotonically.
	// - UpDownCounter: for delta values that can increase and decrease.
	Instrument InstrumentKind `json:"instrument"`

	// Number Kind as defined by OpenTelemetry, defines the measure data type
	// accepted by the Instrument. Supported values are:
	//
	// - Int64.
	// - Float64.
	Number NumberKind `json:"number"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LogzMetricsTargetList is a list of event target instances.
type LogzMetricsTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []LogzMetricsTarget `json:"items"`
}
