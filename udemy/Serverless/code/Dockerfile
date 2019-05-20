FROM golang:1.12

RUN go get github.com/codegangsta/gin
RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get github.com/ktr0731/evans
RUN go get go.mongodb.org/mongo-driver/mongo


RUN mkdir -p /code
WORKDIR /code

COPY . $WORKDIR

RUN apt-get update && apt-get install -y unzip 
RUN export PROTOC_ZIP=protoc-3.7.1-linux-x86_64.zip \
&& curl -OL https://github.com/google/protobuf/releases/download/v3.7.1/$PROTOC_ZIP \
&& unzip -o $PROTOC_ZIP -d /usr/local bin/protoc \
&& rm -f $PROTOC_ZIP


EXPOSE 50051

CMD ["tail", "-f", "/dev/null"]
# ENTRYPOINT  gin  --bin "api" --immediate --appPort 3001 run main.go