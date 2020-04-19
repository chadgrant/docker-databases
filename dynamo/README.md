# dynamodb local
docker image that loads data / schemas from inside container

### Sample Data

```
    data 
        - (directory | table name)
            - schema.json
            - anything_not_named_schema_loaded_as_records.json
```

Json can be formatted as a large blob of arrays:
```json
[
    { 
        "inner": {}
    },
    {
        "inner": {}
    }
]
```

or line by line in these formats:

```json
[
    {},
    {}
]
```


```json
{},
{}
```

```json
{}
{}
```

```docker
FROM chadgrant/dynamo:1.4

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