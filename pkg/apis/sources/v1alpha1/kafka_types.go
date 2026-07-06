package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KafkaSource is the Schema for the KafkaSource.
type KafkaSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KafkaSourceSpec `json:"spec,omitempty"`
	Status v1alpha1.Status `json:"status,omitempty"`
}

// Check the interfaces the event source should be implementing.
var (
	_ v1alpha1.Reconcilable        = (*KafkaSource)(nil)
	_ v1alpha1.EventSender         = (*KafkaSource)(nil)
	_ v1alpha1.AdapterConfigurable = (*KafkaSource)(nil)
	_ v1alpha1.EventSource         = (*KafkaSource)(nil)
)

// KafkaSourceSpec defines the desired state of the event source.
type KafkaSourceSpec struct {
	Auth              KafkaSourceAuth `json:"auth"`
	duckv1.SourceSpec `json:",inline"`
	AdapterOverrides  *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
	Topic             string                     `json:"topic"`
	GroupID           string                     `json:"groupID"`
	BootstrapServers  []string                   `json:"bootstrapServers"`
}

// KafkaSourceAuth contains Authentication method used to interact with Kafka.
type KafkaSourceAuth struct {
	Kerberos           *KafkaSourceKerberos     `json:"kerberos,omitempty"`
	TLS                *KafkaSourceTLSAuth      `json:"tls,omitempty"`
	TLSEnable          *bool                    `json:"tlsEnable,omitempty"`
	SecurityMechanisms *string                  `json:"securityMechanism,omitempty"`
	Username           *string                  `json:"username,omitempty"`
	Password           *v1alpha1.ValueFromField `json:"password,omitempty"`
	SASLEnable         bool                     `json:"saslEnable"`
}

// KafkaSourceTLSAuth contains kerberos credentials.
type KafkaSourceTLSAuth struct {
	CA         *v1alpha1.ValueFromField `json:"ca,omitempty"`
	ClientCert *v1alpha1.ValueFromField `json:"clientCert,omitempty"`
	ClientKey  *v1alpha1.ValueFromField `json:"clientKey,omitempty"`
	SkipVerify *bool                    `json:"skipVerify,omitempty"`
}

// KafkaSourceKerberos contains kerberos credentials.
type KafkaSourceKerberos struct {
	Username    *string                  `json:"username,omitempty"`
	Password    *v1alpha1.ValueFromField `json:"password,omitempty"`
	Realm       *string                  `json:"realm,omitempty"`
	ServiceName *string                  `json:"serviceName,omitempty"`
	ConfigPath  *string                  `json:"configPath,omitempty"`
	KeytabPath  *string                  `json:"keytabPath,omitempty"`
	Config      *v1alpha1.ValueFromField `json:"config,omitempty"`
	Keytab      *v1alpha1.ValueFromField `json:"keytab,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KafkaSourceList contains a list of event sources.
type KafkaSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KafkaSource `json:"items"`
}
