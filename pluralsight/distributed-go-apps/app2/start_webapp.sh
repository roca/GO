#!/bin/sh -f
cd src/distributed/web
go build main.go
./main
cd ../../..
