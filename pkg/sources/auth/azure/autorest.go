package azure

import (
	"errors"
	"fmt"

	coreclientv1 "k8s.io/client-go/kubernetes/typed/core/v1"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	autorestauth "github.com/Azure/go-autorest/autorest/azure/auth"

	"github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	"github.com/zeiss/typhoon/pkg/sources/auth"
	"github.com/zeiss/typhoon/pkg/sources/secret"
)

// NewAADAuthorizer returns a new Authorizer for autorest-based Azure clients
// using the provided Service Principal authentication information.
func NewAADAuthorizer(cli coreclientv1.SecretInterface, spAuth *v1alpha1.AzureServicePrincipal) (autorest.Authorizer, error) {
	if spAuth == nil {
		return nil, errors.New("servicePrincipal auth is undefined")
	}

	requestedSecrets, err := secret.NewGetter(cli).Get(
		spAuth.TenantID,
		spAuth.ClientID,
		spAuth.ClientSecret,
	)
	if err != nil {
		return nil, fmt.Errorf("getting auth secrets: %w", err)
	}

	tenantID, clientID, clientSecret := requestedSecrets[0], requestedSecrets[1], requestedSecrets[2]

	azureEnv := &azure.PublicCloud

	authSettings := autorestauth.EnvironmentSettings{
		Values: map[string]string{
			autorestauth.TenantID:     tenantID,
			autorestauth.ClientID:     clientID,
			autorestauth.ClientSecret: clientSecret,
			autorestauth.Resource:     azureEnv.ResourceManagerEndpoint,
		},
		Environment: *azureEnv,
	}

	authorizer, err := authSettings.GetAuthorizer()
	if err != nil {
		// GetAuthorizer returns an untyped error when it is unable to
		// obtain a non-empty value for any of the required auth settings.
		return nil, auth.NewPermanentCredentialsError(err)
	}

	return authorizer, nil
}
