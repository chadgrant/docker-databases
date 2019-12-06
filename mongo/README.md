## Mongo Container
- Creates a collection for each directory in `collections/`
    - collection name is same as directory name

- Import data into the collection if the directory has a `data.json` file

## Sample

Sample dockerfile: [sample](Sample)

## Building 
```bash
make docker-build
```

## Running
```bash
make docker-run
```