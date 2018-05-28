# Pull base image
#FROM resin/rpi-raspbian:jessie
FROM resin/raspberry-pi-golang:1.9

ADD ./etc/services /etc/services

# Set environment variables
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH

# Define working directory
WORKDIR /go

RUN apt-get update 

# Install dependencies
RUN apt-get update && apt-get install -y \
    ca-certificates \
    gcc \
    libc6-dev \
    make \
    git \
    --no-install-recommends && \
    rm -rf /var/lib/apt/lists/*


RUN go get -d -u gobot.io/x/gobot/...
# RUN apt-get update 
# RUN apt-get install -y libopencv-dev
# RUN go get -u -d gocv.io/x/gocv
# RUN apt-get install -y sudo
# RUN sudo adduser root sudo
# WORKDIR $GOPATH/src/gocv.io/x/gocv
# RUN yes | make deps
# RUN make download
# RUN make build
# RUN make clean

# Define default command
CMD ["bash"]
