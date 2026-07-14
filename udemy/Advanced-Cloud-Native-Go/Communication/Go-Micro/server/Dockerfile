FROM golang:1.15

ENV SOURCES /Go-Micro/
RUN mkdir -p ${SOURCES}
COPY . ${SOURCES}
WORKDIR ${SOURCES}

RUN curl -fsSL https://raw.githubusercontent.com/micro/micro/master/scripts/install.sh | /bin/bash
RUN export PATH=/root/bin:$PATH

RUN go get google.golang.org/grpc

RUN go get -u github.com/golang/protobuf/proto
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get github.com/micro/micro/v3/cmd/protoc-gen-micro
RUN go get github.com/micro/micro/v3/cmd/protoc-gen-openapi

ENV CONSUL_HTTP_ADDR localhost:8500
EXPOSE 8080

RUN cd ${SOURCES}server && CGO_ENABLED=0 go build -o server main.go
CMD server/server --proxy_address=sky-net:8080 --auth_id admin --auth_secret micro
