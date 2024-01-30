package sinks

import "k8s.io/apimachinery/pkg/runtime/schema"

// GroupName ...
const GroupName = "sinks.typhoon.zeiss.com"

// NatsSinkResource ...
var NatsSinkResource = schema.GroupResource{
	Group:    GroupName,
	Resource: "nats-sinks",
}
