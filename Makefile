info:
	@echo "Makefile is your friend"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

start: ## start gin-gonic app example
	@go run src/app/rest/gingonic/gingonic.go

mock: ## generates mocks
	@go install github.com/vektra/mockery/v2@latest
	@mockery --dir=src/repository --all --output=src/repository/mocks
	@mockery --dir=src/service --all --output=src/service/mocks