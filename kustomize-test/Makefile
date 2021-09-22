KUSTOMIZE_VERSION = 4.1.3
BIN_DIR := $(shell pwd)/bin

KUSTOMIZE := $(shell pwd)/bin/kustomize
.PHONY: kustomize
kustomize: $(KUSTOMIZE) ## Download kustomize locally if necessary.

$(KUSTOMIZE):
	mkdir -p $(BIN_DIR)
	curl -fsL https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2Fv$(KUSTOMIZE_VERSION)/kustomize_v$(KUSTOMIZE_VERSION)_linux_amd64.tar.gz | \
	tar -C bin -xzf -

.PHONY: manifests
manifests: kustomize ## Generate WebhookConfiguration, ClusterRole and CustomResourceDefinition objects.
	$(KUSTOMIZE) build . > result.yaml
