BUILD_NUMBER?=1.0
BUILDER?=golang:1.13
RUNNER?=openjdk:11.0.5-stretch

.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
.EXPORT_ALL_VARIABLES:

docker-build:
	docker build -t chadgrant/dynamo:${BUILD_NUMBER} --build-arg builder=${BUILDER} runner=${RUNNER} .

docker-push:
	docker push chadgrant/dynamo:${BUILD_NUMBER}

docker-run:
	docker run -itp 8000:8000 chadgrant/dynamo:${BUILD_NUMBER}

docker-stop:
	-docker container stop `docker container ls -q --filter ancestor=chadgrant/dynamo`

docker-rm: docker-stop
	-docker container rm `docker container ls -aq --filter ancestor=chadgrant/dynamo`

docker-clean: docker-rm
	-docker rmi `docker images --format '{{.Repository}}:{{.Tag}}' | grep "chadgrant/dynamo"` -f
	-docker rmi `docker images -qf dangling=true`