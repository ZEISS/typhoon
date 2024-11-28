// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	scheme "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// WebhookSourcesGetter has a method to return a WebhookSourceInterface.
// A group's client should implement this interface.
type WebhookSourcesGetter interface {
	WebhookSources(namespace string) WebhookSourceInterface
}

// WebhookSourceInterface has methods to work with WebhookSource resources.
type WebhookSourceInterface interface {
	Create(ctx context.Context, webhookSource *v1alpha1.WebhookSource, opts v1.CreateOptions) (*v1alpha1.WebhookSource, error)
	Update(ctx context.Context, webhookSource *v1alpha1.WebhookSource, opts v1.UpdateOptions) (*v1alpha1.WebhookSource, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, webhookSource *v1alpha1.WebhookSource, opts v1.UpdateOptions) (*v1alpha1.WebhookSource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.WebhookSource, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.WebhookSourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.WebhookSource, err error)
	WebhookSourceExpansion
}

// webhookSources implements WebhookSourceInterface
type webhookSources struct {
	*gentype.ClientWithList[*v1alpha1.WebhookSource, *v1alpha1.WebhookSourceList]
}

// newWebhookSources returns a WebhookSources
func newWebhookSources(c *SourcesV1alpha1Client, namespace string) *webhookSources {
	return &webhookSources{
		gentype.NewClientWithList[*v1alpha1.WebhookSource, *v1alpha1.WebhookSourceList](
			"webhooksources",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.WebhookSource { return &v1alpha1.WebhookSource{} },
			func() *v1alpha1.WebhookSourceList { return &v1alpha1.WebhookSourceList{} }),
	}
}
