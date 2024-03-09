info:
	@echo "Makefile is your friend"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

chmod: ## chmod shell scripts
	chmod +x infra/chmod.sh
	./infra/chmod.sh

start: chmod ## starts service
	./infra/start.sh

docker-up-local: chmod ## ups docker service local
	./infra/local/docker-up.sh

docker-stop-local: chmod ## stops docker service local
	./infra/local/docker-stop.sh

docker-down-local: chmod ## removes docker service local
	./infra/local/docker-down.sh