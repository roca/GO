.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o build/alice alice/main.go
	env GOOS=linux go build -ldflags="-s -w" -o build/bob bob/main.go
	env GOOS=linux go build -ldflags="-s -w" -o build/archive archive/main.go

clean:
	rm -rf ./build ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
