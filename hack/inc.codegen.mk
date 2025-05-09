# Code generation
#
# see https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api_changes.md#generate-code

# Name of the Go package for this repository
PKG 		:= github.com/zeiss/typhoon

# List of API groups to generate code for
# e.g. "sources/v1alpha1 sources/v1alpha2"
API_GROUPS 	:= sources/v1alpha1 targets/v1alpha1 flow/v1alpha1 extensions/v1alpha1 routing/v1alpha1
# generates e.g. "PKG/apis/sources/v1alpha1 PKG/apis/sources/v1alpha2"
api-import-paths := $(foreach group,$(API_GROUPS),$(PKG)/pkg/apis/$(group))

generators 	:= deepcopy client lister informer injection

.PHONY: codegen $(generators)
codegen: $(generators)

# http://blog.jgc.org/2007/06/escaping-comma-and-space-in-gnu-make.html
comma := ,
null  :=
space := $(null) $(null)

# Additionally to $(API_GROUPS), generate deepcopy methods for selected shared Go types inside apis/common/...
deepcopy: private api-import-paths += $(PKG)/pkg/apis/common/v1alpha1
deepcopy:
	@echo "+ Generating deepcopy funcs for $(API_GROUPS)"
	$(GO_RUN_TOOLS) k8s.io/code-generator/cmd/deepcopy-gen \
		${api-import-paths} \
		--go-header-file hack/copyright.go.txt \
		--output-file zz_generated.deepcopy.go \
		$(find_dirs_containing_comment_tags "+k8s:deepcopy-gen=")

client:
	@echo "+ Generating clientsets for $(API_GROUPS)"
	@rm -rf pkg/client/generated/clientset
		echo "+ Generating clientsets for $$apigrp" ; \
		$(GO_RUN_TOOLS) k8s.io/code-generator/cmd/client-gen \
			--fake-clientset=true \
			--input $(subst $(space),$(comma),$(API_GROUPS)) \
			--input-base $(PKG)/pkg/apis \
			--go-header-file hack/copyright.go.txt \
			--output-pkg $(PKG)/pkg/client/generated/clientset \
			--output-dir pkg/client/generated/clientset;

lister:
	@echo "+ Generating listers for $(API_GROUPS)"
	@rm -rf pkg/client/generated/listers
		echo "+ Generating listers for $$apigrp" ; \
		$(GO_RUN_TOOLS) k8s.io/code-generator/cmd/lister-gen \
			${api-import-paths} \
			--go-header-file hack/copyright.go.txt \
			--output-pkg $(PKG)/pkg/client/generated/listers \
			--output-dir pkg/client/generated/listers;

informer:
	@echo "+ Generating informers for $(API_GROUPS)"
	@rm -rf pkg/client/generated/informers
		echo "+ Generating informers for $$apigrp" ; \
		$(GO_RUN_TOOLS) k8s.io/code-generator/cmd/informer-gen \
			${api-import-paths} \
			--go-header-file hack/copyright.go.txt \
			--output-pkg $(PKG)/pkg/client/generated/listers \
			--listers-package $(PKG)/pkg/client/generated/listers \
			--versioned-clientset-package $(PKG)/pkg/client/generated/clientset/internalclientset \
			--output-dir pkg/client/generated/listers;

injection:
	@echo "+ Generating injection for $(API_GROUPS)"
	@rm -rf pkg/client/generated/injection
	$(GO_RUN_TOOLS) knative.dev/pkg/codegen/cmd/injection-gen \
		${api-import-paths} \
		--go-header-file hack/copyright.go.txt \
		--input-dirs $(subst $(space),$(comma),$(api-import-paths)) \
		--output-package $(PKG)/pkg/client/generated/injection \
		--versioned-clientset-package $(PKG)/pkg/client/generated/clientset/internalclientset \
		--listers-package $(PKG)/pkg/client/generated/listers \
		--output-dir pkg/client/generated/injection \
		--external-versions-informers-package $(PKG)/pkg/client/generated/listers/externalversions;

# Cleanup codegen
.PHONY: codegen-cleanup
codegen-cleanup:
	@if [ -d "./$(PKG)" ]; then \
		cp -a ./$(PKG)/pkg/client/generated/ pkg/client/generated/ ;\
		cp -a ./$(PKG)/apis/* apis/ ;\
		rm -rf "./$(PKG)" ;\
	fi
