NAME=web-1

.PHONY: install-helm
install-helm:
	@brew install helm

.PHONY: lint
lint:
	@helm lint webapp

.PHONY: template
template:
	@helm template "$(NAME)" webapp --output-dir output

.PHONY: deploy-web-1
deploy-web-1:
	@helm upgrade web-1 webapp --install --values web-1.yaml

.PHONY: deploy-web-2
deploy-web-2:
	@helm upgrade web-2 webapp --install --values web-2.yaml

.PHONY: delete
delete:
	@helm uninstall web-1
	@helm uninstall web-2
