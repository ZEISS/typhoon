package main

import (
	"os"

	"knative.dev/pkg/injection"
	"knative.dev/pkg/injection/sharedmain"
	"knative.dev/pkg/signals"

	"github.com/zeiss/typhoon/pkg/extensions/reconciler/function"
	"github.com/zeiss/typhoon/pkg/flow/reconciler/jqtransformation"
	"github.com/zeiss/typhoon/pkg/flow/reconciler/synchronizer"
	"github.com/zeiss/typhoon/pkg/flow/reconciler/transformation"
	"github.com/zeiss/typhoon/pkg/flow/reconciler/xmltojsontransformation"
	"github.com/zeiss/typhoon/pkg/flow/reconciler/xslttransformation"
	"github.com/zeiss/typhoon/pkg/routing/reconciler/filter"
	"github.com/zeiss/typhoon/pkg/routing/reconciler/splitter"
	"github.com/zeiss/typhoon/pkg/sources/reconciler/cloudeventssource"
	"github.com/zeiss/typhoon/pkg/sources/reconciler/httppollersource"
	"github.com/zeiss/typhoon/pkg/sources/reconciler/kafkasource"
	"github.com/zeiss/typhoon/pkg/sources/reconciler/webhooksource"
	"github.com/zeiss/typhoon/pkg/targets/reconciler/cloudeventstarget"
	"github.com/zeiss/typhoon/pkg/targets/reconciler/httptarget"
	"github.com/zeiss/typhoon/pkg/targets/reconciler/logzmetricstarget"
	"github.com/zeiss/typhoon/pkg/targets/reconciler/logztarget"
	"github.com/zeiss/typhoon/pkg/targets/reconciler/splunktarget"
)

func main() {
	ctx := signals.NewContext()

	if namespace, set := os.LookupEnv("WORKING_NAMESPACE"); set {
		ctx = injection.WithNamespaceScope(ctx, namespace)
	}

	sharedmain.MainWithContext(ctx, "typhoon-controller",
		cloudeventssource.NewController,
		cloudeventstarget.NewController,
		httppollersource.NewController,
		httptarget.NewController,
		kafkasource.NewController,
		logzmetricstarget.NewController,
		logztarget.NewController,
		splunktarget.NewController,
		webhooksource.NewController,
		// flow
		jqtransformation.NewController,
		synchronizer.NewController,
		transformation.NewController,
		xmltojsontransformation.NewController,
		xslttransformation.NewController,
		// extensions
		function.NewController,
		// routing
		filter.NewController,
		splitter.NewController,
	)
}
