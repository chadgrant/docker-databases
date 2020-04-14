#!/bin/bash

cd /opt/dynamodb/

echo "starting dynamo"
java -jar DynamoDBLocal.jar -sharedDb -dbPath /var/dynamodb &
pid=$!

echo "starting populator"
/populator

echo "killing dynamo"
kill $pid

exit 0