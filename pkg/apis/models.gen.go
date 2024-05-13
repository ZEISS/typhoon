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
	"wLdTlcv+fwAAAP//SYaHJz9sAAA=",
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
