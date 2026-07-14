.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./post/main
	rm -rf ./get/main
	rm -rf ./delete/main
	
build:
	GOOS=linux GOARCH=amd64 go build -o post/main ./post
	GOOS=linux GOARCH=amd64 go build -o get/main ./get
	GOOS=linux GOARCH=amd64 go build -o delete/main ./delete