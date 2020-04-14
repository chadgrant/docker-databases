BUILD_NUMBER?=8.0.18

.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
.EXPORT_ALL_VARIABLES:

build: ## build docker image
	docker build -t chadgrant/mysql_sample:${BUILD_NUMBER} --build-arg tag=${BUILD_NUMBER} ./sample

push: ## push docker image
	docker push chadgrant/mysql_sample:${BUILD_NUMBER}

run: ## run container
	docker run -itp 3306:3306 \
	-v $(shell pwd)/sample/data:/docker-entrypoint-initdb.d \
	-e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=sms \
	mysql:${BUILD_NUMBER}

stop: ## stop docker
	-docker container stop `docker container ls -q --filter ancestor=chadgrant/mysql`

clean: ## delete containers / images
	-docker container rm `docker container ls -aq --filter ancestor=chadgrant/mysql`
	-docker rmi `docker images --format '{{.Repository}}:{{.Tag}}' | grep "chadgrant/mysql"` -f
	-docker rmi `docker images -qf dangling=true`