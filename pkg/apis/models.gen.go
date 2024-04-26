// Package apis provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package apis

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	Api_keyScopes    = "api_key.Scopes"
	BearerAuthScopes = "bearerAuth.Scopes"
	CookieAuthScopes = "cookieAuth.Scopes"
)

// Account defines model for Account.
type Account struct {
	ContactEmail *string `json:"contactEmail,omitempty"`

	// CreatedAt Creation date and time
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Creation date and time
	DeletedAt   *time.Time          `json:"deletedAt,omitempty"`
	Description *string             `json:"description,omitempty"`
	Id          *openapi_types.UUID `json:"id,omitempty"`
	Key         *KeyPair            `json:"key,omitempty"`
	Name        string              `json:"name"`
	SigningKeys *[]KeyPair          `json:"signingKeys,omitempty"`

	// UpdatedAt Creation date and time
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// Cluster defines model for Cluster.
type Cluster struct {
	Id  *openapi_types.UUID `json:"id,omitempty"`
	Url *string             `json:"url,omitempty"`
}

// Credentials defines model for Credentials.
type Credentials = openapi_types.File

// JWTToken A JWT token is a JSON Web Token.
type JWTToken = string

// KeyPair defines model for KeyPair.
type KeyPair struct {
	PrivateKey *string `json:"privateKey,omitempty"`
	PublicKey  string  `json:"publicKey"`
}

// NKey An NKey is the private key of a NATS account.
type NKey = string

// Operator defines model for Operator.
type Operator struct {
	ContactEmail *string `json:"contactEmail,omitempty"`

	// CreatedAt Creation date and time
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Creation date and time
	DeletedAt   *time.Time          `json:"deletedAt,omitempty"`
	Description *string             `json:"description,omitempty"`
	Id          *openapi_types.UUID `json:"id,omitempty"`
	Key         *KeyPair            `json:"key,omitempty"`
	Name        string              `json:"name"`
	SigningKeys *[]KeyPair          `json:"signingKeys,omitempty"`

	// UpdatedAt Creation date and time
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// PaginatedResult defines model for PaginatedResult.
type PaginatedResult struct {
	Limit   *float32       `json:"limit,omitempty"`
	Offset  *float32       `json:"offset,omitempty"`
	Results *[]interface{} `json:"results,omitempty"`
	Total   *float32       `json:"total,omitempty"`
}

// SigningKeyGroup defines model for SigningKeyGroup.
type SigningKeyGroup struct {
	// CreatedAt Creation date and time
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Creation date and time
	DeletedAt   *time.Time            `json:"deletedAt,omitempty"`
	Description *string               `json:"description,omitempty"`
	Disabled    *bool                 `json:"disabled,omitempty"`
	Id          *openapi_types.UUID   `json:"id,omitempty"`
	IsScoped    *bool                 `json:"is_scoped,omitempty"`
	Name        string                `json:"name"`
	Scope       *SigningKeyGroupScope `json:"scope,omitempty"`

	// UpdatedAt Creation date and time
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// SigningKeyGroupScope defines model for SigningKeyGroupScope.
type SigningKeyGroupScope struct {
	Data          *int                           `json:"data,omitempty"`
	Payload       *int                           `json:"payload,omitempty"`
	Publish       *SigningKeyGroupScopePublish   `json:"publish,omitempty"`
	Subscribe     *SigningKeyGroupScopeSubscribe `json:"subscribe,omitempty"`
	Subscriptions *int                           `json:"subscriptions,omitempty"`
}

// SigningKeyGroupScopePublish defines model for SigningKeyGroupScopePublish.
type SigningKeyGroupScopePublish struct {
	Allow *[]string `json:"allow,omitempty"`
	Deny  *[]string `json:"deny,omitempty"`
}

// SigningKeyGroupScopeSubscribe defines model for SigningKeyGroupScopeSubscribe.
type SigningKeyGroupScopeSubscribe struct {
	Allow *[]string `json:"allow,omitempty"`
	Deny  *[]string `json:"deny,omitempty"`
}

// System defines model for System.
type System struct {
	Cluster Cluster `json:"cluster"`

	// CreatedAt Creation date and time
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Creation date and time
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Description A description of the system.
	Description *string             `json:"description,omitempty"`
	Id          *openapi_types.UUID `json:"id,omitempty"`

	// Name Name of the system
	Name     string    `json:"name"`
	Operator *Operator `json:"operator,omitempty"`

	// UpdatedAt Creation date and time
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// Systems defines model for Systems.
type Systems = []System

// Team defines model for Team.
type Team struct {
	ContactEmail *string `json:"contactEmail,omitempty"`

	// CreatedAt Creation date and time
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Creation date and time
	DeletedAt   *time.Time          `json:"deletedAt,omitempty"`
	Description *string             `json:"description,omitempty"`
	Id          *openapi_types.UUID `json:"id,omitempty"`
	Name        string              `json:"name"`

	// UpdatedAt Creation date and time
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// User defines model for User.
type User struct {
	// CreatedAt Creation date and time
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Creation date and time
	DeletedAt *time.Time          `json:"deletedAt,omitempty"`
	Email     *string             `json:"email,omitempty"`
	Id        *openapi_types.UUID `json:"id,omitempty"`
	Jwt       *string             `json:"jwt,omitempty"`
	Key       *string             `json:"key,omitempty"`
	Name      string              `json:"name"`

	// UpdatedAt Creation date and time
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// Version defines model for Version.
type Version struct {
	Date    string `json:"date"`
	Version string `json:"version"`
}

// AccountId defines model for accountId.
type AccountId = openapi_types.UUID

// GroupId defines model for groupId.
type GroupId = openapi_types.UUID

// LimitParam defines model for limitParam.
type LimitParam = int

// OffsetParam defines model for offsetParam.
type OffsetParam = int

// OperatorId defines model for operatorId.
type OperatorId = openapi_types.UUID

// TeamId defines model for teamId.
type TeamId = openapi_types.UUID

// UserId defines model for userId.
type UserId = openapi_types.UUID

// CreateAccount defines model for CreateAccount.
type CreateAccount struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name"`
}

// CreateOperator defines model for CreateOperator.
type CreateOperator struct {
	ContactEmail *string `json:"contactEmail,omitempty"`
	Description  *string `json:"description,omitempty"`
	Name         string  `json:"name"`
}

// CreateOperatorAccountUser defines model for CreateOperatorAccountUser.
type CreateOperatorAccountUser struct {
	Email *string `json:"email,omitempty"`
	Name  string  `json:"name"`
}

// CreateTeam defines model for CreateTeam.
type CreateTeam = Team

// ListOperatorParams defines parameters for ListOperator.
type ListOperatorParams struct {
	// Offset The number of items to skip before starting to collect the result set.
	Offset *OffsetParam `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit The numbers of items to return.
	Limit *LimitParam `form:"limit,omitempty" json:"limit,omitempty"`
}

// CreateOperatorJSONBody defines parameters for CreateOperator.
type CreateOperatorJSONBody struct {
	ContactEmail *string `json:"contactEmail,omitempty"`
	Description  *string `json:"description,omitempty"`
	Name         string  `json:"name"`
}

// ListOperatorAccountsParams defines parameters for ListOperatorAccounts.
type ListOperatorAccountsParams struct {
	// Offset The number of items to skip before starting to collect the result set.
	Offset *OffsetParam `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit The numbers of items to return.
	Limit *LimitParam `form:"limit,omitempty" json:"limit,omitempty"`
}

// CreateOperatorAccountJSONBody defines parameters for CreateOperatorAccount.
type CreateOperatorAccountJSONBody struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name"`
}

// ListOperatorAccountSigningKeysParams defines parameters for ListOperatorAccountSigningKeys.
type ListOperatorAccountSigningKeysParams struct {
	// Offset The number of items to skip before starting to collect the result set.
	Offset *OffsetParam `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit The numbers of items to return.
	Limit *LimitParam `form:"limit,omitempty" json:"limit,omitempty"`
}

// ListOperatorAccountUsersParams defines parameters for ListOperatorAccountUsers.
type ListOperatorAccountUsersParams struct {
	// Offset The number of items to skip before starting to collect the result set.
	Offset *OffsetParam `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit The numbers of items to return.
	Limit *LimitParam `form:"limit,omitempty" json:"limit,omitempty"`
}

// CreateOperatorAccountUserJSONBody defines parameters for CreateOperatorAccountUser.
type CreateOperatorAccountUserJSONBody struct {
	Email *string `json:"email,omitempty"`
	Name  string  `json:"name"`
}

// ListOperatorSigningKeysParams defines parameters for ListOperatorSigningKeys.
type ListOperatorSigningKeysParams struct {
	// Offset The number of items to skip before starting to collect the result set.
	Offset *OffsetParam `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit The numbers of items to return.
	Limit *LimitParam `form:"limit,omitempty" json:"limit,omitempty"`
}

// ListTeamsParams defines parameters for ListTeams.
type ListTeamsParams struct {
	// Offset The number of items to skip before starting to collect the result set.
	Offset *OffsetParam `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit The numbers of items to return.
	Limit *LimitParam `form:"limit,omitempty" json:"limit,omitempty"`
}

// ListTeamAccountsParams defines parameters for ListTeamAccounts.
type ListTeamAccountsParams struct {
	// Offset The number of items to skip before starting to collect the result set.
	Offset *OffsetParam `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit The numbers of items to return.
	Limit *LimitParam `form:"limit,omitempty" json:"limit,omitempty"`
}

// ListGroupsParams defines parameters for ListGroups.
type ListGroupsParams struct {
	// Offset The number of items to skip before starting to collect the result set.
	Offset *OffsetParam `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit The numbers of items to return.
	Limit *LimitParam `form:"limit,omitempty" json:"limit,omitempty"`
}

// ListUsersParams defines parameters for ListUsers.
type ListUsersParams struct {
	// Offset The number of items to skip before starting to collect the result set.
	Offset *OffsetParam `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit The numbers of items to return.
	Limit *LimitParam `form:"limit,omitempty" json:"limit,omitempty"`
}

// CreateOperatorJSONRequestBody defines body for CreateOperator for application/json ContentType.
type CreateOperatorJSONRequestBody CreateOperatorJSONBody

// UpdateOperatorJSONRequestBody defines body for UpdateOperator for application/json ContentType.
type UpdateOperatorJSONRequestBody = Operator

// CreateOperatorAccountJSONRequestBody defines body for CreateOperatorAccount for application/json ContentType.
type CreateOperatorAccountJSONRequestBody CreateOperatorAccountJSONBody

// CreateOperatorAccountUserJSONRequestBody defines body for CreateOperatorAccountUser for application/json ContentType.
type CreateOperatorAccountUserJSONRequestBody CreateOperatorAccountUserJSONBody

// CreateSystemJSONRequestBody defines body for CreateSystem for application/json ContentType.
type CreateSystemJSONRequestBody = System

// UpdateSystemJSONRequestBody defines body for UpdateSystem for application/json ContentType.
type UpdateSystemJSONRequestBody = System

// CreateTeamJSONRequestBody defines body for CreateTeam for application/json ContentType.
type CreateTeamJSONRequestBody = Team

// CreateGroupJSONRequestBody defines body for CreateGroup for application/json ContentType.
type CreateGroupJSONRequestBody = SigningKeyGroup

// UpdateGroupJSONRequestBody defines body for UpdateGroup for application/json ContentType.
type UpdateGroupJSONRequestBody = SigningKeyGroup

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody = User

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.
type UpdateUserJSONRequestBody = User

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
