## Mysql

This is meant more as an inspiration since mysql already loads the sql files in this directory, this is a sample meant for reference

## Sample Dockerfile

```docker
FROM mysql:8.0.18

COPY data /docker-entrypoint-initdb.d/
```

### Running locally

```bash
	docker run -itp 3306:3306 \
	-v $(pwd)/sample/data:/docker-entrypoint-initdb.d \
	-e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=sms \
	mysql:8.0.18
```