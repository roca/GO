FROM golang:latest

RUN mkdir -p /go/src/github.com/GOCODE/udemy/MachineLearning
WORKDIR /go/src/github.com/GOCODE/udemy/MachineLearning

COPY . /go/src/github.com/GOCODE/udemy/MachineLearning

RUN go get github.com/codegangsta/gin
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/lib/pq

RUN go get github.com/kniren/gota/dataframe

EXPOSE 3000

CMD ["tail","-f","/dev/null"]