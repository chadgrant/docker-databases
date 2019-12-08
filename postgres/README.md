## Postgres

This is meant more as an inspiration since postgres already loads the sql files in this directory, this is a sample meant for reference

## Sample Dockerfile

```docker
FROM postgres:12.1

COPY data /docker-entrypoint-initdb.d/
```

### Running locally

```bash
	docker run -itp 3306:3306 \
	-v $(pwd)/sample/data:/docker-entrypoint-initdb.d \
	-e POSTGRES_USER=docker -e POSTGRES_PASSWORD=password -e POSTGRES_DB=sms \
	postgres:12.1
```