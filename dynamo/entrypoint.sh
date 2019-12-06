#!/bin/bash
set -m

cd /opt/dynamodb/

eval "$@" &

cd /

./populator

fg 1