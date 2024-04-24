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

	"H4sIAAAAAAAC/+xd6W/bOBb/VwR1PyqWnO4CA3/LTLtB2kEa1OkU2DQY0NKzzUbXkFQyhuH/fcFLh0VZ",
	"snzkGH+LJB6P7/zx8ZlZ2n4SpUkMMaP2aGmniKAIGBDxhHw/yWJ2FfCHAKhPcMpwEtsj+3YOFg6sZGqx",
	"OViqocUSiwAjGB5hYDs25i1TxOa2Y8coAntUGtKxCfyVYQKBPWIkA8em/hwixOeaJiRCzB7ZWYZ5S7ZI",
	"eWfKCI5n9mrl2DOSZGk7YaJZB7L0cLsRFeIIsxvOQjNdcRZNgFBOHGYQUUVYRuKcrL8yIIuCLjGiXaYi",
	"gCnKQmaPzj3HjtDfOMoie/Qf/oBj+TDMacMxgxkQQVwynVJop65CHH3AqTWBaULAogwRhuMZf+8nYQg+",
	"EywmQLOQWRRY0yLkzOZVlOn2zHSnQBBLSLuwdcsO8i4NupvIGaConTLeqgNVarBdKFrJzkDZr0mAQZjx",
	"bwQQgwtpefyFn8QM5J8oTUPsI061+5Ny0pel2VLCGcXUOJUVLten1qtYGrhUrOdOtrrPKU8mP8FnnPKV",
	"oyj9ooSzA6m8G/LZxwjh0EjrsdZyC9LcOq/jXwSm9sh+5xaO2ZVfqSsGW8nh1TvepSTZLbngCxqDC1ZX",
	"YEE+TmIrQAwsFAcWwxHYjg1/oygN+Tjn3vnwzBuevfduvV9G772R5/3PdgoV5T3PVC+DAEJ4tqk3yx4H",
	"JksjgIIvcbjQNlnr9gCLNgF+hsUNwmSDgjk2xbMYx7PPsBAyFL54i3HVeIgQtODPWRo8j4i7mYpj/xZm",
	"lAGpa29PMWQk7G24jv3p++1t8gBxnVsX1qfvtxbjHy1MLWR9Gn+5tr7DxBIdBj/iMmM+fb81qZ4WVGV0",
	"Mbj6IodOs0mIfTcl+JGL5gEWVoowkXNUuaTafJbKl88/QRTen5tIkGN3bb/Gt6KziXnXatQ1xsUW/8BX",
	"xmNheVHJ1ELW9cXtWCPINS42r6IcI05O7+T0XpnTu0EzHHMSvwr0XNdiCfwLVklwbuc43vhJQnFa+pZz",
	"hCUMhYZOKwNx41wYl3xbZDCxkwWZICWmaBJCUNnfTFFIC7uZJEkIKN7B3jD9k/pJ2nWWZoPjg7RZ2Joi",
	"jEWf12BeRsLr+xnEapvR9Q2oY6doESYo6NCQx0Y678PUG9WVCyabcH5OeglnnHcuRhKSoW3krzpy8aZY",
	"ZJWZKAyTp4rvruncuk8KIF5s06MrjeMyC18GlQvKIDI40gL7bhK1hsgn8JIYgXnpWWdcqOD4wDRYT8+r",
	"XWl18msUQXXOyqohO3sCys6GJkKSEobdJP4c67401+vkCnzfqPTd8ZwykpVIa17JHkPPq1ukzqmccP/z",
	"4f5GYPHiscE3ako2/DOVAxqNpada/HxixuHUNvENqdEfQKiyqRqqNC/pseixeU7d0JFjGfLLjk3Bzwhm",
	"izF3ngrbpPhPxWdxtjAHFAApHb2p74VDTfFnEB51AogAuciYAHby6b/VVJbw0mJ3Ib4Wo8wZS4V7TZIH",
	"DHoMQYF8VVDwafxxPL76cn31oU7ESgwR00ys5q6eJb/nljLFMc7xbCkF7ufJ9q9A0ySm0JRP3CLTX8vq",
	"OzaOp0ldI99ZVzEjSZD5/MWP+Ed8u0jnSSIThbEFjxAza0JwMANrmhCrtDgqNJgCecQ+UJn/CrEPagmK",
	"cxcp8udgnQ88W2U4Bd/pyHWfnp4GSHweJGTmqr7U/f3qt4/X449n5wNvMGdRKJAqZsI+NHkXN1d2STHt",
	"4cAbeAqaxCjF9sh+PxiKSVPE5oKNroYt4mkmUxHyHU7iq8Ae2b9jynLQ4lQOle/MGKBo4pbPKVdOa/PS",
	"mevqXiRHg8xvViGi1EMQf+55W53NoDD8Mm1cg8Yx67kdvoiqJpbyNJ2QURkAtm027oWeVvVznPk+UDrN",
	"wtDSDKh4EbGmshO4E65+VIj6njOXZlGEyEJJ2EJhaBUtHJuhmeB7Tu493xUnlMmEVvwgFnoJuW78upBH",
	"p1XtmUGT8iwrh8GcZXItg0kSLN65OJBGWh1u7Uyxi4ro89PG1GbliNVdm2FV07Lh3k4Ai0nqUpZkBFKw",
	"uaTkW2ohK4anXFxmaa2cknW7y4LZK+nzOE6py+uDeN/f3guRruoG+u+6ty20eW2pkhDhcfNSgMnCuvrA",
	"RWJ0VJfADkS3dxSRN7LiElgDH9LMwIdvAm3tjxUVA9o7F1YvjduSfUaGN5uUq468uoXRC914F9k4p+C7",
	"ffDV1RXHjr25ejSFXt1AIsq4xbebI7FaXGMg1ot3misDTWHY2RyoyzrZFrDNFGztkw4U9EvKcbCYn8/R",
	"L+SjnIFbRfzcPbnLXNhbwIB9iK3d/xRquD/ooGtZOyOHF7BU7xjK1go21ji3pW65qgbg7EEVAXSNieNS",
	"7cDRZHAKpT1CaWPFxoFDaUWzmsKpamTxRjqkbvSd5ohaaGNjUC2abBVXy6XzXcNrubKm2vZd9z1zzc6O",
	"6uq2tIT9xd1cW/vF3ZI+7SX2uqyoS9wiAstqxtcWhlWd5ZoZdg/Gz77q/UXkvCJ1+5DcyEXtuDpY/bNz",
	"cngUTnaxaDM3N5jx1ojmeFDmhE5ePzrps+HfDZ7sDju2SNXvCW/8QyHElmhhdzd/kMiflA5V2kL/IZbw",
	"AsN4mSVd4vgh2PJCY3KFNe3nHG9QYfKTiCbGcI9Bi8I8ZVPV0b+KnyTzQUIefpKpFaEYzSBQFY50YDsG",
	"EKPr/V5p+C/KD/cb/eux3MBN7dsVEY3RW3xujty6ALUateU8W8fhYrTtcud70X4tDdNp3/AAs/QKz5pB",
	"NemVrMxdavaXsucG8ckPRgnebf5xuRzfYomlQr3xt+W5Emz6dfl6DdS+YroicT2lvrPbuWzQ+c4c0z/I",
	"PyjPPFPNeqodWO4uxK8hOdRXfrSv96DSfTTHv91YJisz982w43oQ7wgepEOMrpoFdxoMULQ5VXArWpyK",
	"6naO9/IyhSPv9aV8mzb5TMlW27MgsQkL8I+NSED0XMcB+q6SrVCAHukw5+daCAeL8PrKjF57C7n0NWnk",
	"ZuouJUNXjfZ6aRREu7UqQR10e9HEmNa9KKC6x8pZ0a2oic/du6BJM+eU43y7xUydTa+xSqXJHhuLNQ5z",
	"HZOzPNwlaK+oDKSD0FxxBdtmv3EpmxzOY5wKPHbMoazdHHFsN6N0qMMhirwZcEOhx0bsJVbXCL7k2vuW",
	"duQ3G5oLKxtBnNbyJiBnoupAtnP/TDmjde07cPKoPt2OhzxSKXvHPeVC3aVSoS4JpyNqRXtjrfr7yzut",
	"87a9qvO1MMQ7ph63bQo28HlDHuoFs/qZnNPLEGopQdUk164eKaPQ8jvRb6LFCdK9VEgnrks4No6TatME",
	"48TXnsiNL6cRuIm19sVtbcDMsTndW+ffDDS9LdQmFeywUE3P0QefZVIAPSGZ0FV3KUXfBZAdT97m3IRS",
	"0v0mJvqhN05KO2B7kwzzDq76bZCuyvwNKO6t8P+Ynu0ZxVuAu7KEuTsr3U9jtLU/8mtpDrYaPYW+4zyn",
	"WtcGsDlYfkYIxMxS9OoT6oubq8pZOJBHEDWSq/a4RvUtOnfyOpv7NYhSvtvm7p4rYgW0iDf59Tt3EreI",
	"+U0H6x/gEcIkjfgaZKvKnS4j1w0TH4XzhLLRL94vnotS7D4ORYBWqzOc1VOwEIHyWb0uFJgsLHXfy6Cw",
	"mrzaafNQIsRZbI6YeJVRvhuoj6aO8VrGIng248itNBwE4t9MQBzIS3LEf50wjK+6cpj5/wAAAP//Ru04",
	"GrNkAAA=",
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
