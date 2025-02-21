// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/targets/v1alpha1"
	targetsv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/typed/targets/v1alpha1"
	gentype "k8s.io/client-go/gentype"
)

// fakeServiceNowTargets implements ServiceNowTargetInterface
type fakeServiceNowTargets struct {
	*gentype.FakeClientWithList[*v1alpha1.ServiceNowTarget, *v1alpha1.ServiceNowTargetList]
	Fake *FakeTargetsV1alpha1
}

func newFakeServiceNowTargets(fake *FakeTargetsV1alpha1, namespace string) targetsv1alpha1.ServiceNowTargetInterface {
	return &fakeServiceNowTargets{
		gentype.NewFakeClientWithList[*v1alpha1.ServiceNowTarget, *v1alpha1.ServiceNowTargetList](
			fake.Fake,
			namespace,
			v1alpha1.SchemeGroupVersion.WithResource("servicenowtargets"),
			v1alpha1.SchemeGroupVersion.WithKind("ServiceNowTarget"),
			func() *v1alpha1.ServiceNowTarget { return &v1alpha1.ServiceNowTarget{} },
			func() *v1alpha1.ServiceNowTargetList { return &v1alpha1.ServiceNowTargetList{} },
			func(dst, src *v1alpha1.ServiceNowTargetList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.ServiceNowTargetList) []*v1alpha1.ServiceNowTarget {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1alpha1.ServiceNowTargetList, items []*v1alpha1.ServiceNowTarget) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
