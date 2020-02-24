#!/bin/bash

set -m

eval "$@" &

sleep 1

DATA_DIR=/import

## scripts: executes commands stored in a file
files=($DATA_DIR/*.*)
for ((i=0; i<${#files[@]}; i++)); do
    f=${files[$i]}
    name=$(basename "${files[$i]}")
    echo ""
    echo ">>> Importing ${name} ..."
    cat $f | redis-cli --pipe
    echo ""
done

## text blobs does a SET operation on the contents 
## of the file, using the filename as the key
files=($DATA_DIR/SET/*)
for ((i=0; i<${#files[@]}; i++)); do
    f=${files[$i]}
    name=$(basename "${files[$i]}")
    echo ""
    echo ">>> Importing ${f} as ${name%%.*} ..."
    redis-cli -x SET ${name%%.*} <$f
done

## text blobs does a SADD operation on the contents 
## of the file, using the filename as the key
dirs=($DATA_DIR/SADD/*/)
for ((i=0; i<${#dirs[@]}; i++)); do
    d=${dirs[$i]}
    dname=$(basename "${dirs[$i]}")
    files=(${d}*)
    for ((j=0; j<${#files[@]}; j++)); do
        f=${files[$j]}
        name=$(basename "${files[$j]}")
        echo ">>> Importing ${f} into set ${dname} ..."
        redis-cli -x SADD $dname <$f
    done
    echo ""
done

## text blobs does a ZADD operation on the contents 
## of the file, using the filename as the key
dirs=($DATA_DIR/ZADD/*/)
for ((i=0; i<${#dirs[@]}; i++)); do
    d=${dirs[$i]}
    dname=$(basename "${dirs[$i]}")
    files=(${d}*)
    for ((j=0; j<${#files[@]}; j++)); do
        f=${files[$j]}
        name=$(basename "${files[$j]}")
        echo ">>> Importing ${f} into sorted set ${dname} in position $j"
        redis-cli -x ZADD $dname $j <$f
    done
    echo ""
done

## text blobs does a LPUSH operation on the contents 
## of the file, using the filename as the key
dirs=($DATA_DIR/LPUSH/*/)
for ((i=0; i<${#dirs[@]}; i++)); do
    d=${dirs[$i]}
    dname=$(basename "${dirs[$i]}")
    files=(${d}*)
    for ((j=0; j<${#files[@]}; j++)); do
        f=${files[$j]}
        name=$(basename "${files[$j]}")
        echo ">>> Importing ${f} into list ${dname} ..."
        redis-cli -x LPUSH $dname <$f
    done
    echo ""
done

## text blobs does a HSET operation on the contents 
## of the file, using the filename as the key
dirs=($DATA_DIR/HSET/*/)
for ((i=0; i<${#dirs[@]}; i++)); do
    d=${dirs[$i]}
    dname=$(basename "${dirs[$i]}")
    files=(${d}*)
    for ((j=0; j<${#files[@]}; j++)); do
        f=${files[$j]}
        name=$(basename "${files[$j]}")
        echo ">>> Importing ${f} into hash ${dname} as ${name%%.*} ..."
        redis-cli -x HSET $dname ${name%%.*} <$f
    done
    echo ""
done

redis-commander -p 8080  2>&1 &
redis-stat --server > /dev/null 2>&1 &

fg %1