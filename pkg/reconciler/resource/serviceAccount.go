package resource

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ServiceAccountOption is a functional option for a ServiceAccount.
type ServiceAccountOption func(*corev1.ServiceAccount)

// NewServiceAccount creates a ServiceAccount object.
func NewServiceAccount(ns, name string, opts ...ObjectOption) *corev1.ServiceAccount {
	sa := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: ns,
			Name:      name,
		},
	}

	for _, opt := range opts {
		opt(sa)
	}

	return sa
}

func ServiceAccountWithMetaOptions(opts ...MetaOption) ObjectOption {
	return func(obj interface{}) {
		s := obj.(*corev1.ServiceAccount)

		for _, opt := range opts {
			opt(&s.ObjectMeta)
		}
	}
}
