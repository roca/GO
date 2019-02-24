FROM golang:1.11-alpine

RUN apk add --no-cache ca-certificates git

RUN go get github.com/codegangsta/gin
RUN go get  github.com/gorilla/websocket
RUN go get github.com/mitchellh/mapstructure

RUN mkdir -p /go/src/github.com/GOCODE/udemy/RealTime/api
WORKDIR /go/src/github.com/GOCODE/udemy/RealTime/api

COPY . $WORKDIR

EXPOSE 3001

# CMD ["tail", "-f", "/dev/null"]
ENTRYPOINT  gin  --bin "rtsupport" --immediate --appPort 3001 run main.go