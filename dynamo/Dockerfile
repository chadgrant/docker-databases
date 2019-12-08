ARG builder
ARG runner

FROM $builder AS builder

WORKDIR /go/src/github.com/chadgrant/dynamodb-local/
COPY main.go .
RUN go get ./... && \
    go build -o /go/bin/populator

FROM $runner
MAINTAINER chad.grant@me.com
EXPOSE 8000
RUN mkdir -p opt/dynamodb
WORKDIR /opt/dynamodb
RUN wget https://s3-us-west-2.amazonaws.com/dynamodb-local/dynamodb_local_latest.tar.gz -q -O - | tar -xz

COPY entrypoint.sh /entrypoint.sh
COPY --from=builder /go/bin/populator /populator
RUN chmod +x /entrypoint.sh

#ONBUILD COPY data /data

ENTRYPOINT ["/entrypoint.sh"]
CMD ["java","-jar","DynamoDBLocal.jar","-inMemory"]