FROM golang:alpine3.14

RUN apk update && apk upgrade && apk add --no-cache bash git

RUN go get github.com/hashicorp/consul/api

ENV SOURCES /app/Simple/
RUN mkdir -p ${SOURCES}
COPY . ${SOURCES}

RUN cd ${SOURCES}client/ && CGO_ENABLED=0 go build

ENV CONSUL_HTTP_ADDR localhost:8500

WORKDIR ${SOURCES}client/
CMD ${SOURCES}client/client
