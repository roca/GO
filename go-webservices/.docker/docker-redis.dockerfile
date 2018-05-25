FROM 		redis:latest

MAINTAINER 	Romel Campbell

COPY ./db/redis/config/redis.conf /etc/redis.conf

EXPOSE      6379

ENTRYPOINT  ["redis-server", "/etc/redis.conf"]
