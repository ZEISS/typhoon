package resource

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"knative.dev/pkg/kmeta"
	"knative.dev/pkg/ptr"
)

// ObjectOption is a functional option for building Kubernetes API objects.
type ObjectOption func(interface{})

// Controller sets the given object as the controller (main owner) of an API object.
func Controller(obj kmeta.OwnerRefable) ObjectOption {
	return func(object interface{}) {
		meta := object.(metav1.Object)

		meta.SetOwnerReferences(append(meta.GetOwnerReferences(),
			*kmeta.NewControllerRef(obj),
		))
	}
}

// Owners sets the given objects as regular owners of an API object.
func Owners(owners ...kmeta.OwnerRefable) ObjectOption {
	ownerRefs := make([]metav1.OwnerReference, len(owners))
	for i, owner := range owners {
		ownerRefs[i] = *kmeta.NewControllerRef(owner)
		ownerRefs[i].Controller = ptr.Bool(false)
	}

	return func(object interface{}) {
		meta := object.(metav1.Object)

		meta.SetOwnerReferences(append(meta.GetOwnerReferences(),
			ownerRefs...,
		))
	}
}

// Label sets the value of an API object's label.
func Label(key, val string) ObjectOption {
	return func(object interface{}) {
		meta := object.(metav1.Object)

		lbls := meta.GetLabels()

		if lbls == nil {
			lbls = make(labels.Set, 1)
			meta.SetLabels(lbls)
		}
		lbls[key] = val
	}
}

// Labels sets multiple labels on an API object.
func Labels(ls labels.Set) ObjectOption {
	return func(object interface{}) {
		meta := object.(metav1.Object)

		lbls := meta.GetLabels()

		if lbls == nil {
			lbls = make(labels.Set, len(ls))
			meta.SetLabels(lbls)
		}

		for k, v := range ls {
			lbls[k] = v
		}
	}
}

// Annotation sets the value of an API object's annotation.
func Annotation(key, val string) ObjectOption {
	return func(object interface{}) {
		meta := object.(metav1.Object)

		annots := meta.GetAnnotations()

		if annots == nil {
			annots = make(labels.Set, 1)
			meta.SetAnnotations(annots)
		}
		annots[key] = val
	}
}
