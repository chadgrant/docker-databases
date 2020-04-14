FROM redis:5.0.7

RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y ruby-full build-essential nodejs npm && \
    gem install redis-stat && \
    npm install -g redis-commander && \
    apt-get remove build-essential -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

COPY *.sh /
RUN chmod +x /*.sh
EXPOSE 8080
EXPOSE 63790
EXPOSE 6379
ENTRYPOINT [ "/entrypoint.sh" ]
CMD [ "redis-server"," --appendonly yes" ]
