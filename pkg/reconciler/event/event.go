// Package event contains functions for generating Kubernetes API events.
package event

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"knative.dev/pkg/controller"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// Normal records a normal event for an API object.
func Normal(ctx context.Context, reason, msgFmt string, args ...interface{}) {
	recordEvent(ctx, corev1.EventTypeNormal, reason, msgFmt, args...)
}

// Warn records a warning event for an API object.
func Warn(ctx context.Context, reason, msgFmt string, args ...interface{}) {
	recordEvent(ctx, corev1.EventTypeWarning, reason, msgFmt, args...)
}

func recordEvent(ctx context.Context, typ, reason, msgFmt string, args ...interface{}) {
	recorder := controller.GetEventRecorder(ctx)
	if recorder != nil {
		recorder.Eventf(v1alpha1.ReconcilableFromContext(ctx), typ, reason, msgFmt, args...)
	}
}
