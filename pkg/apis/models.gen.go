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
type JWTToken struct {
	Token *string `json:"token,omitempty"`
}

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
type CreateSystem struct {
	// Clusters A list of clusters the system is part of.
	Clusters []Cluster `json:"clusters"`

	// Description A description of the system.
	Description *string `json:"description,omitempty"`

	// Name Name of the system
	Name       string             `json:"name"`
	OperatorId openapi_types.UUID `json:"operatorId"`
}

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

// ListOperatorsParams defines parameters for ListOperators.
type ListOperatorsParams struct {
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

// ListOperatorSigningKeyGroupsParams defines parameters for ListOperatorSigningKeyGroups.
type ListOperatorSigningKeyGroupsParams struct {
	// Offset The number of items to skip before starting to collect the result set.
	Offset *OffsetParam `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit The numbers of items to return.
	Limit *LimitParam `form:"limit,omitempty" json:"limit,omitempty"`
}

// CreateSystemJSONBody defines parameters for CreateSystem.
type CreateSystemJSONBody struct {
	// Clusters A list of clusters the system is part of.
	Clusters []Cluster `json:"clusters"`

	// Description A description of the system.
	Description *string `json:"description,omitempty"`

	// Name Name of the system
	Name       string             `json:"name"`
	OperatorId openapi_types.UUID `json:"operatorId"`
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

// CreateOperatorSigningKeyGroupJSONRequestBody defines body for CreateOperatorSigningKeyGroup for application/json ContentType.
type CreateOperatorSigningKeyGroupJSONRequestBody = SigningKeyGroup

// CreateSystemJSONRequestBody defines body for CreateSystem for application/json ContentType.
type CreateSystemJSONRequestBody CreateSystemJSONBody

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
