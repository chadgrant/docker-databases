BUILD_NUMBER?=8.0.18

.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
.EXPORT_ALL_VARIABLES:

docker-build:
	docker build -t chadgrant/mysql:${BUILD_NUMBER} --build-arg tag=${BUILD_NUMBER} ./sample

docker-push:
	docker push chadgrant/mysql:${BUILD_NUMBER}

docker-run:
	docker run -itp 3306:3306 \
	-e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=sms \
	chadgrant/mysql:${BUILD_NUMBER}

docker-stop:
	-docker container stop `docker container ls -q --filter ancestor=chadgrant/mysql`

docker-rm: docker-stop
	-docker container rm `docker container ls -aq --filter ancestor=chadgrant/mysql`

docker-clean: docker-rm
	-docker rmi `docker images --format '{{.Repository}}:{{.Tag}}' | grep "chadgrant/mysql"` -f
	-docker rmi `docker images -qf dangling=true`