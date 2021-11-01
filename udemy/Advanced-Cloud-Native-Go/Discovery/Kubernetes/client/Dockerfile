FROM golang:alpine3.14

ENV SOURCES /app/K8s/
RUN mkdir -p ${SOURCES}
COPY . ${SOURCES}

RUN cd ${SOURCES}client/ && CGO_ENABLED=0 go build -o client

WORKDIR ${SOURCES}client/
CMD ${SOURCES}client/client
