package resource

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"knative.dev/networking/pkg/apis/networking"
	"knative.dev/serving/pkg/apis/serving"
	servingv1 "knative.dev/serving/pkg/apis/serving/v1"
)

// NewKnService creates a Knative Service object.
func NewKnService(ns, name string, opts ...ObjectOption) *servingv1.Service {
	ks := &servingv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: ns,
			Name:      name,
		},
	}

	for _, opt := range opts {
		opt(ks)
	}

	// ensure the container name is not empty
	containers := ks.Spec.Template.Spec.Containers
	if len(containers) == 1 && containers[0].Name == "" {
		containers[0].Name = defaultContainerName
	}

	return ks
}

// VisibilityClusterLocal makes the Knative Service only available on the
// cluster's local network.
func VisibilityClusterLocal(object interface{}) {
	ks := object.(*servingv1.Service)
	Label(networking.VisibilityLabelKey, serving.VisibilityClusterLocal)(ks)
}

// VisibilityPublic makes the Knative Service available on the public internet.
func VisibilityPublic(object interface{}) {
	ks := object.(*servingv1.Service)
	delete(ks.Labels, networking.VisibilityLabelKey)
}
