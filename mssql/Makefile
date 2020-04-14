BUILD_NUMBER?=2019-GA-ubuntu-16.04
#2017-latest
#2019-GA-ubuntu-16.04

.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
.EXPORT_ALL_VARIABLES:

build: ## build docker image
	docker build -t chadgrant/mssql:${BUILD_NUMBER} --build-arg tag=${BUILD_NUMBER} .

push: ## push docker image
	docker push chadgrant/mssql:${BUILD_NUMBER}

run: ## run container
	docker run -itp 1433:1433 \
	-v $(shell pwd)/sample/data:/usr/config/setup_data \
	-e ACCEPT_EULA=Y -e SA_PASSWORD=!StrongPassw0rd \
	chadgrant/mssql:${BUILD_NUMBER}

stop: ## stop docker
	-docker container stop `docker container ls -q --filter ancestor=chadgrant/mssql`

clean: stop ## delete containers / images
	-docker container rm `docker container ls -aq --filter ancestor=chadgrant/mssql`
	-docker rmi `docker images --format '{{.Repository}}:{{.Tag}}' | grep "chadgrant/mssql"` -f
	-docker rmi `docker images -qf dangling=true`