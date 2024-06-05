//go:build generate
// +build generate

package api

//go:generate go run -modfile ../../tools/go.mod github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen -config config.client.yml api.yml
//go:generate go run -modfile ../../tools/go.mod github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen -config config.server.yml api.yml
//go:generate go run -modfile ../../tools/go.mod github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen -config config.models.yml api.yml
