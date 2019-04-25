FROM golang:1.12

RUN go get github.com/codegangsta/gin
RUN go get google.golang.org/grpc
RUN go get github.com/golang/protobuf/protoc-gen-go

RUN mkdir -p /code
WORKDIR /code

COPY . $WORKDIR

EXPOSE 3001

CMD ["tail", "-f", "/dev/null"]
# ENTRYPOINT  gin  --bin "api" --immediate --appPort 3001 run main.go