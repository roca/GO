.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o build/invokeImageProcessor invokeImageProcessor/main.go
	env GOOS=linux go build -ldflags="-s -w" -o build/getFileType getFileType/main.go
clean:
	rm -rf ./build ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
