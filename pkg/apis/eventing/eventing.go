package eventing

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	GroupName = "eventing.typhoon.zeiss.com"
)

// BrokersResource represents a Redis broker resource.
var RedisBrokersResource = schema.GroupResource{
	Group:    GroupName,
	Resource: "redisbrokers",
}

// Given an object accessor, return a list of owner references that are brokers.
func GetOwnerBrokers(object metav1.ObjectMetaAccessor) []metav1.OwnerReference {
	ors := []metav1.OwnerReference{}

	for _, or := range object.GetObjectMeta().GetOwnerReferences() {
		gv, err := schema.ParseGroupVersion(or.APIVersion)
		if err != nil {
			continue
		}

		if gv.Group == GroupName && IsBrokerKind(or.Kind) {
			ors = append(ors, or)
		}
	}

	return ors
}

func IsBrokerKind(kind string) bool {
	if kind == "RedisBroker" || kind == "MemoryBroker" {
		return true
	}

	return false
}
