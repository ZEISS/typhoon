package auth

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

const (
	// grant type for OAuth JWT.
	// See: https://tools.ietf.org/html/rfc7523#page-10
	grantJWT = "urn:ietf:params:oauth:grant-type:jwt-bearer"

	oauthTokenPath = "/services/oauth2/token"
)

// JWTAuthenticator is the JWT OAuth implementation.
// See: https://help.salesforce.com/articleView?id=remoteaccess_oauth_jwt_flow.htm
type JWTAuthenticator struct {
	authURL string
	signKey *rsa.PrivateKey
	claims  *claims

	client *http.Client
	logger *zap.SugaredLogger
}

type claims struct {
	jwt.RegisteredClaims
}

// Credentials returned from Salesforce Auth.
type Credentials struct {
	Token        string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	InstanceURL  string `json:"instance_url"`
	ID           string `json:"id"`
	Signature    string `json:"signature"`
	Scope        string `json:"scope"`
	IDToken      string `json:"id_token"`
	CommunityURL string `json:"sfdc_community_url"`
	CommunityID  string `json:"sfdc_community_id"`
}

// NewJWTAuthenticator creates an OAuth JWT authenticator for Salesforce.
func NewJWTAuthenticator(certKey, clientID, user, server string, client *http.Client, logger *zap.SugaredLogger) (*JWTAuthenticator, error) {
	audience := strings.TrimSuffix(server, "/")

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(certKey))
	if err != nil {
		return nil, fmt.Errorf("unable to parse PEM private key: %w", err)
	}

	claims := &claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   clientID,
			Subject:  user,
			Audience: jwt.ClaimStrings{audience},
		},
	}

	return &JWTAuthenticator{
		authURL: audience + oauthTokenPath,
		claims:  claims,
		signKey: signKey,
		client:  client,
		logger:  logger,
	}, nil
}

// NewCredentials generates a new set of credentials.
func (j *JWTAuthenticator) NewCredentials(ctx context.Context) (*Credentials, error) {
	// expiry needs to be set to 3 minutes or less
	// See: https://help.salesforce.com/articleView?id=remoteaccess_oauth_jwt_flow.htm
	j.claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 3))

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, j.claims)
	signedToken, err := token.SignedString(j.signKey)
	if err != nil {
		return nil, fmt.Errorf("could not sign JWT token: %w", err)
	}

	form := url.Values{
		"grant_type": []string{grantJWT},
		"assertion":  []string{signedToken},
	}

	req, err := http.NewRequestWithContext(ctx, "POST", j.authURL, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("could not build authentication request: %w", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := j.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not execute authentication request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		msg := fmt.Sprintf("received unexpected status code %d from authentication", res.StatusCode)
		if resb, err := io.ReadAll(res.Body); err == nil {
			msg += ": " + string(resb)
		}
		return nil, errors.New(msg)
	}

	c := &Credentials{}
	err = json.NewDecoder(res.Body).Decode(c)
	if err != nil {
		return nil, fmt.Errorf("could not decode authentication response into credentials: %w", err)
	}

	return c, nil
}

// RefreshCredentials renews credentials.
func (j *JWTAuthenticator) RefreshCredentials() (*Credentials, error) {
	return nil, errors.New("not implemented")
}

// CreateOrRenewCredentials will always create a new set of credentials.
func (j *JWTAuthenticator) CreateOrRenewCredentials(ctx context.Context) (*Credentials, error) {
	return j.NewCredentials(ctx)
}
