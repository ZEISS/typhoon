// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	sourcesv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/typed/sources/v1alpha1"
	gentype "k8s.io/client-go/gentype"
)

// fakeCloudEventsSources implements CloudEventsSourceInterface
type fakeCloudEventsSources struct {
	*gentype.FakeClientWithList[*v1alpha1.CloudEventsSource, *v1alpha1.CloudEventsSourceList]
	Fake *FakeSourcesV1alpha1
}

func newFakeCloudEventsSources(fake *FakeSourcesV1alpha1, namespace string) sourcesv1alpha1.CloudEventsSourceInterface {
	return &fakeCloudEventsSources{
		gentype.NewFakeClientWithList[*v1alpha1.CloudEventsSource, *v1alpha1.CloudEventsSourceList](
			fake.Fake,
			namespace,
			v1alpha1.SchemeGroupVersion.WithResource("cloudeventssources"),
			v1alpha1.SchemeGroupVersion.WithKind("CloudEventsSource"),
			func() *v1alpha1.CloudEventsSource { return &v1alpha1.CloudEventsSource{} },
			func() *v1alpha1.CloudEventsSourceList { return &v1alpha1.CloudEventsSourceList{} },
			func(dst, src *v1alpha1.CloudEventsSourceList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.CloudEventsSourceList) []*v1alpha1.CloudEventsSource {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1alpha1.CloudEventsSourceList, items []*v1alpha1.CloudEventsSource) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
