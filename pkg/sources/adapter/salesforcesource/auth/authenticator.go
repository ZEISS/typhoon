package auth

// Authenticator manages and stores Authentication Credentials
// for services.
type Authenticator interface {
	// CreateOrRenewCredentials is a best effort wrapper on credential
	// provisioning functions that can decide if it is better to create
	// new credentials or refresh using an existing token.
	CreateOrRenewCredentials() (*Credentials, error)
	// NewCredentials retrieve a new set of credentials.
	NewCredentials() (*Credentials, error)
	// RefreshCredentials uses credentials refresh tokens to create a new set of credentials.
	RefreshCredentials() (*Credentials, error)
}

// Credentials returned from Salesforce Auth.
type Credentials struct {
	CommunityID  string `json:"sfdc_community_id"`
	CommunityURL string `json:"sfdc_community_url"`
	ID           string `json:"id"`
	IDToken      string `json:"id_token"`
	InstanceURL  string `json:"instance_url"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	Signature    string `json:"signature"`
	Token        string `json:"access_token"`
	TokenType    string `json:"token_type"`
}
