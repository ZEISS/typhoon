package common

import (
	"strings"

	"knative.dev/pkg/kmeta"

	eventingv1alpha1 "github.com/zeiss/typhoon/pkg/apis/eventing/v1alpha1"
)

const (
	configMapResourceSuffix = "status"

	// Name of the status key inside the Status ConfigMap
	ConfigMapStatusKey = "status"
)

func AppAnnotationValue(or kmeta.OwnerRefable) string {
	return strings.ToLower(or.GetGroupVersionKind().Kind)
}

func GetBrokerConfigMapName(b eventingv1alpha1.ReconcilableBroker) string {
	if b == nil {
		return ""
	}

	return b.GetObjectMeta().GetName() + "-" + b.GetOwnedObjectsSuffix() + "-" + configMapResourceSuffix
}
