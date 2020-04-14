BUILD_NUMBER?=5.0.7

.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
.EXPORT_ALL_VARIABLES:

docker-build: ## build docker image
	docker build -t chadgrant/redis:${BUILD_NUMBER} .

docker-push: ## push docker image
	docker push chadgrant/redis:${BUILD_NUMBER}

docker-run: ## run docker image
	cd sample; docker build -t chadgrant/redis-sample .
	docker run -itp 6379:6379 -p 63790:63790 -p 8080:8080 --rm \
	chadgrant/redis-sample

docker-stop: ## stop container
	-docker container stop `docker container ls -q --filter ancestor=chadgrant/redis:${BUILD_NUMBER}`

docker-rm: docker-stop # remove container after stopping
	-docker container rm `docker container ls -aq --filter ancestor=chadgrant/redis:${BUILD_NUMBER}`

docker-clean: docker-rm # remove image after stopping all containers
	-docker rmi `docker images --format '{{.Repository}}:{{.Tag}}' | grep "chadgrant/redis"` -f
	-docker rmi `docker images -qf dangling=true`