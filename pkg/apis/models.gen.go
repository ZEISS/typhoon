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

// ListAccounts defines model for ListAccounts.
type ListAccounts struct {
	SystemId *openapi_types.UUID `json:"systemId,omitempty"`
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

// ListAccountsJSONBody defines parameters for ListAccounts.
type ListAccountsJSONBody struct {
	SystemId *openapi_types.UUID `json:"systemId,omitempty"`
}

// ListAccountsParams defines parameters for ListAccounts.
type ListAccountsParams struct {
	// Offset The number of items to skip before starting to collect the result set.
	Offset *OffsetParam `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit The numbers of items to return.
	Limit *LimitParam `form:"limit,omitempty" json:"limit,omitempty"`
}

// CreateAccountJSONBody defines parameters for CreateAccount.
type CreateAccountJSONBody struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name"`
}

// ListAccountSigningKeyGroupsParams defines parameters for ListAccountSigningKeyGroups.
type ListAccountSigningKeyGroupsParams struct {
	// Offset The number of items to skip before starting to collect the result set.
	Offset *OffsetParam `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit The numbers of items to return.
	Limit *LimitParam `form:"limit,omitempty" json:"limit,omitempty"`
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

// ListOperatorSigningKeyGroupsParams defines parameters for ListOperatorSigningKeyGroups.
type ListOperatorSigningKeyGroupsParams struct {
	// Offset The number of items to skip before starting to collect the result set.
	Offset *OffsetParam `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit The numbers of items to return.
	Limit *LimitParam `form:"limit,omitempty" json:"limit,omitempty"`
}

// ListSystemsParams defines parameters for ListSystems.
type ListSystemsParams struct {
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

// ListTeamSystemsParams defines parameters for ListTeamSystems.
type ListTeamSystemsParams struct {
	// Offset The number of items to skip before starting to collect the result set.
	Offset *OffsetParam `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit The numbers of items to return.
	Limit *LimitParam `form:"limit,omitempty" json:"limit,omitempty"`
}

// ListAccountsJSONRequestBody defines body for ListAccounts for application/json ContentType.
type ListAccountsJSONRequestBody ListAccountsJSONBody

// CreateAccountJSONRequestBody defines body for CreateAccount for application/json ContentType.
type CreateAccountJSONRequestBody CreateAccountJSONBody

// UpdateAccountJSONRequestBody defines body for UpdateAccount for application/json ContentType.
type UpdateAccountJSONRequestBody = Account

// CreateAccountSigningKeyGroupJSONRequestBody defines body for CreateAccountSigningKeyGroup for application/json ContentType.
type CreateAccountSigningKeyGroupJSONRequestBody = SigningKeyGroup

// CreateOperatorJSONRequestBody defines body for CreateOperator for application/json ContentType.
type CreateOperatorJSONRequestBody CreateOperatorJSONBody

// UpdateOperatorJSONRequestBody defines body for UpdateOperator for application/json ContentType.
type UpdateOperatorJSONRequestBody = Operator

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

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"ymT/NwAA//+CqmpsAWwAAA==",
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
