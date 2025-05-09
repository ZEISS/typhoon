// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/targets/v1alpha1"
	targetsv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/typed/targets/v1alpha1"
	gentype "k8s.io/client-go/gentype"
)

// fakeKafkaTargets implements KafkaTargetInterface
type fakeKafkaTargets struct {
	*gentype.FakeClientWithList[*v1alpha1.KafkaTarget, *v1alpha1.KafkaTargetList]
	Fake *FakeTargetsV1alpha1
}

func newFakeKafkaTargets(fake *FakeTargetsV1alpha1, namespace string) targetsv1alpha1.KafkaTargetInterface {
	return &fakeKafkaTargets{
		gentype.NewFakeClientWithList[*v1alpha1.KafkaTarget, *v1alpha1.KafkaTargetList](
			fake.Fake,
			namespace,
			v1alpha1.SchemeGroupVersion.WithResource("kafkatargets"),
			v1alpha1.SchemeGroupVersion.WithKind("KafkaTarget"),
			func() *v1alpha1.KafkaTarget { return &v1alpha1.KafkaTarget{} },
			func() *v1alpha1.KafkaTargetList { return &v1alpha1.KafkaTargetList{} },
			func(dst, src *v1alpha1.KafkaTargetList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.KafkaTargetList) []*v1alpha1.KafkaTarget {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1alpha1.KafkaTargetList, items []*v1alpha1.KafkaTarget) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
