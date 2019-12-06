FROM golang:1.13 AS builder
MAINTAINER chad.grant@me.com
WORKDIR /go/src/github.com/chadgrant/dynamodb-local/
COPY main.go .
RUN go get ./... && \
    go build -o /go/bin/populator

FROM openjdk:11.0.5-stretch
EXPOSE 8000
RUN mkdir -p opt/dynamodb
WORKDIR /opt/dynamodb
RUN wget https://s3-us-west-2.amazonaws.com/dynamodb-local/dynamodb_local_latest.tar.gz -q -O - | tar -xz

COPY entrypoint.sh /entrypoint.sh
COPY --from=builder /go/bin/populator /populator
RUN chmod +x /entrypoint.sh

#ONBUILD COPY schema /schema
#ONBUILD COPY data /data

ENTRYPOINT ["/entrypoint.sh"]
CMD ["java","-jar","DynamoDBLocal.jar","-inMemory"]