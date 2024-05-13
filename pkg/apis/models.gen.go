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
	Description *string            `json:"description,omitempty"`
	Name        string             `json:"name"`
	OperatorId  openapi_types.UUID `json:"operatorId"`
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
	OperatorId *openapi_types.UUID `json:"operatorId,omitempty"`
}

// UpdateAccount defines model for UpdateAccount.
type UpdateAccount struct {
	Claims      *JWTAccountClaims `json:"claims,omitempty"`
	Description *string           `json:"description,omitempty"`
	Name        string            `json:"name"`
}

// UpdateOperatorSystemAccount defines model for UpdateOperatorSystemAccount.
type UpdateOperatorSystemAccount struct {
	AccountId openapi_types.UUID `json:"accountId"`
}

// UpdateSystemOperator defines model for UpdateSystemOperator.
type UpdateSystemOperator struct {
	OperatorId openapi_types.UUID `json:"operatorId"`
}

// ListAccountsJSONBody defines parameters for ListAccounts.
type ListAccountsJSONBody struct {
	OperatorId *openapi_types.UUID `json:"operatorId,omitempty"`
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
	Description *string            `json:"description,omitempty"`
	Name        string             `json:"name"`
	OperatorId  openapi_types.UUID `json:"operatorId"`
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

// UpdateOperatorSystemAccountJSONBody defines parameters for UpdateOperatorSystemAccount.
type UpdateOperatorSystemAccountJSONBody struct {
	AccountId openapi_types.UUID `json:"accountId"`
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

// UpdateOperatorSystemAccountJSONRequestBody defines body for UpdateOperatorSystemAccount for application/json ContentType.
type UpdateOperatorSystemAccountJSONRequestBody UpdateOperatorSystemAccountJSONBody

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
