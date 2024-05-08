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
	ApiKeyScopes     = "apiKey.Scopes"
	BearerAuthScopes = "bearerAuth.Scopes"
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
	// CreatedAt Creation date and time
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Delete date and time
	DeletedAt   *time.Time          `json:"deletedAt,omitempty"`
	Description *string             `json:"description,omitempty"`
	Id          *openapi_types.UUID `json:"id,omitempty"`
	Name        string              `json:"name"`
	ServerURL   string              `json:"serverURL"`

	// UpdatedAt Update date and time
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// Clusters defines model for Clusters.
type Clusters = []Cluster

// Credentials defines model for Credentials.
type Credentials = openapi_types.File

// Error defines model for Error.
type Error struct {
	// Code Error code
	Code int32 `json:"code"`

	// Message Error message
	Message string `json:"message"`

	// Ref Reference to the error
	Ref *string `json:"ref,omitempty"`
}

// JWTAccountClaims defines model for JWTAccountClaims.
type JWTAccountClaims struct {
	Exports *[]JWTExport `json:"exports,omitempty"`
}

// JWTExport defines model for JWTExport.
type JWTExport struct {
	AccountTokenPosition *uint    `json:"account_token_position,omitempty"`
	Info                 *JWTInfo `json:"info,omitempty"`
	Name                 *string  `json:"name,omitempty"`
	ResponseType         *string  `json:"response_type,omitempty"`
	Subject              *string  `json:"subject,omitempty"`
	Type                 *int     `json:"type,omitempty"`
}

// JWTInfo defines model for JWTInfo.
type JWTInfo struct {
	Description *string `json:"description,omitempty"`
	InfoUrl     *string `json:"info_url,omitempty"`
}

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

// Operators defines model for Operators.
type Operators struct {
	Results *[]Operator `json:"results,omitempty"`
}

// PaginatedResult defines model for PaginatedResult.
type PaginatedResult struct {
	Limit   *int           `json:"limit,omitempty"`
	Offset  *int           `json:"offset,omitempty"`
	Results *[]interface{} `json:"results,omitempty"`
	Total   *int           `json:"total,omitempty"`
}

// SigningKeyGroup defines model for SigningKeyGroup.
type SigningKeyGroup struct {
	// CreatedAt Creation date and time
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Creation date and time
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Description A description of the signing key group.
	Description *string             `json:"description,omitempty"`
	Disabled    *bool               `json:"disabled,omitempty"`
	Id          *openapi_types.UUID `json:"id,omitempty"`
	IsScoped    *bool               `json:"is_scoped,omitempty"`

	// Name The name of the signing key group. This is used to identify the group.
	Name string `json:"name"`

	// Scope A Signing Key Group Limits is a set of rules that define
	// what a user that is signed with this key can do.
	Scope *SigningKeyGroupLimits `json:"scope,omitempty"`

	// UpdatedAt Creation date and time
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// SigningKeyGroupLimits A Signing Key Group Limits is a set of rules that define
// what a user that is signed with this key can do.
type SigningKeyGroupLimits struct {
	// AllowedConnectionTypes The allowed connection types.
	AllowedConnectionTypes *[]string `json:"allowed_connection_types,omitempty"`

	// BearerToken The bearer token that is required to use this key.
	BearerToken *bool `json:"bearer_token,omitempty"`

	// Data The maximum size of the data in bytes.
	Data *int `json:"data,omitempty"`

	// Payload The maximum size of the payload in bytes.
	Payload *int `json:"payload,omitempty"`
	Pub     *struct {
		Allow *[]string `json:"allow,omitempty"`
		Deny  *[]string `json:"deny,omitempty"`
	} `json:"pub,omitempty"`
	Resp *struct {
		Max *int `json:"max,omitempty"`
		Ttl *int `json:"ttl,omitempty"`
	} `json:"resp,omitempty"`

	// Src The list of allowed sources.
	Src *[]string `json:"src,omitempty"`
	Sub *struct {
		Allow *[]string `json:"allow,omitempty"`
		Deny  *[]string `json:"deny,omitempty"`
	} `json:"sub,omitempty"`

	// Subs The maximum number of subscriptions.
	Subs *int `json:"subs,omitempty"`

	// Tags Tags are a way to group keys together. This can be used
	// to apply a policy to a group of keys.
	Tags *[]string `json:"tags,omitempty"`

	// Times The times that the key is allowed to be used.
	Times *struct {
		End   *time.Time `json:"end,omitempty"`
		Start *time.Time `json:"start,omitempty"`
	} `json:"times,omitempty"`
	TimesLocations *string `json:"times_locations,omitempty"`
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
	Clusters Clusters `json:"clusters"`

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

// Teams defines model for Teams.
type Teams struct {
	Results *[]Team `json:"results,omitempty"`
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

// SystemId defines model for systemId.
type SystemId = openapi_types.UUID

// TeamId defines model for teamId.
type TeamId = openapi_types.UUID

// UserId defines model for userId.
type UserId = openapi_types.UUID

// BadRequest defines model for BadRequest.
type BadRequest = Error

// Duplicate defines model for Duplicate.
type Duplicate = Error

// InternalError defines model for InternalError.
type InternalError = Error

// NotFound defines model for NotFound.
type NotFound = Error

// Unauthorized defines model for Unauthorized.
type Unauthorized = Error

// Unimplemented defines model for Unimplemented.
type Unimplemented = Error

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

// CreateSystem defines model for CreateSystem.
type CreateSystem = System

// CreateTeam defines model for CreateTeam.
type CreateTeam struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name"`
}

// UpdateAccount defines model for UpdateAccount.
type UpdateAccount struct {
	Claims      *JWTAccountClaims `json:"claims,omitempty"`
	Description *string           `json:"description,omitempty"`
	Name        string            `json:"name"`
}

// UpdateSystemOperator defines model for UpdateSystemOperator.
type UpdateSystemOperator struct {
	OperatorId openapi_types.UUID `json:"operatorId"`
}

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

// UpdateOperatorAccountJSONBody defines parameters for UpdateOperatorAccount.
type UpdateOperatorAccountJSONBody struct {
	Claims      *JWTAccountClaims `json:"claims,omitempty"`
	Description *string           `json:"description,omitempty"`
	Name        string            `json:"name"`
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

// UpdateSystemOperatorJSONBody defines parameters for UpdateSystemOperator.
type UpdateSystemOperatorJSONBody struct {
	OperatorId openapi_types.UUID `json:"operatorId"`
}

// ListTeamsParams defines parameters for ListTeams.
type ListTeamsParams struct {
	// Offset The number of items to skip before starting to collect the result set.
	Offset *OffsetParam `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit The numbers of items to return.
	Limit *LimitParam `form:"limit,omitempty" json:"limit,omitempty"`
}

// CreateTeamJSONBody defines parameters for CreateTeam.
type CreateTeamJSONBody struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name"`
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

// ListTeamSystemsParams defines parameters for ListTeamSystems.
type ListTeamSystemsParams struct {
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

// UpdateOperatorAccountJSONRequestBody defines body for UpdateOperatorAccount for application/json ContentType.
type UpdateOperatorAccountJSONRequestBody UpdateOperatorAccountJSONBody

// CreateOperatorAccountUserJSONRequestBody defines body for CreateOperatorAccountUser for application/json ContentType.
type CreateOperatorAccountUserJSONRequestBody CreateOperatorAccountUserJSONBody

// CreateSystemJSONRequestBody defines body for CreateSystem for application/json ContentType.
type CreateSystemJSONRequestBody = System

// UpdateSystemJSONRequestBody defines body for UpdateSystem for application/json ContentType.
type UpdateSystemJSONRequestBody = System

// UpdateSystemOperatorJSONRequestBody defines body for UpdateSystemOperator for application/json ContentType.
type UpdateSystemOperatorJSONRequestBody UpdateSystemOperatorJSONBody

// CreateTeamJSONRequestBody defines body for CreateTeam for application/json ContentType.
type CreateTeamJSONRequestBody CreateTeamJSONBody

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
	"BBqYr4sI9dMxvu2qIPE/AQAA//8lZBlDhpYAAA==",
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
