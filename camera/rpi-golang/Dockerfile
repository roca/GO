# Pull base image
#FROM resin/rpi-raspbian:jessie
FROM resin/raspberry-pi-golang:latest

ADD ./etc/services /etc/services

# Set environment variables
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH

# Define working directory
WORKDIR /go

RUN apt-get update
RUN apt-get upgrade
RUN apt-get install libopencv-dev

RUN go get -d -u gobot.io/x/gobot/...
RUN go get -u -d gocv.io/x/gocv
RUN cd $GOPATH/src/gocv.io/x/gocv
WORKDIR $GOPATH/src/gocv.io/x/gocv
RUN yes | make deps
RUN make download
RUN make build
RUN make clean

# Define default command
CMD ["bash"]
