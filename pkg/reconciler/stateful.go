package reconciler

import (
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Stateful events related
const (
	LabelBridgeUsedByPrefix  = "flow.typhoon.zeiss.com/used-by."
	LabelValueBridgeDominant = "dominant"

	// Bridge identifier for stateful flows
	EnvBridgeID = "EVENTS_BRIDGE_IDENTIFIER"
)

// GetStatefulBridgeID returns the BridgeID based on an object metadata.
//
// All bridge components controlled by Triggerflow have labels informing their
// relation with the bridges they are part of. A component can only have one
// dominant bridge, which is the one it synchronizes with. This function uses
// that label to retrieve the bridge name and use it as unique ID.
func GetStatefulBridgeID(object metav1.Object) string {
	labels := object.GetLabels()

	for k, v := range labels {
		if strings.HasPrefix(k, LabelBridgeUsedByPrefix) && v == LabelValueBridgeDominant {
			return k[len(LabelBridgeUsedByPrefix):]
		}
	}

	return ""
}
