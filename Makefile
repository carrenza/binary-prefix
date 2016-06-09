
build:
	go build -v ./...

install:
	go install -v

lint:
	golint ./...
	go vet ./...

test:
	go test -v ./...

cover:
	go test -v ./... --cover

deps: dev-deps

dev-deps:
	go get -u github.com/golang/lint/golint
	go get -u github.com/smartystreets/goconvey/convey
