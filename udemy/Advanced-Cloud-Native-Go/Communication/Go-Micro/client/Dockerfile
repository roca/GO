FROM golang:1.15

# RUN apk update && apk upgrade && apk add --no-cache bash git

# RUN go get -u github.com/micro/micro && \
#     go get github.com/micro/protobuf/proto && \
#     go get -u github.com/micro/protobuf/protoc-gen-go && \
#     go get github.com/micro/go-plugins/wrapper/breaker/hystrix && \
#     go get github.com/afex/hystrix-go/hystrix

ENV SOURCES /app/Go-Micro/
RUN mkdir -p ${SOURCES}
COPY . ${SOURCES}

RUN cd ${SOURCES}client/ && CGO_ENABLED=0 go build

ENV CONSUL_HTTP_ADDR localhost:8500

WORKDIR ${SOURCES}client/
# CMD ["tail", "-f", "/dev/null"]
CMD ${SOURCES}client/client
