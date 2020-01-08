.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o build/add-note 				api/add-note/main.go
	env GOOS=linux go build -ldflags="-s -w" -o build/delete-note 			api/delete-note/main.go
	env GOOS=linux go build -ldflags="-s -w" -o build/get-note 				api/get-note/main.go
	env GOOS=linux go build -ldflags="-s -w" -o build/update-note 			api/update-note/main.go
	env GOOS=linux go build -ldflags="-s -w" -o build/authenticate-google	api/authenticate-google/main.go
	env GOOS=linux go build -ldflags="-s -w" -o build/archive 				api/archive/main.go
	env GOOS=linux go build -ldflags="-s -w" -o build/kinesisProcessor 		api/kinesisProcessor/main.go
	env GOOS=linux go build -ldflags="-s -w" -o build/kinesisProducer 		api/kinesisProducer/main.go

clean:
	rm -rf ./build ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
