FROM couchbase/server:enterprise-6.0.0

# in derived projects, create this volume and copy
#buckets /opt/couchbase/buckets

COPY *.sh /

RUN chmod +x /*.sh 

ENTRYPOINT ["/configure.sh"]
CMD ["couchbase-server"]