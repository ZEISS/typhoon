package main

import (
	"context"
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/injection/sharedmain"
	"knative.dev/pkg/signals"
	"knative.dev/pkg/webhook"
	"knative.dev/pkg/webhook/certificates"
	"knative.dev/pkg/webhook/resourcesemantics"
	"knative.dev/pkg/webhook/resourcesemantics/defaulting"
	"knative.dev/pkg/webhook/resourcesemantics/validation"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	extensionsv1alpha1 "github.com/zeiss/typhoon/pkg/apis/extensions/v1alpha1"
	flowv1alpha1 "github.com/zeiss/typhoon/pkg/apis/flow/v1alpha1"
	routingv1alpha1 "github.com/zeiss/typhoon/pkg/apis/routing/v1alpha1"
	sourcesv1alpha1 "github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	targetsv1alpha1 "github.com/zeiss/typhoon/pkg/apis/targets/v1alpha1"
)

var (
	validationTypes = map[schema.GroupVersionKind]resourcesemantics.GenericCRD{}
	defaultingTypes = map[schema.GroupVersionKind]resourcesemantics.GenericCRD{
		sourcesv1alpha1.SchemeGroupVersion.WithKind("CloudEventsSource"): &sourcesv1alpha1.CloudEventsSource{},
		routingv1alpha1.SchemeGroupVersion.WithKind("Filter"):            &routingv1alpha1.Filter{},
		flowv1alpha1.SchemeGroupVersion.WithKind("XSLTTransformation"):   &flowv1alpha1.XSLTTransformation{},
	}
)

// NewDefaultingAdmissionController returns defaulting webhook controller implementation.
func NewDefaultingAdmissionController(ctx context.Context, cmw configmap.Watcher) *controller.Impl {
	return defaulting.NewAdmissionController(ctx,
		// Name of the resource webhook.
		"defaulting.webhook.typhoon.zeiss.com",

		// The path on which to serve the webhook.
		"/defaulting",

		// The resources to default.
		defaultingTypes,

		// A function that infuses the context passed to Validate/SetDefaults with custom metadata.
		func(ctx context.Context) context.Context {
			return ctx
		},

		// Whether to disallow unknown fields.
		true,
	)
}

// NewValidationAdmissionController returns validation webhook controller implementation.
func NewValidationAdmissionController(ctx context.Context, cmw configmap.Watcher) *controller.Impl {
	return validation.NewAdmissionController(ctx,
		// Name of the resource webhook.
		"validation.webhook.typhoon.zeiss.com",

		// The path on which to serve the webhook.
		"/validation",

		// The resources to validate.
		validationTypes,

		// A function that infuses the context passed to Validate/SetDefaults with custom metadata.
		func(ctx context.Context) context.Context {
			return ctx
		},

		// Whether to disallow unknown fields.
		true,
	)
}

func main() {
	webhookName := webhook.NameFromEnv()

	// Set up a signal context with our webhook options
	ctx := webhook.WithOptions(signals.NewContext(), webhook.Options{
		ServiceName: webhookName,
		Port:        webhook.PortFromEnv(8443),
		SecretName:  webhookName + "-certs",
	})

	registerValidationType(sourcesv1alpha1.SchemeGroupVersion, sourcesv1alpha1.AllTypes)
	registerValidationType(targetsv1alpha1.SchemeGroupVersion, targetsv1alpha1.AllTypes)
	registerValidationType(flowv1alpha1.SchemeGroupVersion, flowv1alpha1.AllTypes)
	registerValidationType(extensionsv1alpha1.SchemeGroupVersion, extensionsv1alpha1.AllTypes)
	registerValidationType(routingv1alpha1.SchemeGroupVersion, routingv1alpha1.AllTypes)

	sharedmain.MainWithContext(ctx, webhookName,
		certificates.NewController,
		NewDefaultingAdmissionController,
		NewValidationAdmissionController,
	)
}

// registerValidationType registers components in the validation controller.
func registerValidationType(gv schema.GroupVersion, objects []v1alpha1.GroupObject) {
	for _, object := range objects {
		t := reflect.TypeOf(object.Single)
		if admissible, ok := object.Single.(resourcesemantics.GenericCRD); ok {
			validationTypes[gv.WithKind(t.Elem().Name())] = admissible
		}
	}
}
