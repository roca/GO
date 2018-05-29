FROM rcampbell/rpi-golang

RUN apt-get update 
RUN apt-get install -y libopencv-dev
RUN go get -u -d gocv.io/x/gocv
WORKDIR $GOPATH/src/gocv.io/x/gocv
RUN yes | make deps
RUN make download
RUN make build
RUN make clean

# Define default command
CMD ["bash"]
