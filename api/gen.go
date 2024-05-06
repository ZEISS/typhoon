//go:build generate
// +build generate

package api

//go:generate npx -y @redocly/openapi-cli@latest bundle api.yml -o out.yml
//go:generate go run -modfile ../tools/go.mod github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen -config config.client.yml out.yml
//go:generate go run -modfile ../tools/go.mod github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen -config config.server.yml out.yml
//go:generate go run -modfile ../tools/go.mod github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen -config config.models.yml out.yml
