.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o build/add-note 		api/add-note/main.go
	env GOOS=linux go build -ldflags="-s -w" -o build/delete-note 	api/delete-note/main.go
	env GOOS=linux go build -ldflags="-s -w" -o build/get-note 		api/get-note/main.go
	env GOOS=linux go build -ldflags="-s -w" -o build/update-note 	api/update-note/main.go

clean:
	rm -rf ./build ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
