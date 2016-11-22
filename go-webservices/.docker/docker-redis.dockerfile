FROM 		redis:latest

MAINTAINER 	Romel Campbell

#Change as appropriate for build
ENV APP_ENV development

COPY ./.docker/config/redis.${APP_ENV}.conf /etc/redis.conf

EXPOSE      6379

ENTRYPOINT  ["redis-server", "/etc/redis.conf"]

# To build:
# docker build -f docker-redis.dockerfile --tag rcampbell/redis ../

# To run:
# docker run -d -p 6379:6379 --name redis rcampbell/redis