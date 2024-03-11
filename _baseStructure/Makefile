info:
	@echo "Makefile is your friend"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

chmod: ## chmod shell scripts
	chmod +x infra/chmod.sh
	./infra/chmod.sh

start: chmod ## starts service
	./infra/start.sh

docker-up: chmod ## ups docker service
	./infra/docker-up.sh

docker-stop: chmod ## stops docker service
	./infra/docker-stop.sh

docker-down: chmod ## removes docker service
	./infra/docker-down.sh