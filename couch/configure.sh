#!/bin/bash

set -m

BUCKET_DIR=/opt/couchbase/buckets

buckets=($BUCKET_DIR/*)

# BUCKETS contains a csv list of bucket names to create (they must have configuration in buckets/ dir)
if [ -n "$BUCKETS" ]; then
    buckets=(${BUCKETS//,/ })
fi

# at least 1 bucket name must be provided
if [ ${#buckets[@]} == 0 ]; then
    echo "Please provide a csv of bucket names via BUCKETS environment variable"
    exit 1
fi

# memory for entire cluster (and index memory)
# divide memory evenly across buckets
CLUSTER_RAM=300
BUCKET_RAM=$(($CLUSTER_RAM/${#buckets[@]}))

echo "CLUSTER_RAM: $CLUSTER_RAM mb"
echo "BUCKET_DIR: $BUCKET_DIR"
echo "BUCKET_RAM: $BUCKET_RAM mb"

/entrypoint.sh "$@" &

until $(curl --output /dev/null --silent --head --fail http://127.0.0.1:8091); do
    echo 'waiting on rest api ...'
    sleep 3
done

# configure cluster
couchbase-cli cluster-init --cluster-username=$CB_USERNAME --cluster-password=$CB_PASSWORD \
    --cluster-port=8091 --cluster-ramsize=$CLUSTER_RAM --cluster-index-ramsize=$CLUSTER_RAM \
    --services=data,index,query,fts

# create buckets
for ((i=0; i<${#buckets[@]}; i++)); do
    b=$(basename ${buckets[$i]})
    echo "creating bucket $b ..."
        couchbase-cli bucket-create -c 127.0.0.1 --bucket-type=couchbase --bucket-ramsize=$BUCKET_RAM \
            --bucket=$b -u $CB_USERNAME -p $CB_PASSWORD    

        # code connects to the buckets using a username that is the same name as the bucket
        # this creates that user with the same password as our admin
        couchbase-cli user-manage -c 127.0.0.1 --set --rbac-username $b --rbac-password $CB_PASSWORD --auth-domain local  \
            --roles bucket_admin[$b] -u $CB_USERNAME -p $CB_PASSWORD
done

for attempt in $(seq 5)
do
    curl -s http://127.0.0.1:8093/admin/ping > /dev/null \
    && break

    echo "Waiting for query service..."
    sleep 3
done

# this creates indexes if there is an indexes.txt file within a bucket directory
# it also imports data if there is a data.json file within a bucket directory
# it also creates a full text search index if fts.sh exists within a bucket directory
for ((i=0; i<${#buckets[@]}; i++)); do
    b="$BUCKET_DIR/$(basename ${buckets[$i]})"

    if [ -f "$b/indexes.txt" ]; then
        echo "creating indexes $b/indexes.txt..."
        cbq -e http://127.0.0.1:8093 -u $CB_USERNAME -p $CB_PASSWORD -q=true -f="$b/indexes.txt"
    fi
    if [ -f "$b/data.json" ]; then
        echo "importing data $b/data.json ..."
        cbimport json -v -u $CB_USERNAME -p $CB_PASSWORD -c 127.0.0.1:8091 -b "$(basename ${buckets[$i]})" -f list -g key::%_id% -d "file://$b/data.json"
    fi
    if [ -f "$b/fts.sh" ]; then
        echo "adding full text search index $b/fts.sh ..."
        bash $b/fts.sh
    fi
done

fg 1
