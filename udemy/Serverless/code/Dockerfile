FROM golang:1.12

RUN go get github.com/codegangsta/gin
# RUN go get github.com/aws/aws-lambda-go
# RUN go get github.com/aws/aws-sdk-go

RUN mkdir -p /code
WORKDIR /code

COPY . $WORKDIR

RUN apt-get update && apt-get install -y unzip zip

ENV GOOS=linux
ENV GOARCH=amd64

#CMD ["tail", "-f", "/dev/null"]
ENTRYPOINT  gin  --bin "main"
# ENTRYPOINT  gin  --bin "api" --immediate --appPort 3001 run main.go