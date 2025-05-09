// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/flow/v1alpha1"
	flowv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/typed/flow/v1alpha1"
	gentype "k8s.io/client-go/gentype"
)

// fakeBridges implements BridgeInterface
type fakeBridges struct {
	*gentype.FakeClientWithList[*v1alpha1.Bridge, *v1alpha1.BridgeList]
	Fake *FakeFlowV1alpha1
}

func newFakeBridges(fake *FakeFlowV1alpha1, namespace string) flowv1alpha1.BridgeInterface {
	return &fakeBridges{
		gentype.NewFakeClientWithList[*v1alpha1.Bridge, *v1alpha1.BridgeList](
			fake.Fake,
			namespace,
			v1alpha1.SchemeGroupVersion.WithResource("bridges"),
			v1alpha1.SchemeGroupVersion.WithKind("Bridge"),
			func() *v1alpha1.Bridge { return &v1alpha1.Bridge{} },
			func() *v1alpha1.BridgeList { return &v1alpha1.BridgeList{} },
			func(dst, src *v1alpha1.BridgeList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.BridgeList) []*v1alpha1.Bridge { return gentype.ToPointerSlice(list.Items) },
			func(list *v1alpha1.BridgeList, items []*v1alpha1.Bridge) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
