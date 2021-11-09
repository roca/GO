FROM golang:1.17

# RUN apk update && apk upgrade && apk add --no-cache bash git

RUN go install github.com/micro/micro@latest && \
    go install github.com/micro/protobuf/proto@latest && \
    go install github.com/micro/protobuf/protoc-gen-go@latest

ENV SOURCES /app/Go-Micro/
RUN mkdir -p ${SOURCES}
COPY . ${SOURCES}

RUN cd ${SOURCES}server/ && CGO_ENABLED=0 go build

ENV CONSUL_HTTP_ADDR localhost:8500

WORKDIR ${SOURCES}server/
# CMD ["tail", "-f", "/dev/null"]
CMD ${SOURCES}server/server
