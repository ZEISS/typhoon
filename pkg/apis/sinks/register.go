package sinks

import "k8s.io/apimachinery/pkg/runtime/schema"

// GroupName is the group name use in this package
const GroupName = "sinks.typhoon.zeiss.com"

// NatsSinkResource is the resource for the NATS sink
var NatsSinkResource = schema.GroupResource{
	Group:    GroupName,
	Resource: "natssinks",
}
