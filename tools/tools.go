//go:build tools
// +build tools

package tools

import (
	_ "github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen"
	_ "github.com/golang/mock/mockgen/model"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/google/ko"
	_ "github.com/goreleaser/goreleaser"
	_ "gotest.tools/gotestsum"
	_ "k8s.io/code-generator"
	_ "knative.dev/pkg/codegen/cmd/injection-gen"
	_ "mvdan.cc/gofumpt"
)
