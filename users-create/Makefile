.PHONY: build clean deploy

build: 
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/v1 v1/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: 
	sls deploy --verbose
