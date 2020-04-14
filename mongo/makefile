BUILD_NUMBER?=3.4

.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
.EXPORT_ALL_VARIABLES:

docker-build:
	docker build -t chadgrant/mongo:${BUILD_NUMBER} --build-arg tag=${BUILD_NUMBER} .

docker-push:
	docker push chadgrant/mongo:${BUILD_NUMBER}

docker-run:
	docker run -itp 27017-27018:27017-27018 \
            -v $(shell pwd)/sample/collections:/docker-entrypoint-initdb.d/collections \
			chadgrant/mongo:${BUILD_NUMBER}

docker-stop:
	-docker container stop `docker container ls -q --filter ancestor=chadgrant/mongo`

docker-rm: docker-stop
	-docker container rm `docker container ls -aq --filter ancestor=chadgrant/mongo`

docker-clean: docker-rm
	-docker rmi `docker images --format '{{.Repository}}:{{.Tag}}' | grep "chadgrant/mongo"` -f
	-docker rmi `docker images -qf dangling=true`