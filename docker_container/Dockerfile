FROM golang:latest

RUN mkdir -p /go/src/container_demo/ROOT_FOR_CONTAINER/proc
WORKDIR /go/src/container_demo
# RUN mkdir -p /sys/module/user_namespace/parameters
# RUN mkdir -p /proc/sys/kernel

# RUN touch /proc/sys/kernel/unprivileged_userns_clone
# RUN echo Y > /sys/module/user_namespace/parameters/enabled


COPY ./main.go /go/src/container_demo

RUN go get github.com/codegangsta/gin

ENV GOOS linux
ENV GOARCH amd64

CMD ["cp", "-r", "/bin", "ROOT_FOR_CONTAINER"]
CMD gin run