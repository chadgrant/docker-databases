#!/bin/bash
set -m

cd /opt/dynamodb/

eval "$@" &

/populator

DYNAMO_ENDPOINT=http://localhost:8000 dynamodb-admin &

fg 1