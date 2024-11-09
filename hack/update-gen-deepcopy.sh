#!/bin/bash
source "$(dirname "${BASH_SOURCE}")/lib/init.sh"
source "$(dirname "${BASH_SOURCE}")/lib/generator-help.sh"

go run -modfile ./tools/go.mod k8s.io/code-generator/cmd/deepcopy-gen \
                --go-header-file "hack/copyright.go.txt" \
                --output-file zz_generated.deepcopy.go \
                $(find_dirs_containing_comment_tags "+k8s:deepcopy-gen=")
