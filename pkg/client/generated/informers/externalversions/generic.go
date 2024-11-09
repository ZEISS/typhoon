// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/extensions/v1alpha1"
	flowv1alpha1 "github.com/zeiss/typhoon/pkg/apis/flow/v1alpha1"
	routingv1alpha1 "github.com/zeiss/typhoon/pkg/apis/routing/v1alpha1"
	sourcesv1alpha1 "github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	targetsv1alpha1 "github.com/zeiss/typhoon/pkg/apis/targets/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=extensions.typhoon.zeiss.com, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("functions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().V1alpha1().Functions().Informer()}, nil

		// Group=flow.typhoon.zeiss.com, Version=v1alpha1
	case flowv1alpha1.SchemeGroupVersion.WithResource("bridges"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Flow().V1alpha1().Bridges().Informer()}, nil
	case flowv1alpha1.SchemeGroupVersion.WithResource("jqtransformations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Flow().V1alpha1().JQTransformations().Informer()}, nil
	case flowv1alpha1.SchemeGroupVersion.WithResource("synchronizers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Flow().V1alpha1().Synchronizers().Informer()}, nil
	case flowv1alpha1.SchemeGroupVersion.WithResource("transformations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Flow().V1alpha1().Transformations().Informer()}, nil
	case flowv1alpha1.SchemeGroupVersion.WithResource("workertransformations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Flow().V1alpha1().WorkerTransformations().Informer()}, nil
	case flowv1alpha1.SchemeGroupVersion.WithResource("xmltojsontransformations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Flow().V1alpha1().XMLToJSONTransformations().Informer()}, nil
	case flowv1alpha1.SchemeGroupVersion.WithResource("xslttransformations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Flow().V1alpha1().XSLTTransformations().Informer()}, nil

		// Group=routing.typhoon.zeiss.com, Version=v1alpha1
	case routingv1alpha1.SchemeGroupVersion.WithResource("filters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Routing().V1alpha1().Filters().Informer()}, nil
	case routingv1alpha1.SchemeGroupVersion.WithResource("splitters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Routing().V1alpha1().Splitters().Informer()}, nil

		// Group=sources.typhoon.zeiss.com, Version=v1alpha1
	case sourcesv1alpha1.SchemeGroupVersion.WithResource("azureservicebusqueuesources"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Sources().V1alpha1().AzureServiceBusQueueSources().Informer()}, nil
	case sourcesv1alpha1.SchemeGroupVersion.WithResource("azureservicebussources"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Sources().V1alpha1().AzureServiceBusSources().Informer()}, nil
	case sourcesv1alpha1.SchemeGroupVersion.WithResource("azureservicebustopicsources"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Sources().V1alpha1().AzureServiceBusTopicSources().Informer()}, nil
	case sourcesv1alpha1.SchemeGroupVersion.WithResource("cloudeventssources"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Sources().V1alpha1().CloudEventsSources().Informer()}, nil
	case sourcesv1alpha1.SchemeGroupVersion.WithResource("httppollersources"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Sources().V1alpha1().HTTPPollerSources().Informer()}, nil
	case sourcesv1alpha1.SchemeGroupVersion.WithResource("kafkasources"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Sources().V1alpha1().KafkaSources().Informer()}, nil
	case sourcesv1alpha1.SchemeGroupVersion.WithResource("ocimetricssources"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Sources().V1alpha1().OCIMetricsSources().Informer()}, nil
	case sourcesv1alpha1.SchemeGroupVersion.WithResource("salesforcesources"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Sources().V1alpha1().SalesforceSources().Informer()}, nil
	case sourcesv1alpha1.SchemeGroupVersion.WithResource("webhooksources"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Sources().V1alpha1().WebhookSources().Informer()}, nil

		// Group=targets.typhoon.zeiss.com, Version=v1alpha1
	case targetsv1alpha1.SchemeGroupVersion.WithResource("cloudeventstargets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Targets().V1alpha1().CloudEventsTargets().Informer()}, nil
	case targetsv1alpha1.SchemeGroupVersion.WithResource("datadogtargets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Targets().V1alpha1().DatadogTargets().Informer()}, nil
	case targetsv1alpha1.SchemeGroupVersion.WithResource("httptargets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Targets().V1alpha1().HTTPTargets().Informer()}, nil
	case targetsv1alpha1.SchemeGroupVersion.WithResource("jiratargets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Targets().V1alpha1().JiraTargets().Informer()}, nil
	case targetsv1alpha1.SchemeGroupVersion.WithResource("kafkatargets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Targets().V1alpha1().KafkaTargets().Informer()}, nil
	case targetsv1alpha1.SchemeGroupVersion.WithResource("logzmetricstargets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Targets().V1alpha1().LogzMetricsTargets().Informer()}, nil
	case targetsv1alpha1.SchemeGroupVersion.WithResource("logztargets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Targets().V1alpha1().LogzTargets().Informer()}, nil
	case targetsv1alpha1.SchemeGroupVersion.WithResource("natstargets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Targets().V1alpha1().NatsTargets().Informer()}, nil
	case targetsv1alpha1.SchemeGroupVersion.WithResource("salesforcetargets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Targets().V1alpha1().SalesforceTargets().Informer()}, nil
	case targetsv1alpha1.SchemeGroupVersion.WithResource("servicenowtargets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Targets().V1alpha1().ServiceNowTargets().Informer()}, nil
	case targetsv1alpha1.SchemeGroupVersion.WithResource("splunktargets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Targets().V1alpha1().SplunkTargets().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
