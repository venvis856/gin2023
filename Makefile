GO_VERSION ?= 1.18.3
GOOS ?= linux
GOARCH ?= amd64
GOPATH ?= $(shell go env GOPATH)

.DEFAULT_GOAL := help
.PHONY: help
help:
	@awk 'BEGIN { \
		FS = ":.*##"; \
		printf "\nUsage:\n  make \033[36m<target>\033[0m\n"\
	} \
	/^[a-zA-Z_-]+:.*?##/ { \
		printf "  \033[36m%-17s\033[0m %s\n", $$1, $$2 \
	} \
	/^##@/ { \
		printf "\n\033[1m%s\033[0m\n", substr($$0, 5) \
	} ' $(MAKEFILE_LIST)

.PHONY: release
release: system_api http ##构建所有发布版本包


.PHONY: system_api
system_api: ## 编译构建system_api
	@echo "Running ${@}"
	./script/release.sh build -o release/system_api ./cmd/system_api/main.go

.PHONY: http
http: ## 编译构建system_api
	@echo "Running ${@}"
	./script/release.sh build -o release/http ./cmd/http/main.go