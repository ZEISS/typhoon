package servicebustopics

import (
	"fmt"

	coreclientv1 "k8s.io/client-go/kubernetes/typed/core/v1"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/servicebus/mgmt/servicebus"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/servicebus/mgmt/servicebus/servicebusapi"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"

	"github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	"github.com/zeiss/typhoon/pkg/sources/auth/azure"
)

// SubscriptionsClient is an alias for the SubscriptionsClientAPI interface.
type SubscriptionsClient = servicebusapi.SubscriptionsClientAPI

// ClientGetter can obtain clients for Azure Service Bus APIs.
type ClientGetter interface {
	Get(*v1alpha1.AzureServiceBusTopicSource) (SubscriptionsClient, error)
}

// NewClientGetter returns a ClientGetter for the given secrets getter.
func NewClientGetter(sg NamespacedSecretsGetter) *ClientGetterWithSecretGetter {
	return &ClientGetterWithSecretGetter{
		sg: sg,
	}
}

// NamespacedSecretsGetter returns a SecretInterface for the given namespace.
type NamespacedSecretsGetter func(namespace string) coreclientv1.SecretInterface

// ClientGetterWithSecretGetter gets Azure clients using static credentials
// retrieved using a Secret getter.
type ClientGetterWithSecretGetter struct {
	sg NamespacedSecretsGetter
}

// ClientGetterWithSecretGetter implements ClientGetter.
var _ ClientGetter = (*ClientGetterWithSecretGetter)(nil)

// Get implements ClientGetter.
func (g *ClientGetterWithSecretGetter) Get(src *v1alpha1.AzureServiceBusTopicSource) (SubscriptionsClient, error) {
	var authorizer autorest.Authorizer
	var err error

	if src.Spec.Auth.ServicePrincipal != nil {
		authorizer, err = azure.NewAADAuthorizer(g.sg(src.Namespace), src.Spec.Auth.ServicePrincipal)
		if err != nil {
			return nil, fmt.Errorf("retrieving Azure service principal credentials: %w", err)
		}
	} else {
		// Use Azure AKS Managed Identity
		authorizer, err = auth.NewAuthorizerFromEnvironment()
		if err != nil {
			return nil, fmt.Errorf("retrieving Azure AKS Managed Identity: %w", err)
		}
	}

	subsCli := servicebus.NewSubscriptionsClient(src.Spec.TopicID.SubscriptionID)
	subsCli.Authorizer = authorizer

	return subsCli, nil
}

// ClientGetterFunc allows the use of ordinary functions as ClientGetter.
type ClientGetterFunc func(*v1alpha1.AzureServiceBusTopicSource) (SubscriptionsClient, error)

// ClientGetterFunc implements ClientGetter.
var _ ClientGetter = (ClientGetterFunc)(nil)

// Get implements ClientGetter.
func (f ClientGetterFunc) Get(src *v1alpha1.AzureServiceBusTopicSource) (SubscriptionsClient, error) {
	return f(src)
}
