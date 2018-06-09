FROM ubuntu:16.04

USER root

ENV PG_VERSION 9.6

RUN echo "deb http://apt.postgresql.org/pub/repos/apt/ xenial-pgdg main" > /etc/apt/sources.list.d/pgdg.list

RUN apt-get -y update && apt-get install -y wget

RUN wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add -

RUN apt-get -y update && apt-get install -y postgresql-$PG_VERSION

USER postgres

RUN /etc/init.d/postgresql start &&\
    psql -c "ALTER USER postgres WITH PASSWORD 'postgres';" &&\
    psql -c "CREATE DATABASE tpdb;" &&\
    psql -d tpdb -c "CREATE EXTENSION IF NOT EXISTS citext;" &&\
    psql -c "GRANT ALL PRIVILEGES ON DATABASE tpdb TO postgres;" &&\
    /etc/init.d/postgresql stop

RUN echo "host all  all    0.0.0.0/0  md5" >>\
/etc/postgresql/$PG_VERSION/main/pg_hba.conf

RUN echo "listen_addresses='*'" >> /etc/postgresql/$PG_VERSION/main/postgresql.conf
RUN echo "synchronous_commit = off" >> /etc/postgresql/$PG_VERSION/main/postgresql.conf
RUN echo "fsync = 'off'" >> /etc/postgresql/$PG_VERSION/main/postgresql.conf
RUN echo "max_wal_size = 1GB" >> /etc/postgresql/$PG_VERSION/main/postgresql.conf
RUN echo "shared_buffers = 128MB" >> /etc/postgresql/$PG_VERSION/main/postgresql.conf
RUN echo "effective_cache_size = 512MB" >> /etc/postgresql/$PG_VERSION/main/postgresql.conf
RUN echo "work_mem = 64MB" >> /etc/postgresql/$PG_VERSION/main/postgresql.conf
RUN echo "autovacuum = off" >> /etc/postgresql/$PG_VERSION/main/postgresql.conf

USER root

RUN wget "https://dl.google.com/go/go1.10.linux-amd64.tar.gz"
RUN tar -C /usr/local -xzf go1.10.linux-amd64.tar.gz &&\
mkdir go && mkdir go/src && mkdir go/bin && mkdir go/pkg

ENV GOPATH $HOME/go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH &&\
export PATH=$PATH:/usr/local/go/bin

ADD ./ $GOPATH/src/go_tp_db/
EXPOSE 5000
WORKDIR $GOPATH/src/go_tp_db

CMD service postgresql start && go run main.go
