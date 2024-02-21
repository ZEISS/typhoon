package resource

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewConfigMap creates a ConfigMap object.
func NewConfigMap(ns, name string, opts ...ObjectOption) *corev1.ConfigMap {
	cmap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: ns,
			Name:      name,
		},
	}

	for _, opt := range opts {
		opt(cmap)
	}

	return cmap
}

// Data sets one UTF-8 data entry in a ConfigMap.
func Data(key, value string) ObjectOption {
	return func(object interface{}) {
		cmap := object.(*corev1.ConfigMap)

		bdata := &cmap.Data

		if *bdata == nil {
			*bdata = make(map[string]string, 1)
		}

		(*bdata)[key] = value
	}
}
