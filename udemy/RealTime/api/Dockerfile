FROM golang:latest

RUN mkdir -p /go/src/github.com/GOCODE/udemy/RealTime
WORKDIR /go/src/github.com/GOCODE/udemy/RealTime

COPY . /go/src/github.com/GOCODE/udemy/RealTime

EXPOSE 3001

CMD ["tail","-f","/dev/null"]