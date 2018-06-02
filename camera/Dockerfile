FROM rcampbell/gocv

RUN mkdir -p /var/app

COPY camera.go /var/app
WORKDIR /var/app

# RUN source /go/src/gocv.io/x/gocv/env.sh
#RUN GOARM=7 GOARCH=arm GOOS=linux go build camera.go
RUN go build camera.go
ENV PORT 3000
EXPOSE 3000
CMD ["./camera"]