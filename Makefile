.DEFAULT_GOAL := build

GO 				?= go
GO_RUN_TOOLS	?= $(GO) run -modfile ./tools/go.mod
GO_TEST 		?= $(GO_RUN_TOOLS) gotest.tools/gotestsum --format pkgname
GO_RELEASER 	?= $(GO_RUN_TOOLS) github.com/goreleaser/goreleaser
GO_KO 			?= $(GO_RUN_TOOLS) github.com/google/ko
GO_MOD 			?= $(shell ${GO} list -m)
BASE_DIR		?= $(CURDIR)

COMMANDS		:= $(notdir $(wildcard cmd/*))

.PHONY: build
build: $(COMMANDS) ## Build the application.

$(filter-out $(CUSTOM_BUILD_BINARIES), $(COMMANDS)): ## Build artifact
	$(GO) build -ldflags "$(LDFLAGS_STATIC)" -o $(BIN_OUTPUT_DIR)/$@ ./cmd/$@	

.PHONY: generate
generate: ## Generate code.
	$(GO) generate ./...
	$(GO_RUN_TOOLS) github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen -config ./api/config.models.yml ./api/api.yml
	$(GO_RUN_TOOLS) github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen -config ./api/config.client.yml ./api/api.yml
	$(GO_RUN_TOOLS) github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen -config ./api/config.server.yml ./api/api.yml

.PHONY: fmt
fmt: ## Run go fmt against code.
	$(GO_RUN_TOOLS) mvdan.cc/gofumpt -w .

.PHONY: vet
vet: ## Run go vet against code.
	$(GO) vet ./...

.PHONY: test
test: fmt vet ## Run tests.
	mkdir -p .test/reports
	$(GO_TEST) --junitfile .test/reports/unit-test.xml -- -race ./... -count=1 -short -cover -coverprofile .test/reports/unit-test-coverage.out

.PHONY: lint
lint: ## Run lint.
	$(GO_RUN_TOOLS) github.com/golangci/golangci-lint/cmd/golangci-lint run --timeout 5m -c .golangci.yml

.PHONY: clean
clean: ## Remove previous build.
	rm -rf .test .dist
	find . -type f -name '*.gen.go' -exec rm {} +
	git checkout go.mod

.PHONY: deploy
deploy: ## Deploy the application.
	$(GO_KO) resolve -f $(BASE_DIR)/config > $(BASE_DIR)/typhoon.yaml

.PHONY: help
help: ## Display this help screen.
	@grep -E '^[a-z.A-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# codegen
include hack/inc.codegen.mk
