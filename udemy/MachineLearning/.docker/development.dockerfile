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
RUN go get github.com/gonum/blas/blas64

RUN go get github.com/gonum/stat
RUN go get github.com/kniren/gota/dataframe
RUN go get github.com/montanaflynn/stats

RUN go get gonum.org/v1/plot
RUN go get github.com/gonum/mathext
RUN go get github.com/sajari/regression
RUN go get github.com/berkmancenter/ridge
RUN go get github.com/sjwhitworth/golearn/base
RUN go get github.com/sjwhitworth/golearn/evaluation
RUN go get github.com/sjwhitworth/golearn/knn

EXPOSE 3000

CMD ["tail","-f","/dev/null"]