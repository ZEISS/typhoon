//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	commonv1alpha1 "github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudEventsSource) DeepCopyInto(out *CloudEventsSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudEventsSource.
func (in *CloudEventsSource) DeepCopy() *CloudEventsSource {
	if in == nil {
		return nil
	}
	out := new(CloudEventsSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudEventsSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudEventsSourceList) DeepCopyInto(out *CloudEventsSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudEventsSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudEventsSourceList.
func (in *CloudEventsSourceList) DeepCopy() *CloudEventsSourceList {
	if in == nil {
		return nil
	}
	out := new(CloudEventsSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudEventsSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudEventsSourceSpec) DeepCopyInto(out *CloudEventsSourceSpec) {
	*out = *in
	in.SourceSpec.DeepCopyInto(&out.SourceSpec)
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(HTTPCredentials)
		(*in).DeepCopyInto(*out)
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.RateLimiter != nil {
		in, out := &in.RateLimiter, &out.RateLimiter
		*out = new(RateLimiter)
		**out = **in
	}
	if in.AdapterOverrides != nil {
		in, out := &in.AdapterOverrides, &out.AdapterOverrides
		*out = new(commonv1alpha1.AdapterOverrides)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudEventsSourceSpec.
func (in *CloudEventsSourceSpec) DeepCopy() *CloudEventsSourceSpec {
	if in == nil {
		return nil
	}
	out := new(CloudEventsSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPBasicAuth) DeepCopyInto(out *HTTPBasicAuth) {
	*out = *in
	in.Password.DeepCopyInto(&out.Password)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPBasicAuth.
func (in *HTTPBasicAuth) DeepCopy() *HTTPBasicAuth {
	if in == nil {
		return nil
	}
	out := new(HTTPBasicAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPCredentials) DeepCopyInto(out *HTTPCredentials) {
	*out = *in
	if in.BasicAuths != nil {
		in, out := &in.BasicAuths, &out.BasicAuths
		*out = make([]HTTPBasicAuth, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPCredentials.
func (in *HTTPCredentials) DeepCopy() *HTTPCredentials {
	if in == nil {
		return nil
	}
	out := new(HTTPCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPPollerSource) DeepCopyInto(out *HTTPPollerSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPPollerSource.
func (in *HTTPPollerSource) DeepCopy() *HTTPPollerSource {
	if in == nil {
		return nil
	}
	out := new(HTTPPollerSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HTTPPollerSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPPollerSourceList) DeepCopyInto(out *HTTPPollerSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HTTPPollerSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPPollerSourceList.
func (in *HTTPPollerSourceList) DeepCopy() *HTTPPollerSourceList {
	if in == nil {
		return nil
	}
	out := new(HTTPPollerSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HTTPPollerSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPPollerSourceSpec) DeepCopyInto(out *HTTPPollerSourceSpec) {
	*out = *in
	in.SourceSpec.DeepCopyInto(&out.SourceSpec)
	if in.EventSource != nil {
		in, out := &in.EventSource, &out.EventSource
		*out = new(string)
		**out = **in
	}
	in.Endpoint.DeepCopyInto(&out.Endpoint)
	if in.SkipVerify != nil {
		in, out := &in.SkipVerify, &out.SkipVerify
		*out = new(bool)
		**out = **in
	}
	if in.CACertificate != nil {
		in, out := &in.CACertificate, &out.CACertificate
		*out = new(string)
		**out = **in
	}
	if in.BasicAuthUsername != nil {
		in, out := &in.BasicAuthUsername, &out.BasicAuthUsername
		*out = new(string)
		**out = **in
	}
	if in.BasicAuthPassword != nil {
		in, out := &in.BasicAuthPassword, &out.BasicAuthPassword
		*out = new(commonv1alpha1.ValueFromField)
		(*in).DeepCopyInto(*out)
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AdapterOverrides != nil {
		in, out := &in.AdapterOverrides, &out.AdapterOverrides
		*out = new(commonv1alpha1.AdapterOverrides)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPPollerSourceSpec.
func (in *HTTPPollerSourceSpec) DeepCopy() *HTTPPollerSourceSpec {
	if in == nil {
		return nil
	}
	out := new(HTTPPollerSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaSource) DeepCopyInto(out *KafkaSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaSource.
func (in *KafkaSource) DeepCopy() *KafkaSource {
	if in == nil {
		return nil
	}
	out := new(KafkaSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaSourceAuth) DeepCopyInto(out *KafkaSourceAuth) {
	*out = *in
	if in.Kerberos != nil {
		in, out := &in.Kerberos, &out.Kerberos
		*out = new(KafkaSourceKerberos)
		(*in).DeepCopyInto(*out)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(KafkaSourceTLSAuth)
		(*in).DeepCopyInto(*out)
	}
	if in.TLSEnable != nil {
		in, out := &in.TLSEnable, &out.TLSEnable
		*out = new(bool)
		**out = **in
	}
	if in.SecurityMechanisms != nil {
		in, out := &in.SecurityMechanisms, &out.SecurityMechanisms
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(commonv1alpha1.ValueFromField)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaSourceAuth.
func (in *KafkaSourceAuth) DeepCopy() *KafkaSourceAuth {
	if in == nil {
		return nil
	}
	out := new(KafkaSourceAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaSourceKerberos) DeepCopyInto(out *KafkaSourceKerberos) {
	*out = *in
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(commonv1alpha1.ValueFromField)
		(*in).DeepCopyInto(*out)
	}
	if in.Realm != nil {
		in, out := &in.Realm, &out.Realm
		*out = new(string)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
	if in.ConfigPath != nil {
		in, out := &in.ConfigPath, &out.ConfigPath
		*out = new(string)
		**out = **in
	}
	if in.KeytabPath != nil {
		in, out := &in.KeytabPath, &out.KeytabPath
		*out = new(string)
		**out = **in
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(commonv1alpha1.ValueFromField)
		(*in).DeepCopyInto(*out)
	}
	if in.Keytab != nil {
		in, out := &in.Keytab, &out.Keytab
		*out = new(commonv1alpha1.ValueFromField)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaSourceKerberos.
func (in *KafkaSourceKerberos) DeepCopy() *KafkaSourceKerberos {
	if in == nil {
		return nil
	}
	out := new(KafkaSourceKerberos)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaSourceList) DeepCopyInto(out *KafkaSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KafkaSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaSourceList.
func (in *KafkaSourceList) DeepCopy() *KafkaSourceList {
	if in == nil {
		return nil
	}
	out := new(KafkaSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaSourceSpec) DeepCopyInto(out *KafkaSourceSpec) {
	*out = *in
	in.SourceSpec.DeepCopyInto(&out.SourceSpec)
	if in.BootstrapServers != nil {
		in, out := &in.BootstrapServers, &out.BootstrapServers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Auth.DeepCopyInto(&out.Auth)
	if in.AdapterOverrides != nil {
		in, out := &in.AdapterOverrides, &out.AdapterOverrides
		*out = new(commonv1alpha1.AdapterOverrides)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaSourceSpec.
func (in *KafkaSourceSpec) DeepCopy() *KafkaSourceSpec {
	if in == nil {
		return nil
	}
	out := new(KafkaSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaSourceTLSAuth) DeepCopyInto(out *KafkaSourceTLSAuth) {
	*out = *in
	if in.CA != nil {
		in, out := &in.CA, &out.CA
		*out = new(commonv1alpha1.ValueFromField)
		(*in).DeepCopyInto(*out)
	}
	if in.ClientCert != nil {
		in, out := &in.ClientCert, &out.ClientCert
		*out = new(commonv1alpha1.ValueFromField)
		(*in).DeepCopyInto(*out)
	}
	if in.ClientKey != nil {
		in, out := &in.ClientKey, &out.ClientKey
		*out = new(commonv1alpha1.ValueFromField)
		(*in).DeepCopyInto(*out)
	}
	if in.SkipVerify != nil {
		in, out := &in.SkipVerify, &out.SkipVerify
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaSourceTLSAuth.
func (in *KafkaSourceTLSAuth) DeepCopy() *KafkaSourceTLSAuth {
	if in == nil {
		return nil
	}
	out := new(KafkaSourceTLSAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIMetrics) DeepCopyInto(out *OCIMetrics) {
	*out = *in
	if in.Compartment != nil {
		in, out := &in.Compartment, &out.Compartment
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIMetrics.
func (in *OCIMetrics) DeepCopy() *OCIMetrics {
	if in == nil {
		return nil
	}
	out := new(OCIMetrics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in OCIMetricsDecodedList) DeepCopyInto(out *OCIMetricsDecodedList) {
	{
		in := &in
		*out = make(OCIMetricsDecodedList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIMetricsDecodedList.
func (in OCIMetricsDecodedList) DeepCopy() OCIMetricsDecodedList {
	if in == nil {
		return nil
	}
	out := new(OCIMetricsDecodedList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIMetricsSource) DeepCopyInto(out *OCIMetricsSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIMetricsSource.
func (in *OCIMetricsSource) DeepCopy() *OCIMetricsSource {
	if in == nil {
		return nil
	}
	out := new(OCIMetricsSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OCIMetricsSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIMetricsSourceList) DeepCopyInto(out *OCIMetricsSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OCIMetricsSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIMetricsSourceList.
func (in *OCIMetricsSourceList) DeepCopy() *OCIMetricsSourceList {
	if in == nil {
		return nil
	}
	out := new(OCIMetricsSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OCIMetricsSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIMetricsSourceSpec) DeepCopyInto(out *OCIMetricsSourceSpec) {
	*out = *in
	in.SourceSpec.DeepCopyInto(&out.SourceSpec)
	in.OracleAPIPrivateKey.DeepCopyInto(&out.OracleAPIPrivateKey)
	in.OracleAPIPrivateKeyPassphrase.DeepCopyInto(&out.OracleAPIPrivateKeyPassphrase)
	in.OracleAPIPrivateKeyFingerprint.DeepCopyInto(&out.OracleAPIPrivateKeyFingerprint)
	if in.PollingFrequency != nil {
		in, out := &in.PollingFrequency, &out.PollingFrequency
		*out = new(string)
		**out = **in
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]OCIMetrics, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AdapterOverrides != nil {
		in, out := &in.AdapterOverrides, &out.AdapterOverrides
		*out = new(commonv1alpha1.AdapterOverrides)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIMetricsSourceSpec.
func (in *OCIMetricsSourceSpec) DeepCopy() *OCIMetricsSourceSpec {
	if in == nil {
		return nil
	}
	out := new(OCIMetricsSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimiter) DeepCopyInto(out *RateLimiter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimiter.
func (in *RateLimiter) DeepCopy() *RateLimiter {
	if in == nil {
		return nil
	}
	out := new(RateLimiter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookEventExtensionAttributes) DeepCopyInto(out *WebhookEventExtensionAttributes) {
	*out = *in
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookEventExtensionAttributes.
func (in *WebhookEventExtensionAttributes) DeepCopy() *WebhookEventExtensionAttributes {
	if in == nil {
		return nil
	}
	out := new(WebhookEventExtensionAttributes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookSource) DeepCopyInto(out *WebhookSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookSource.
func (in *WebhookSource) DeepCopy() *WebhookSource {
	if in == nil {
		return nil
	}
	out := new(WebhookSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebhookSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookSourceList) DeepCopyInto(out *WebhookSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WebhookSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookSourceList.
func (in *WebhookSourceList) DeepCopy() *WebhookSourceList {
	if in == nil {
		return nil
	}
	out := new(WebhookSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebhookSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookSourceSpec) DeepCopyInto(out *WebhookSourceSpec) {
	*out = *in
	in.SourceSpec.DeepCopyInto(&out.SourceSpec)
	if in.EventSource != nil {
		in, out := &in.EventSource, &out.EventSource
		*out = new(string)
		**out = **in
	}
	if in.EventExtensionAttributes != nil {
		in, out := &in.EventExtensionAttributes, &out.EventExtensionAttributes
		*out = new(WebhookEventExtensionAttributes)
		(*in).DeepCopyInto(*out)
	}
	if in.BasicAuthUsername != nil {
		in, out := &in.BasicAuthUsername, &out.BasicAuthUsername
		*out = new(string)
		**out = **in
	}
	if in.BasicAuthPassword != nil {
		in, out := &in.BasicAuthPassword, &out.BasicAuthPassword
		*out = new(commonv1alpha1.ValueFromField)
		(*in).DeepCopyInto(*out)
	}
	if in.CORSAllowOrigin != nil {
		in, out := &in.CORSAllowOrigin, &out.CORSAllowOrigin
		*out = new(string)
		**out = **in
	}
	if in.AdapterOverrides != nil {
		in, out := &in.AdapterOverrides, &out.AdapterOverrides
		*out = new(commonv1alpha1.AdapterOverrides)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookSourceSpec.
func (in *WebhookSourceSpec) DeepCopy() *WebhookSourceSpec {
	if in == nil {
		return nil
	}
	out := new(WebhookSourceSpec)
	in.DeepCopyInto(out)
	return out
}
