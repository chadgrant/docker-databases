# redis
docker image that loads data / keys from inside container

To restore a dump.rdb simply use:

```docker
COPY dump.rdb /data/dump.rdb
```

/import/*.* = files that act as scripts for example:

/import/simple-keys.txt would contain:

```
SET Key0 Value0
SET Key1 Value1
SET Key2 Value2
SET Key3 Value3
```

### SET
    /import/SET/ = SET [filename - extension] [file contents]
### SADD
    /import/SADD/directory/ = SADD [directory name] [file contents]
### ZADD
    /import/ZADD/directory/ = ZADD [directory name] [int filename sort] [file contents]
### LPUSH
    /import/LPUSH/directory/ = LPUSH [directory name] [file contents]
### HSET 
    /import/HSET/directory/ = HSET [directory name] [filename - extension] [file contents]


# Redis-Stat and Redis-Commander

Redis-stat is exposed on 63790
Redis-commander is exposed on 8080

# Sample Docker-compose

```yaml
  redis:
    image: chadgrant/redis:5.0.7
    restart: unless-stopped
    ports: 
      - 6379:6379
      - 63790:63790
      - 8080:8080
```

# Sample Dockerfile

```docker
FROM chadgrant/redis:1.0
COPY dump.rdb /data/dump.rdb
COPY import /import/
```

## Sample

Sample dockerfile and scripts: [https://github.com/chadgrant/docker-database/redis/sample](Sample)

## Building 
```bash
make docker-build
```

## Running
```bash
make docker-run
```