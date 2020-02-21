BUILD_NUMBER?=1.1
BUILDER?=golang:1.13
RUNNER?=openjdk:11.0.5-stretch

.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
.EXPORT_ALL_VARIABLES:

docker-build: ## build docker image
	docker build -t chadgrant/dynamo:${BUILD_NUMBER} --build-arg builder=${BUILDER} --build-arg runner=${RUNNER} .

docker-push: ## push docker image
	docker push chadgrant/dynamo:${BUILD_NUMBER}

docker-run: ## run docker image
	docker run -itp 8000-8001:8000-8001 --rm \
	-v $(shell pwd)/sample/data:/data \
	chadgrant/dynamo:${BUILD_NUMBER}

docker-stop: ## stop container
	-docker container stop `docker container ls -q --filter ancestor=chadgrant/dynamo:${BUILD_NUMBER}`

docker-rm: docker-stop # remove container after stopping
	-docker container rm `docker container ls -aq --filter ancestor=chadgrant/dynamo:${BUILD_NUMBER}`

docker-clean: docker-rm # remove image after stopping all containers
	-docker rmi `docker images --format '{{.Repository}}:{{.Tag}}' | grep "chadgrant/dynamo"` -f
	-docker rmi `docker images -qf dangling=true`