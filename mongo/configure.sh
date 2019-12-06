#!/bin/bash

COLLECTION_DIR=/docker-entrypoint-initdb.d/collections

collections=($COLLECTION_DIR/*)

for ((i=0; i<${#collections[@]}; i++)); do
    c=$(basename ${collections[$i]})
    echo "creating collection $c ..."
    mongoimport --db ${MONGO_INITDB_DATABASE:-$c} --collection $c --type json --file $COLLECTION_DIR/$c/data.json --jsonArray
done
