FROM golang:1.11-alpine

RUN apk add --no-cache ca-certificates git

RUN go get github.com/codegangsta/gin

RUN go get github.com/subosito/gotenv

RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/gorilla/mux
RUN go get github.com/lib/pq
RUN go get golang.org/x/crypto/bcrypt
RUN go get github.com/davecgh/go-spew/spew

RUN mkdir -p /go/src/github.com/GOCODE/udemy/RESTfulJWT/api
WORKDIR /go/src/github.com/GOCODE/udemy/RESTfulJWT/api

COPY . $WORKDIR

EXPOSE 3001

# CMD ["tail", "-f", "/dev/null"]
ENTRYPOINT  gin  --bin "rtsupport" --immediate --appPort 3001 run main.go