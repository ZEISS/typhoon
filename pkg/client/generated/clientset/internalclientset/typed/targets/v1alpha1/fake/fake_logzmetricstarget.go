// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/targets/v1alpha1"
	targetsv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/typed/targets/v1alpha1"
	gentype "k8s.io/client-go/gentype"
)

// fakeLogzMetricsTargets implements LogzMetricsTargetInterface
type fakeLogzMetricsTargets struct {
	*gentype.FakeClientWithList[*v1alpha1.LogzMetricsTarget, *v1alpha1.LogzMetricsTargetList]
	Fake *FakeTargetsV1alpha1
}

func newFakeLogzMetricsTargets(fake *FakeTargetsV1alpha1, namespace string) targetsv1alpha1.LogzMetricsTargetInterface {
	return &fakeLogzMetricsTargets{
		gentype.NewFakeClientWithList[*v1alpha1.LogzMetricsTarget, *v1alpha1.LogzMetricsTargetList](
			fake.Fake,
			namespace,
			v1alpha1.SchemeGroupVersion.WithResource("logzmetricstargets"),
			v1alpha1.SchemeGroupVersion.WithKind("LogzMetricsTarget"),
			func() *v1alpha1.LogzMetricsTarget { return &v1alpha1.LogzMetricsTarget{} },
			func() *v1alpha1.LogzMetricsTargetList { return &v1alpha1.LogzMetricsTargetList{} },
			func(dst, src *v1alpha1.LogzMetricsTargetList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.LogzMetricsTargetList) []*v1alpha1.LogzMetricsTarget {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1alpha1.LogzMetricsTargetList, items []*v1alpha1.LogzMetricsTarget) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
