#!/bin/sh 

# CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o cloud-native-go
# docker build -t cloud-native-go:1.0.2 .

docker buildx use rpi3-ssh
docker buildx build -t rcampbell/cloud-native-go:1.0.2 --platform linux/arm/v7  .

# docker tag cloud-native-go:1.0.2 rcampbell/cloud-native-go:1.0.2
docker push rcampbell/cloud-native-go:1.0.2
