BUILD_NUMBER?=12.1

.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
.EXPORT_ALL_VARIABLES:

build: ## build
	docker build -t chadgrant/postgres_sample:${BUILD_NUMBER} --build-arg tag=${BUILD_NUMBER} --no-cache ./sample

push: ## push
	docker push chadgrant/postgres_sample:${BUILD_NUMBER}

run: ## run
	docker run -itp 5432:5432 \
	-e POSTGRES_USER=docker -e POSTGRES_PASSWORD=password -e POSTGRES_DB=sms \
	chadgrant/postgres_sample:${BUILD_NUMBER}

stop: ## stop
	-docker container stop `docker container ls -q --filter ancestor=chadgrant/postgres`

clean: ## clean
	-docker container rm `docker container ls -aq --filter ancestor=chadgrant/postgres`
	-docker rmi `docker images --format '{{.Repository}}:{{.Tag}}' | grep "chadgrant/postgres"` -f
	-docker rmi `docker images -qf dangling=true`