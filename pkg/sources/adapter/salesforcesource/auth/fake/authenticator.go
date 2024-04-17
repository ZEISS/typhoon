package fake

import (
	"errors"

	"github.com/zeiss/typhoon/pkg/sources/adapter/salesforcesource/auth"
)

// Authenticator is a test oriented fake Authenticator.
type Authenticator struct {
	defaultCredentials auth.Credentials
}

var _ auth.Authenticator = (*Authenticator)(nil)

// NewFakeAuthenticator creates a Fake authenticator for Bayeux/Salesforce.
func NewFakeAuthenticator(defaultCredentials auth.Credentials) auth.Authenticator {
	return &Authenticator{
		defaultCredentials: defaultCredentials,
	}
}

// NewCredentials generates a new set of credentials.
func (a *Authenticator) NewCredentials() (*auth.Credentials, error) {
	return &a.defaultCredentials, nil
}

// RefreshCredentials renews credentials.
func (a *Authenticator) RefreshCredentials() (*auth.Credentials, error) {
	return nil, errors.New("not implemented")
}

// CreateOrRenewCredentials will always create a new set of credentials.
func (a *Authenticator) CreateOrRenewCredentials() (*auth.Credentials, error) {
	return a.NewCredentials()
}
