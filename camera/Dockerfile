FROM rcampbell/gocv

RUN mkdir -p /var/app
WORKDIR /var/app

COPY camera.go /var/app
RUN GOARM=7 GOARCH=arm GOOS=linux go build camera.go
ENV PORT 3000
EXPOSE 3000
CMD ["./camera"]