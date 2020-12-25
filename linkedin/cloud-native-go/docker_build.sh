#!/bin/sh 

CGO_ENABLED=0 GOOS=linux go build -o cloud-native-go
docker build -t cloud-native-go:1.0.2 .

docker tag cloud-native-go:1.0.2 rcampbell/cloud-native-go:1.0.2
docker push rcampbell/cloud-native-go:1.0.2
