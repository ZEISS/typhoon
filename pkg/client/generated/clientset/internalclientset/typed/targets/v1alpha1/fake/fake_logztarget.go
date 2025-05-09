// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/targets/v1alpha1"
	targetsv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/typed/targets/v1alpha1"
	gentype "k8s.io/client-go/gentype"
)

// fakeLogzTargets implements LogzTargetInterface
type fakeLogzTargets struct {
	*gentype.FakeClientWithList[*v1alpha1.LogzTarget, *v1alpha1.LogzTargetList]
	Fake *FakeTargetsV1alpha1
}

func newFakeLogzTargets(fake *FakeTargetsV1alpha1, namespace string) targetsv1alpha1.LogzTargetInterface {
	return &fakeLogzTargets{
		gentype.NewFakeClientWithList[*v1alpha1.LogzTarget, *v1alpha1.LogzTargetList](
			fake.Fake,
			namespace,
			v1alpha1.SchemeGroupVersion.WithResource("logztargets"),
			v1alpha1.SchemeGroupVersion.WithKind("LogzTarget"),
			func() *v1alpha1.LogzTarget { return &v1alpha1.LogzTarget{} },
			func() *v1alpha1.LogzTargetList { return &v1alpha1.LogzTargetList{} },
			func(dst, src *v1alpha1.LogzTargetList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.LogzTargetList) []*v1alpha1.LogzTarget { return gentype.ToPointerSlice(list.Items) },
			func(list *v1alpha1.LogzTargetList, items []*v1alpha1.LogzTarget) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
