## Mongo
- Creates a collection for each directory in `collections/`
    - collection name is same as directory name

- Imports data into the collection if the directory has a `data.json` file

[Github Repository](https://github.com/chadgrant/docker-database)

## Sample

```
    ./collections
        - collection - collection name
           - data.json
```

```docker
FROM chadgrant/mongo:3.4

COPY collections /docker-entrypoint-initdb.d/collections
```

## Running locally without Dockerfile

```bash
	docker run -itp 27017-27018:27017-27018 \
            -v $(pwd)/collections:/docker-entrypoint-initdb.d/collections \
			chadgrant/mongo:${BUILD_NUMBER}
```

## Building 
```bash
make docker-build
```

## Running
```bash
make docker-run
```