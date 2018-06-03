FROM rcampbell/rpi-golang

RUN apt-get update 
RUN echo 'APT::Get::Assume-Yes "true";' >> /etc/apt/apt.conf
RUN echo 'APT::Get::force-yes "true";' >> /etc/apt/apt.conf
RUN sudo apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    software-properties-common \
    make
RUN sudo echo "deb http://deb.debian.org/debian  jessie main" >> /etc/apt/sources.list
RUN apt-get update 

RUN apt-get install libopencv-dev
RUN go get -u -d gocv.io/x/gocv
WORKDIR $GOPATH/src/gocv.io/x/gocv
RUN yes | make deps
RUN make download
RUN make build
RUN make clean

RUN go get github.com/hybridgroup/mjpeg

# Define default command
CMD ["bash"]
