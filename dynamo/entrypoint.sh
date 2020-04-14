#!/bin/bash
set -m

cd /opt/dynamodb/

eval "$@" &

if [ "${POPULATE}x" = "x" ]; then
    echo "not populating data, POPULATE env var not set"
else
    /populator
fi

DYNAMO_ENDPOINT=http://localhost:8000 dynamodb-admin &

fg 1