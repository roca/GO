FROM golang:alpine3.14

ENV SOURCES /app/K8s/
RUN mkdir -p ${SOURCES}
COPY . ${SOURCES}

RUN cd ${SOURCES}server/ && CGO_ENABLED=0 go build -o server

WORKDIR ${SOURCES}server/
CMD ${SOURCES}server/server
