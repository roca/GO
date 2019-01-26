FROM postgres:latest

MAINTAINER Romel Campbell

ENV SCRIPTS_HOME /var/db/scripts
RUN mkdir -p $SCRIPTS_HOME
WORKDIR $SCRIPTS_HOME

RUN chmod -R 777 /var/lib/postgresql/data

ADD ./.docker/pgsetup/test-pg-data.sql $SCRIPTS_HOME/

EXPOSE 5432