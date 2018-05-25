FROM golang

RUN go get -d -u gobot.io/x/gobot/...
RUN go get -u -d gocv.io/x/gocv
RUN apt-get update 
RUN apt-get install -y sudo
RUN sudo adduser root sudo
WORKDIR $GOPATH/src/gocv.io/x/gocv
RUN yes | make deps
RUN make download
RUN make build
RUN make clean
