package resource

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type (
	PodTemplateSpecOption func(*corev1.PodTemplateSpec)
	DeploymentOption      func(*appsv1.Deployment)
	PodSpecOption         func(*corev1.PodSpec)
)

// NewDeployment creates a Deployment object.
func NewDeployment(ns, name string, opts ...ObjectOption) *appsv1.Deployment {
	d := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: ns,
			Name:      name,
		},
	}

	for _, opt := range opts {
		opt(d)
	}

	// If the Deployment was created without defining a Container
	// explicitly, ensure its default container's name is not empty.
	containers := d.Spec.Template.Spec.Containers
	if len(containers) == 1 && containers[0].Name == "" {
		containers[0].Name = defaultContainerName
	}

	return d
}

// Selector adds a label selector to a Deployment's spec, ensuring a
// corresponding label exists in the Pod template.
func Selector(key, val string) ObjectOption {
	return func(object interface{}) {
		d := object.(*appsv1.Deployment)

		selector := &d.Spec.Selector

		if *selector == nil {
			*selector = &metav1.LabelSelector{}
		}
		*selector = metav1.AddLabelToSelector(*selector, key, val)

		PodLabel(key, val)(d)
	}
}

func DeploymentSetReplicas(replicas int32) ObjectOption {
	return func(object interface{}) {
		d := object.(*appsv1.Deployment)

		d.Spec.Replicas = &replicas
	}
}

func DeploymentWithMetaOptions(opts ...MetaOption) ObjectOption {
	return func(object interface{}) {
		d := object.(*appsv1.Deployment)

		for _, opt := range opts {
			opt(&d.ObjectMeta)
		}
	}
}

func DeploymentAddSelectorForTemplate(key, value string) ObjectOption {
	return func(object interface{}) {
		d := object.(*appsv1.Deployment)

		if d.Spec.Selector == nil {
			d.Spec.Selector = &metav1.LabelSelector{}
		}

		sl := d.Spec.Selector.MatchLabels
		if sl == nil {
			sl = make(map[string]string, 1)
			d.Spec.Selector.MatchLabels = sl
		}
		sl[key] = value

		MetaAddLabel(key, value)(&d.Spec.Template.ObjectMeta)
	}
}

func DeploymentWithTemplateSpecOptions(opts ...ObjectOption) ObjectOption {
	return func(object interface{}) {
		d := object.(*appsv1.Deployment)

		for _, opt := range opts {
			opt(&d.Spec.Template)
		}
	}
}

func PodTemplateSpecWithMetaOptions(opts ...MetaOption) ObjectOption {
	return func(obj interface{}) {
		p := obj.(*corev1.PodTemplateSpec)

		for _, opt := range opts {
			opt(&p.ObjectMeta)
		}
	}
}

func PodSpecAddContainer(opt *corev1.Container) ObjectOption {
	return func(object interface{}) {
		p := object.(*corev1.PodSpec)

		if p.Containers == nil {
			p.Containers = make([]corev1.Container, 0, 1)
		}
		p.Containers = append(p.Containers, *opt)
	}
}

func PodTemplateSpecWithPodSpecOptions(opts ...ObjectOption) ObjectOption {
	return func(obj interface{}) {
		pts := obj.(*corev1.PodTemplateSpec)

		for _, opt := range opts {
			opt(&pts.Spec)
		}
	}
}

func PodSpecWithServiceAccountName(saName string) ObjectOption {
	return func(obj interface{}) {
		ps := obj.(*corev1.PodSpec)

		ps.ServiceAccountName = saName
	}
}
