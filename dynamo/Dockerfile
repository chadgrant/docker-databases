ARG builder
ARG runner

FROM $builder AS builder

WORKDIR /go/src/github.com/chadgrant/dynamodb-local/
COPY main.go .
RUN go get ./... && \
    go build -o /go/bin/populator

FROM $runner
EXPOSE 8000
EXPOSE 8001
RUN mkdir -p opt/dynamodb
WORKDIR /opt/dynamodb
RUN wget https://s3-us-west-2.amazonaws.com/dynamodb-local/dynamodb_local_latest.tar.gz -q -O - | tar -xz

RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends curl gnupg && \
    curl -sL https://deb.nodesource.com/setup_12.x | bash && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y nodejs && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* && \
    npm install -g dynamodb-admin

COPY entrypoint.sh /entrypoint.sh
COPY --from=builder /go/bin/populator /populator
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
CMD ["java","-jar","DynamoDBLocal.jar","-inMemory"]