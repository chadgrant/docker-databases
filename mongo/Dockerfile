ARG tag
FROM mongo:$tag
MAINTAINER chad.grant@me.com

# in derived projects, create this volume and copy
# COPY collections /docker-entrypoint-initdb.d/collections 

COPY configure.sh /docker-entrypoint-initdb.d/configure.sh