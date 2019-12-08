# SQL Server
Docker containers that load data from disk for testing purposes

[Github Repository](https://github.com/chadgrant/docker-database)

## Conventions

Creates database of each directory name, then executes sql files in that directory against that database

```
/usr/config/setup_data
                *- DB_NAME - directory
                    - 01-sql-file.sql
```

## Sample Dockerfile

directory(used as database name) containing sql files

```
./databasename/
    - 01-sql-file.sql
```

```docker
FROM chadgrant/mssql:2019-GA-ubuntu-16.04

COPY databasename /usr/config/setup_data/databasename
```

## Running without dockerfile

./sample/data directory with a directory(used as database name) containing sql files

```
./sample/data/
        *- DB_NAME - directory
            - 01-sql-file.sql
```

```bash
	docker run -itp 1433:1433 \
	-v $(pwd)/sample/data:/usr/config/setup_data \
	-e ACCEPT_EULA=Y -e SA_PASSWORD=!StrongPassw0rd \
	chadgrant/mssql:${BUILD_NUMBER}

```

## Building 
```bash
make docker-build
```

## Running
```bash
make docker-run
```