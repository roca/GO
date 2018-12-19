FROM golang:latest

RUN mkdir -p /go/src/github.com/GOCODE/udemy/ByExample
WORKDIR /go/src/github.com/GOCODE/udemy/ByExample

COPY . /go/src/github.com/GOCODE/udemy/ByExample

RUN go get github.com/codegangsta/gin

EXPOSE 3000

CMD ["tail","-f","/dev/null"]