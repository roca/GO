FROM golang:1.12.10-alpine

RUN apk update
RUN apk upgrade
RUN apk add bash
RUN apk add git
RUN apk add gcc
RUN apk add curl

WORKDIR /app

COPY . .

EXPOSE 80

CMD ["tail","-f","/dev/null"]


