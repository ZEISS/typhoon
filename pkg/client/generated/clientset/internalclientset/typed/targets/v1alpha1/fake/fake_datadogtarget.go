// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/targets/v1alpha1"
	targetsv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/typed/targets/v1alpha1"
	gentype "k8s.io/client-go/gentype"
)

// fakeDatadogTargets implements DatadogTargetInterface
type fakeDatadogTargets struct {
	*gentype.FakeClientWithList[*v1alpha1.DatadogTarget, *v1alpha1.DatadogTargetList]
	Fake *FakeTargetsV1alpha1
}

func newFakeDatadogTargets(fake *FakeTargetsV1alpha1, namespace string) targetsv1alpha1.DatadogTargetInterface {
	return &fakeDatadogTargets{
		gentype.NewFakeClientWithList[*v1alpha1.DatadogTarget, *v1alpha1.DatadogTargetList](
			fake.Fake,
			namespace,
			v1alpha1.SchemeGroupVersion.WithResource("datadogtargets"),
			v1alpha1.SchemeGroupVersion.WithKind("DatadogTarget"),
			func() *v1alpha1.DatadogTarget { return &v1alpha1.DatadogTarget{} },
			func() *v1alpha1.DatadogTargetList { return &v1alpha1.DatadogTargetList{} },
			func(dst, src *v1alpha1.DatadogTargetList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.DatadogTargetList) []*v1alpha1.DatadogTarget {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1alpha1.DatadogTargetList, items []*v1alpha1.DatadogTarget) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
