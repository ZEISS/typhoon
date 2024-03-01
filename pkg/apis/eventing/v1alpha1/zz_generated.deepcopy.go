//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	broker "github.com/zeiss/typhoon/pkg/brokers/config/broker"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "knative.dev/eventing/pkg/apis/duck/v1"
	apis "knative.dev/pkg/apis"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Broker) DeepCopyInto(out *Broker) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
	if in.Observability != nil {
		in, out := &in.Observability, &out.Observability
		*out = new(Observability)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Broker.
func (in *Broker) DeepCopy() *Broker {
	if in == nil {
		return nil
	}
	out := new(Broker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Observability) DeepCopyInto(out *Observability) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Observability.
func (in *Observability) DeepCopy() *Observability {
	if in == nil {
		return nil
	}
	out := new(Observability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Redis) DeepCopyInto(out *Redis) {
	*out = *in
	if in.Connection != nil {
		in, out := &in.Connection, &out.Connection
		*out = new(RedisConnection)
		(*in).DeepCopyInto(*out)
	}
	if in.Stream != nil {
		in, out := &in.Stream, &out.Stream
		*out = new(string)
		**out = **in
	}
	if in.StreamMaxLen != nil {
		in, out := &in.StreamMaxLen, &out.StreamMaxLen
		*out = new(int)
		**out = **in
	}
	if in.EnableTrackingID != nil {
		in, out := &in.EnableTrackingID, &out.EnableTrackingID
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Redis.
func (in *Redis) DeepCopy() *Redis {
	if in == nil {
		return nil
	}
	out := new(Redis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisBroker) DeepCopyInto(out *RedisBroker) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisBroker.
func (in *RedisBroker) DeepCopy() *RedisBroker {
	if in == nil {
		return nil
	}
	out := new(RedisBroker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisBroker) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisBrokerList) DeepCopyInto(out *RedisBrokerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisBroker, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisBrokerList.
func (in *RedisBrokerList) DeepCopy() *RedisBrokerList {
	if in == nil {
		return nil
	}
	out := new(RedisBrokerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisBrokerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisBrokerSpec) DeepCopyInto(out *RedisBrokerSpec) {
	*out = *in
	if in.Redis != nil {
		in, out := &in.Redis, &out.Redis
		*out = new(Redis)
		(*in).DeepCopyInto(*out)
	}
	in.Broker.DeepCopyInto(&out.Broker)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisBrokerSpec.
func (in *RedisBrokerSpec) DeepCopy() *RedisBrokerSpec {
	if in == nil {
		return nil
	}
	out := new(RedisBrokerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisBrokerStatus) DeepCopyInto(out *RedisBrokerStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	in.Address.DeepCopyInto(&out.Address)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisBrokerStatus.
func (in *RedisBrokerStatus) DeepCopy() *RedisBrokerStatus {
	if in == nil {
		return nil
	}
	out := new(RedisBrokerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisConnection) DeepCopyInto(out *RedisConnection) {
	*out = *in
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	if in.ClusterURLs != nil {
		in, out := &in.ClusterURLs, &out.ClusterURLs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(SecretValueFromSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(SecretValueFromSource)
		(*in).DeepCopyInto(*out)
	}
	if in.TLSCACertificate != nil {
		in, out := &in.TLSCACertificate, &out.TLSCACertificate
		*out = new(SecretValueFromSource)
		(*in).DeepCopyInto(*out)
	}
	if in.TLSCertificate != nil {
		in, out := &in.TLSCertificate, &out.TLSCertificate
		*out = new(SecretValueFromSource)
		(*in).DeepCopyInto(*out)
	}
	if in.TLSKey != nil {
		in, out := &in.TLSKey, &out.TLSKey
		*out = new(SecretValueFromSource)
		(*in).DeepCopyInto(*out)
	}
	if in.TLSEnabled != nil {
		in, out := &in.TLSEnabled, &out.TLSEnabled
		*out = new(bool)
		**out = **in
	}
	if in.TLSSkipVerify != nil {
		in, out := &in.TLSSkipVerify, &out.TLSSkipVerify
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisConnection.
func (in *RedisConnection) DeepCopy() *RedisConnection {
	if in == nil {
		return nil
	}
	out := new(RedisConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretValueFromSource) DeepCopyInto(out *SecretValueFromSource) {
	*out = *in
	in.SecretKeyRef.DeepCopyInto(&out.SecretKeyRef)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretValueFromSource.
func (in *SecretValueFromSource) DeepCopy() *SecretValueFromSource {
	if in == nil {
		return nil
	}
	out := new(SecretValueFromSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trigger) DeepCopyInto(out *Trigger) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trigger.
func (in *Trigger) DeepCopy() *Trigger {
	if in == nil {
		return nil
	}
	out := new(Trigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Trigger) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerBounds) DeepCopyInto(out *TriggerBounds) {
	*out = *in
	if in.ById != nil {
		in, out := &in.ById, &out.ById
		*out = new(TriggerBoundsByID)
		(*in).DeepCopyInto(*out)
	}
	if in.ByDate != nil {
		in, out := &in.ByDate, &out.ByDate
		*out = new(TriggerBoundsByDate)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerBounds.
func (in *TriggerBounds) DeepCopy() *TriggerBounds {
	if in == nil {
		return nil
	}
	out := new(TriggerBounds)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerBoundsByDate) DeepCopyInto(out *TriggerBoundsByDate) {
	*out = *in
	if in.Start != nil {
		in, out := &in.Start, &out.Start
		*out = new(string)
		**out = **in
	}
	if in.End != nil {
		in, out := &in.End, &out.End
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerBoundsByDate.
func (in *TriggerBoundsByDate) DeepCopy() *TriggerBoundsByDate {
	if in == nil {
		return nil
	}
	out := new(TriggerBoundsByDate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerBoundsByID) DeepCopyInto(out *TriggerBoundsByID) {
	*out = *in
	if in.Start != nil {
		in, out := &in.Start, &out.Start
		*out = new(string)
		**out = **in
	}
	if in.End != nil {
		in, out := &in.End, &out.End
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerBoundsByID.
func (in *TriggerBoundsByID) DeepCopy() *TriggerBoundsByID {
	if in == nil {
		return nil
	}
	out := new(TriggerBoundsByID)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerList) DeepCopyInto(out *TriggerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Trigger, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerList.
func (in *TriggerList) DeepCopy() *TriggerList {
	if in == nil {
		return nil
	}
	out := new(TriggerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TriggerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpec) DeepCopyInto(out *TriggerSpec) {
	*out = *in
	in.Broker.DeepCopyInto(&out.Broker)
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make([]broker.Filter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Target.DeepCopyInto(&out.Target)
	if in.Delivery != nil {
		in, out := &in.Delivery, &out.Delivery
		*out = new(v1.DeliverySpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpec.
func (in *TriggerSpec) DeepCopy() *TriggerSpec {
	if in == nil {
		return nil
	}
	out := new(TriggerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpecBounded) DeepCopyInto(out *TriggerSpecBounded) {
	*out = *in
	in.TriggerSpec.DeepCopyInto(&out.TriggerSpec)
	if in.Bounds != nil {
		in, out := &in.Bounds, &out.Bounds
		*out = new(TriggerBounds)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpecBounded.
func (in *TriggerSpecBounded) DeepCopy() *TriggerSpecBounded {
	if in == nil {
		return nil
	}
	out := new(TriggerSpecBounded)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerStatus) DeepCopyInto(out *TriggerStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.TargetURI != nil {
		in, out := &in.TargetURI, &out.TargetURI
		*out = new(apis.URL)
		(*in).DeepCopyInto(*out)
	}
	in.DeliveryStatus.DeepCopyInto(&out.DeliveryStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerStatus.
func (in *TriggerStatus) DeepCopy() *TriggerStatus {
	if in == nil {
		return nil
	}
	out := new(TriggerStatus)
	in.DeepCopyInto(out)
	return out
}