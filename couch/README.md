## Couchbase Container
- Initializes couchbase cluster
- Creates a bucket for each directory in `buckets/`
    - bucket's name is same as directory name
- Create indexes if the bucket directory contains an `indexes.txt` file
- Import data into the bucket if the bucket directory has a `data.json` file

## Exporting Couchbase data

```
docker run -it --rm -v $(pwd):/data --entrypoint sh couchbase:enterprise-6.0.0

cbexport json -c couchbase://IPADDRESS -u XXX -p XXX -b lucky_life_qa -o data/lucky_life_qa.json -f lines -t 4 --include-key id
```

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