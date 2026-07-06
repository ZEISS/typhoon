package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KafkaTarget is the Schema for an KafkaTarget.
type KafkaTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KafkaTargetSpec `json:"spec"`
	Status v1alpha1.Status `json:"status,omitempty"`
}

// Check the interfaces the event target should be implementing.
var (
	_ v1alpha1.Reconcilable        = (*KafkaTarget)(nil)
	_ v1alpha1.AdapterConfigurable = (*KafkaTarget)(nil)
)

// KafkaTargetSpec defines the desired state of the event target.
type KafkaTargetSpec struct {
	TopicReplicationFactor *int16                     `json:"topicReplicationFactor,omitempty"`
	TopicPartitions        *int32                     `json:"topicPartitions,omitempty"`
	Auth                   *KafkaTargetAuth           `json:"auth"`
	AdapterOverrides       *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
	Topic                  string                     `json:"topic"`
	BootstrapServers       []string                   `json:"bootstrapServers"`
	DiscardCEContext       bool                       `json:"discardCloudEventContext"`
}

// KafkaTargetAuth contains Authentication method used to interact with Kafka.
type KafkaTargetAuth struct {
	Kerberos           *KafkaTargetKerberos     `json:"kerberos,omitempty"`
	TLS                *KafkaTargetTLSAuth      `json:"tls,omitempty"`
	TLSEnable          *bool                    `json:"tlsEnable,omitempty"`
	SecurityMechanisms *string                  `json:"securityMechanism,omitempty"`
	Username           *string                  `json:"username,omitempty"`
	Password           *v1alpha1.ValueFromField `json:"password,omitempty"`
	SASLEnable         bool                     `json:"saslEnable"`
}

// KafkaTargetTLSAuth contains kerberos credentials.
type KafkaTargetTLSAuth struct {
	CA         *v1alpha1.ValueFromField `json:"ca,omitempty"`
	ClientCert *v1alpha1.ValueFromField `json:"clientCert,omitempty"`
	ClientKey  *v1alpha1.ValueFromField `json:"clientKey,omitempty"`
	SkipVerify *bool                    `json:"skipVerify,omitempty"`
}

// KafkaTargetKerberos contains kerberos credentials.
type KafkaTargetKerberos struct {
	Username    *string                  `json:"username,omitempty"`
	Password    *v1alpha1.ValueFromField `json:"password,omitempty"`
	ServiceName *string                  `json:"serviceName,omitempty"`
	ConfigPath  *string                  `json:"configPath,omitempty"`
	KeytabPath  *string                  `json:"keytabPath,omitempty"`
	Config      *v1alpha1.ValueFromField `json:"config,omitempty"`
	Keytab      *v1alpha1.ValueFromField `json:"keytab,omitempty"`
	Realm       *string                  `json:"realm,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KafkaTargetList is a list of event target instances.
type KafkaTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []KafkaTarget `json:"items"`
}
