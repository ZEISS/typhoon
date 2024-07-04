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

	"H4sIAAAAAAAC/6xUTW/jNhD9K8S0R0VysntY6FQDC2yzBYoAdrEFEh8YahSxoUiWHHljBPrvxZCyZdcq",
	"UBS96WM+3rw3895Bud47i5Yi1O/gZZA9Eob8Njz/ggd+ajCqoD1pZ6GGbYfCD89GK/GKB+FaQR0KqZQb",
	"LJVQgOYoL6mDAqzskd9yrQIC/jnogA3UFAYsIKoOe8lN6OA5MlLQ9gXGcTz+TGC+ftsuI/n6bSvIvaLl",
	"zq0LvSSo4Y/vBMVVyQIiqiFoOmy4MqbK0utpzIS7Q9lgmJH/frN+uL/J4Kd6U8ZYwDPKgGE9UMf5CS4H",
	"5M9zQkfk80Tato5DlbMkFfEj9lIbRonPMpKW9qdXSdIY+aaxbHA/Q9kcI8Rnh8YwgEs+fhD3loJrBsUf",
	"nuyT3c7KiIhhj0H44Pa6wZhUO7EXRevCuZBRaCt+XW834q5clU8WCjBaoY2YqGrQkm41Bqhh7aXq8Oau",
	"XM1Y8zfOZZhx6HsZDpNk/wcgJleTwVTy4DvnrFhPdTe57vrhHgrYY4iZndtyldE4j1Z6DTV8KG8TaF7W",
	"tAzVsVnFby9IS0uno0DbeKctCR3FELER5ARhpAT4+BNSqyA58b6BGr4g/YzGpzOI3tmYN/ButbpusxkU",
	"xtgOJtXNWrdyMAuIfrP45lERNgJDcCFv+hnlOjJOmRGeoSP5EqF+hARqx1nz/O/5ZMd/5OEL0rk8gleb",
	"r087e5LuRe/RnplFuUTJJNuWVU9azCb0+A4/Bmx5tavZqqo5pJqMZdwtc8p3hpbynXujVepcsT0cDzaZ",
	"z1KXyX0qtp50vJfzr6/nZpU+rD4uG1XaavFdRmEdid41fD5NyTkf/11O6wabE/7rKrBmC3qdrcI0FW/D",
	"mV0mJc7N7nE3FrN1Pu6Y/3zOWbVLVJ9xj8b5Hk9HDwUMwUzWWFeVcUqazkWqP60+rVifan8L3OOy0sPJ",
	"3f5e6JSzOw2zwGhEIQOm3ezQ+NMxxHK2rnQN160vs48sLhU4cjjuxr8CAAD///RxaZJiBwAA",
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
