#!/bin/bash
set -m

cd /opt/dynamodb/

eval "$@" &

/populator

fg 1