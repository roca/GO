FROM golang:latest

RUN mkdir -p /go/src/github.com/GOCODE/udemy/MachineLearning
WORKDIR /go/src/github.com/GOCODE/udemy/MachineLearning

COPY . /go/src/github.com/GOCODE/udemy/MachineLearning

RUN go get github.com/codegangsta/gin
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/lib/pq

RUN go get github.com/kniren/gota/dataframe
RUN go get github.com/patrickmn/go-cache
RUN go get github.com/boltdb/bolt
RUN go get github.com/gonum/matrix/mat64
RUN go get github.com/gonum/floats
RUN go gtet github.com/gonum/blas/blas64

EXPOSE 3000

CMD ["tail","-f","/dev/null"]