## Couchbase Container
- Initializes couchbase cluster
- Creates a bucket for each directory in `buckets/`
    - bucket's name is same as directory name
- Create indexes if the bucket directory contains an `indexes.txt` file
- Import data into the bucket if the bucket directory has a `data.json` file

[Github Repository](https://github.com/chadgrant/docker-database)

## Sample

```
    ./buckets
        - buckets - buckets name
           - data.json
```

```docker
FROM chadgrant/couch:enterprise-6.0.0

COPY buckets /opt/couchbase/buckets
```

## Running locally without Dockerfile

```bash
    docker run -itp 8091-8094:8091-8094 \
        -p 11210-11211:11210-11211 \
		-e CB_USERNAME=docker -e CB_PASSWORD=password \
		-v $(pwd)/sample/buckets:/opt/couchbase/buckets \
		chadgrant/couchbase:enterprise-6.0.0
```

## Exporting Couchbase data

```
docker run -it --rm -v $(pwd):/data --entrypoint sh couchbase:enterprise-6.0.0

cbexport json -c couchbase://IPADDRESS -u XXX -p XXX -b bucketname -o data/data.json -f lines -t 4 --include-key id
```

## Sample

Sample dockerfile: [https://github.com/chadgrant/docker-database/couch/sample](Sample)

## Building 
```bash
make docker-build
```

## Running
```bash
make docker-run
```