ARG tag
FROM mcr.microsoft.com/mssql/server:$tag
MAINTAINER chad.grant@me.com

RUN mkdir -p /usr/config
WORKDIR /usr/config

COPY entrypoint.sh entrypoint.sh
RUN chmod +X /usr/config/entrypoint.sh

#COPY data setup_data

ENTRYPOINT ["bash","./entrypoint.sh"]
CMD ["/opt/mssql/bin/sqlservr"]