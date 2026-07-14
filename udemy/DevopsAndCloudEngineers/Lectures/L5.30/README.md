
# Start up the container
docker run --rm -it -v $PWD:/opt/app alpine


## Inside container
```
apk add go

cd /opt/app
go build -o main main.go 
ldd main
CGO_ENABLED=0 go build -o main-nocgo main.go 
ldd main-nocgo 
GOOS=darwin ARCH=amd64  go build -o main-darwim64 main.go 
ldd main-darwim64 
```