FROM rcampbell/opencv

ENV GOPATH /go

RUN go get -u -d gocv.io/x/gocv
WORKDIR $GOPATH/src/gocv.io/x/gocv
RUN make download
RUN make build
RUN make clean

RUN go get github.com/hybridgroup/mjpeg

# Define default command
CMD ["bash"]
