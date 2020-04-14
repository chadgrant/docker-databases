BUILD_NUMBER?=enterprise-6.0.0

.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
.EXPORT_ALL_VARIABLES:

build:
	docker build -t chadgrant/couchbase:${BUILD_NUMBER} --build-arg tag=${BUILD_NUMBER} .

push:
	docker push chadgrant/couchbase:${BUILD_NUMBER}

run:
	docker run -itp 8091-8094:8091-8094 \
        -p 11210-11211:11210-11211 \
		-e CB_USERNAME=docker -e CB_PASSWORD=password \
		-v $(shell pwd)/sample/buckets:/opt/couchbase/buckets \
		chadgrant/couchbase:${BUILD_NUMBER}

stop:
	-docker container stop `docker container ls -q --filter ancestor=chadgrant/couchbase`

clean: stop
	-docker container rm `docker container ls -aq --filter ancestor=chadgrant/couchbase`
	-docker rmi `docker images --format '{{.Repository}}:{{.Tag}}' | grep "chadgrant/couchbase"` -f
	-docker rmi `docker images -qf dangling=true`