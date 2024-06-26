// Package accounts provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package accounts

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

const (
	ApiKeyScopes     = "apiKey.Scopes"
	BearerAuthScopes = "bearerAuth.Scopes"
)

// JWT The JWT token.
type JWT = string

// PubKey defines model for pubKey.
type PubKey = string

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/6xUTW/bOBD9K8TsHhVJ+TgEOq2BALvOAkUAu0iBxAdaGllsKJIlKSeGof9eDClZdqND",
	"D71Z8sybN/Pe0xFK3RqtUHkHxREMt7xFjzY+ddv/8UC/KnSlFcYLraCAdYPMdFspSvaGB6Zr5htkvCx1",
	"p3wKCQiqMtw3kIDiLdJTxErA4o9OWKyg8LbDBFzZYMtpiD8YqnTeCrWDvu/HPwOZx+f1PJPH5zXz+g0V",
	"Ta61bbmHAr6/e0g+QSbgsOys8IcVIWNA5kYMawbeDfIK7cT829XiaXkVyQ94Q0efwBa5RbvofEP9gS4V",
	"xNdTQ+O9iRsJVWsqLbXyvPT0E1suJLHELXdecPXPG/dcSv4hMK1wP1FZjRXsQaOURODyHn+xpfJWV11J",
	"L17Vq1pPyjCHdo+WGav3okIXVDtdz7Fa23MhHROKfVmsV+wmzdNXBQlIUaJyGE5VofKiFmihgIXhZYNX",
	"N2k+cY3vqJdouq5tuT0Mkv0JQnRc4SUGyINptFZsMeCuIu7iaQkJ7NG6eJ3rNI9stEHFjYACbtPrQJrM",
	"GsyQjcOyY7RsT2936D+b71/05/QYSUvuE1qdqO/EHtVZWIi2NmhD1bKKKAPtNW0duEwhfDnC3xZrkjab",
	"oppNJdkQrH5D0XJGKxddfZPno89Q+ehzI0UZJmcUj9GwIXxzU4b0ZRS9YN7L/Ref96bj3uZ380ENqrJ3",
	"7pjSnrW6IvtUKfXc/V5PrTsVGyqseSdnRPmq8MNg6bFiaK22MfQn95FmM3qRmfiO7j1uBZv+/HMRlDgP",
	"+8umT6ZPx8uG7h/tHFW7ZPWAe5TatHgyPSTQWTl8Goosk7rkstHOF/f5fU76ZPtroBmXSE+ndP8KdOrZ",
	"nJaZuahDxi0GbzYoDUNVGS2Ud+kU3f9QmpnRl93jFecAxhv2m/5nAAAA//9+O3u1YgYAAA==",
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
