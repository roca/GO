FROM golang:alpine3.14

RUN apk update && apk upgrade && apk add --no-cache bash git

RUN go get github.com/gin-gonic/gin

ENV SOURCES /go/src/github.com/Gin-Web/

RUN mkdir -p ${SOURCES}

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build -o web-app

WORKDIR ${SOURCES}
CMD ${SOURCES}web-app
EXPOSE 8080