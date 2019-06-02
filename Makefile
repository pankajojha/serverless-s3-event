.PHONY: build clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/events events/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/postEvent upload/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/handlers/addEvent handlers/addEvent.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/handlers/listEvents handlers/listEvents.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
