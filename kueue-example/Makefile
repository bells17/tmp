KUBERNETES_VERSION := 1.23.5
KIND_VERSION := 0.12.0
KIND_CLUSTER_NAME := kueue-example
CERT_MANAGER_VERSION := 1.8.0

KIND_NODE_IMAGE=kindest/node:v$(KUBERNETES_VERSION)

BINDIR := $(shell pwd)/bin
KUBECTL := $(BINDIR)/kubectl
KIND := $(BINDIR)/kind


.PHONY: all
all: setup

.PHONY: clean
clean: shutdown-kind
	rm -rf $(BINDIR)

.PHONY: setup
setup: $(BINDIR) $(KUBECTL) $(KIND)

$(BINDIR):
	mkdir -p $(BINDIR)

$(KUBECTL):
	curl -o $(KUBECTL) -sSfL https://storage.googleapis.com/kubernetes-release/release/v$(KUBERNETES_VERSION)/bin/linux/amd64/kubectl
	chmod a+x $(KUBECTL)

$(KIND):
	GOBIN=$(BINDIR) go install sigs.k8s.io/kind@v$(KIND_VERSION)

.PHONY: setup-cluster
# setup-cluster: launch-kind
setup-cluster:
	$(KUBECTL) apply -f https://github.com/cert-manager/cert-manager/releases/download/v$(CERT_MANAGER_VERSION)/cert-manager.yaml
	$(KUBECTL) wait --for=condition=available --timeout=120s -n cert-manager deployments/cert-manager-webhook
	$(KUBECTL) apply -k examples/kueue
	$(KUBECTL) wait --for=condition=available --timeout=120s -n kueue-system deployments/kueue-controller-manager
	$(KUBECTL) apply -k examples/configs

.PHONY: launch-kind
launch-kind:
	sed -e "s|@KUBERNETES_VERSION@|$(KUBERNETES_VERSION)|" kueue-cluster-template.yaml > kueue-cluster.yaml
	$(KIND) create cluster \
		--name=$(KIND_CLUSTER_NAME) \
		--image $(KIND_NODE_IMAGE) \
		--config kueue-cluster.yaml

.PHONY: shutdown-kind
shutdown-kind:
	$(KIND) delete cluster --name=$(KIND_CLUSTER_NAME) || true
