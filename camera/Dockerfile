FROM rcampbell/gocv
ENV GOPATH /go

# RUN source /go/src/gocv.io/x/gocv/env.sh
RUN mkdir -p /go/src/camera
WORKDIR /go/src/camera
COPY camera.go $GOPATH/src/camera
#RUN GOARM=7 GOARCH=arm GOOS=linux go build camera.go
RUN go build camera.go
ENV PORT 3000
EXPOSE 3000
CMD ["./camera"]