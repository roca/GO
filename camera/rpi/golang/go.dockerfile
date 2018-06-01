# Pull base image
#FROM resin/rpi-raspbian:jessie
FROM resin/raspberry-pi-golang:1.9-slim

ADD ./etc/services /etc/services

# Set environment variables
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH

# Define working directory
WORKDIR /go

# Define default command
CMD ["bash"]
