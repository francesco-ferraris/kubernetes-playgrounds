.PHONY: install
requirements:
	@brew install kubectl kind

.PHONY: create
create:
	@kind create cluster --config config/kind.yaml
	@kubectl cluster-info --context kind-kind
	@kubectl apply -f https://kind.sigs.k8s.io/examples/ingress/deploy-ingress-nginx.yaml

.PHONY: connect
connect:
	@kubectl cluster-info --context kind-kind

.PHONY: delete-cluster
delete-cluster:
	@kind delete cluster

.PHONY: build
build:
	@docker build -t helloworld:latest .

.PHONY: load
load: build
	@kind load docker-image helloworld:latest

.PHONY: deploy
deploy: load
	@kubectl apply -f manifests

.PHONY: delete
delete:
	@kubectl delete -f manifests
