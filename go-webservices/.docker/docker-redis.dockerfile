FROM 		redis:latest

MAINTAINER 	Romel Campbell

#Change as appropriate for build
ENV APP_ENV development

COPY ./.docker/config/redis.${APP_ENV}.conf /etc/redis.conf

EXPOSE      6379

ENTRYPOINT  ["redis-server", "/etc/redis.conf"]
