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

// JiraTargetsGetter has a method to return a JiraTargetInterface.
// A group's client should implement this interface.
type JiraTargetsGetter interface {
	JiraTargets(namespace string) JiraTargetInterface
}

// JiraTargetInterface has methods to work with JiraTarget resources.
type JiraTargetInterface interface {
	Create(ctx context.Context, jiraTarget *v1alpha1.JiraTarget, opts v1.CreateOptions) (*v1alpha1.JiraTarget, error)
	Update(ctx context.Context, jiraTarget *v1alpha1.JiraTarget, opts v1.UpdateOptions) (*v1alpha1.JiraTarget, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, jiraTarget *v1alpha1.JiraTarget, opts v1.UpdateOptions) (*v1alpha1.JiraTarget, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.JiraTarget, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.JiraTargetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.JiraTarget, err error)
	JiraTargetExpansion
}

// jiraTargets implements JiraTargetInterface
type jiraTargets struct {
	*gentype.ClientWithList[*v1alpha1.JiraTarget, *v1alpha1.JiraTargetList]
}

// newJiraTargets returns a JiraTargets
func newJiraTargets(c *TargetsV1alpha1Client, namespace string) *jiraTargets {
	return &jiraTargets{
		gentype.NewClientWithList[*v1alpha1.JiraTarget, *v1alpha1.JiraTargetList](
			"jiratargets",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.JiraTarget { return &v1alpha1.JiraTarget{} },
			func() *v1alpha1.JiraTargetList { return &v1alpha1.JiraTargetList{} }),
	}
}
