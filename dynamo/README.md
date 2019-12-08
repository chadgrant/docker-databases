# dynamodb local
docker image that loads data / schemas from inside container

# Sample Dockerfile

```
    data 
        - (directory | table name)
            - schema.json
            - data.json
```

```docker
FROM chadgrant/dynamo:1.0

COPY data /data/
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