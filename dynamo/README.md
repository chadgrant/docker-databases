# dynamodb local
docker image that loads data / schemas from inside container

# Sample Dockerfile

```
    data 
        - (directory | table name)
            - schema.json
            - anything_not_named_schema_loaded_as_records.json
```

```docker
FROM chadgrant/dynamo:1.3

COPY data /data/

RUN /build.sh
```

## Sample

Sample dockerfile: [https://github.com/chadgrant/docker-database/dynamo/sample](Sample)

## Building 
```bash
make docker-build
```

## Running
```bash
make docker-run
```