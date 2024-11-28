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

// CloudEventsTargetsGetter has a method to return a CloudEventsTargetInterface.
// A group's client should implement this interface.
type CloudEventsTargetsGetter interface {
	CloudEventsTargets(namespace string) CloudEventsTargetInterface
}

// CloudEventsTargetInterface has methods to work with CloudEventsTarget resources.
type CloudEventsTargetInterface interface {
	Create(ctx context.Context, cloudEventsTarget *v1alpha1.CloudEventsTarget, opts v1.CreateOptions) (*v1alpha1.CloudEventsTarget, error)
	Update(ctx context.Context, cloudEventsTarget *v1alpha1.CloudEventsTarget, opts v1.UpdateOptions) (*v1alpha1.CloudEventsTarget, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, cloudEventsTarget *v1alpha1.CloudEventsTarget, opts v1.UpdateOptions) (*v1alpha1.CloudEventsTarget, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.CloudEventsTarget, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.CloudEventsTargetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CloudEventsTarget, err error)
	CloudEventsTargetExpansion
}

// cloudEventsTargets implements CloudEventsTargetInterface
type cloudEventsTargets struct {
	*gentype.ClientWithList[*v1alpha1.CloudEventsTarget, *v1alpha1.CloudEventsTargetList]
}

// newCloudEventsTargets returns a CloudEventsTargets
func newCloudEventsTargets(c *TargetsV1alpha1Client, namespace string) *cloudEventsTargets {
	return &cloudEventsTargets{
		gentype.NewClientWithList[*v1alpha1.CloudEventsTarget, *v1alpha1.CloudEventsTargetList](
			"cloudeventstargets",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.CloudEventsTarget { return &v1alpha1.CloudEventsTarget{} },
			func() *v1alpha1.CloudEventsTargetList { return &v1alpha1.CloudEventsTargetList{} }),
	}
}
