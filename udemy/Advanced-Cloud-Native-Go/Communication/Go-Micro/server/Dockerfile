FROM golang:1.16.2

ENV SOURCES /Go-Micro/
RUN mkdir -p ${SOURCES}
COPY . ${SOURCES}
WORKDIR ${SOURCES}

RUN go mod init udemy.com

RUN go get github.com/golang/protobuf/proto
RUN go get github.com/micro/go-micro
RUN go get github.com/micro/go-micro/client
RUN go get github.com/micro/go-micro/server
RUN go get golang.org/x/net/context
RUN go get github.com/afex/hystrix-go/hystrix

RUN go mod download golang.org/x/sys
RUN go mod download golang.org/x/text

ENV CONSUL_HTTP_ADDR localhost:8500
RUN cd ${SOURCES}server && CGO_ENABLED=0 go build -o server main.go

CMD server/server
