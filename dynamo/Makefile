BUILD_NUMBER?=1.4
BUILDER?=golang:1.14.2
RUNNER?=openjdk:slim-buster

.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
.EXPORT_ALL_VARIABLES:

build: ## build docker image
	docker build -t chadgrant/dynamo:${BUILD_NUMBER} --build-arg builder=${BUILDER} --build-arg runner=${RUNNER} .

push: ## push docker image
	docker push chadgrant/dynamo:${BUILD_NUMBER}

run: ## run docker image
	docker run -itp 8000-8001:8000-8001 --rm \
	-v $(shell pwd)/sample/data:/data \
	chadgrant/dynamo:${BUILD_NUMBER}

stop: ## stop container
	-docker container stop `docker container ls -q --filter ancestor=chadgrant/dynamo:${BUILD_NUMBER}`

clean: stop # remove image after stopping all containers
	-docker container rm `docker container ls -aq --filter ancestor=chadgrant/dynamo:${BUILD_NUMBER}`
	-docker rmi `docker images --format '{{.Repository}}:{{.Tag}}' | grep "chadgrant/dynamo"` -f
	-docker rmi `docker images -qf dangling=true`