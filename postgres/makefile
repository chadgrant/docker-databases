BUILD_NUMBER?=12.1

.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
.EXPORT_ALL_VARIABLES:

docker-build:
	docker build -t chadgrant/postgres:${BUILD_NUMBER} --build-arg tag=${BUILD_NUMBER} ./sample

docker-run:
	docker run -itp 5432:5432 \
	-e POSTGRES_USER=docker -e POSTGRES_PASSWORD=password -e POSTGRES_DB=sms \
	chadgrant/postgres:${BUILD_NUMBER}

docker-stop:
	-docker container stop `docker container ls -q --filter ancestor=chadgrant/postgres`

docker-rm: docker-stop
	-docker container rm `docker container ls -aq --filter ancestor=chadgrant/postgres`

docker-clean: docker-rm
	-docker rmi `docker images --format '{{.Repository}}:{{.Tag}}' | grep "chadgrant/postgres"` -f
	-docker rmi `docker images -qf dangling=true`