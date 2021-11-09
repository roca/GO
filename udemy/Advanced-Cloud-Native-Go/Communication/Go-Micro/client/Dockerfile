FROM golang:1.15

# RUN apk update && apk upgrade && apk add --no-cache bash git

# RUN go clean -modcache


# RUN go mod tidy

ENV SOURCES /app/Go-Micro/
RUN mkdir -p ${SOURCES}
COPY . ${SOURCES}

WORKDIR ${SOURCES}

RUN go mod init udemy.com

RUN go install github.com/micro/micro@latest && \
    go install github.com/golang/protobuf/proto@latest && \
    go install github.com/golang/protobuf/protoc-gen-go@latest && \
    go install github.com/micro/go-plugins/wrapper/breaker/hystrix@latest && \
    go install github.com/afex/hystrix-go/hystrix@latest


RUN CGO_ENABLED=0  go build -o client/client client/main.go

ENV CONSUL_HTTP_ADDR localhost:8500

# CMD ["tail", "-f", "/dev/null"]
CMD client/client
