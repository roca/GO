FROM golang:1.12.10-alpine

RUN apk update
RUN apk upgrade
RUN apk add bash

WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .



# Build the Go app
RUN go build -o rec /app/receiver/main.go

RUN go build -o sen /app/sender/main.go


CMD ["/bin/bash"]


