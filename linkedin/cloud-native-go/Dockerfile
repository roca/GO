FROM golang:1.13.0-alpine

ENV GO111MODULE=on
ENV HOME /linkedin/cloud-native-go/api
ENV PORT 9090

WORKDIR ${HOME}

COPY . ${HOME}

RUN CGO_ENABLED=0 GOOS=linux go build -o cloud-native-go

EXPOSE ${PORT}

ENTRYPOINT ["./cloud-native-go"]
