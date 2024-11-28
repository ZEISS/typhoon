// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/targets/v1alpha1"
	scheme "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// NatsTargetsGetter has a method to return a NatsTargetInterface.
// A group's client should implement this interface.
type NatsTargetsGetter interface {
	NatsTargets(namespace string) NatsTargetInterface
}

// NatsTargetInterface has methods to work with NatsTarget resources.
type NatsTargetInterface interface {
	Create(ctx context.Context, natsTarget *v1alpha1.NatsTarget, opts v1.CreateOptions) (*v1alpha1.NatsTarget, error)
	Update(ctx context.Context, natsTarget *v1alpha1.NatsTarget, opts v1.UpdateOptions) (*v1alpha1.NatsTarget, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, natsTarget *v1alpha1.NatsTarget, opts v1.UpdateOptions) (*v1alpha1.NatsTarget, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.NatsTarget, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.NatsTargetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NatsTarget, err error)
	NatsTargetExpansion
}

// natsTargets implements NatsTargetInterface
type natsTargets struct {
	*gentype.ClientWithList[*v1alpha1.NatsTarget, *v1alpha1.NatsTargetList]
}

// newNatsTargets returns a NatsTargets
func newNatsTargets(c *TargetsV1alpha1Client, namespace string) *natsTargets {
	return &natsTargets{
		gentype.NewClientWithList[*v1alpha1.NatsTarget, *v1alpha1.NatsTargetList](
			"natstargets",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.NatsTarget { return &v1alpha1.NatsTarget{} },
			func() *v1alpha1.NatsTargetList { return &v1alpha1.NatsTargetList{} }),
	}
}
