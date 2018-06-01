FROM rcampbell/rpi-golang

RUN apt-get update 
RUN apt-get install -y libopencv-dev
RUN go get -u -d gocv.io/x/gocv
WORKDIR $GOPATH/src/gocv.io/x/gocv
RUN apt-get install -y make
RUN sudo echo "deb https://packages.ubuntu.com trusty universe" >> /etc/apt/sources.list
RUN sudo apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    software-properties-common
RUN apt-get update 
RUN yes | make deps
RUN make download
RUN make build
RUN make clean

# Define default command
CMD ["bash"]
