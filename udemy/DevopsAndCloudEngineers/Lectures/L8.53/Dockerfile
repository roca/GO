#
# Build go project
#
FROM golang:1.18-alpine as go-builder

WORKDIR /github-deploy

COPY . .

RUN go mod init rcampbell/github-deploy && \
go mod tidy

# Tooling and rerquired packages
RUN apk add -u -t build-tools curl git  && \
go get github.com/google/go-github/v47/github && \
go get k8s.io/api/apps/v1 && \
go get rcampbell/github-deploy && \
go get k8s.io/apimachinery/pkg/api/errors && \
go get k8s.io/apimachinery/pkg/apis/meta/v1 && \
go get k8s.io/apimachinery/pkg/labels && \
go get k8s.io/client-go/kubernetes && \
go get k8s.io/client-go/kubernetes/scheme && \
go get k8s.io/client-go/rest && \
go get k8s.io/client-go/tools/clientcmd && \
go get k8s.io/client-go/util/homedir

# Build the binary and clean up
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o github-deploy . && \
apk del build-tools && \
rm -rf /var/cache/apk/*

#
# Runtime container
#
FROM alpine:latest  

WORKDIR /app

RUN apk --no-cache add ca-certificates bash curl

COPY --from=go-builder /github-deploy/github-deploy /github-deploy

ENTRYPOINT ["/github-deploy"]