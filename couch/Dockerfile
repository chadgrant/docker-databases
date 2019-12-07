ARG tag
FROM couchbase/server:$tag
MAINTAINER chad.grant@me.com

# in derived projects, create this volume and copy
#buckets /opt/couchbase/buckets

COPY *.sh /

RUN chmod +x /*.sh 

ENTRYPOINT ["/configure.sh"]
CMD ["couchbase-server"]