.PHONY: build
build:
	go get ./...
	GOOS=linux go build -ldflags "-s -w" ./
	goupx siteback

.PHONY: test
test:
	GO111MODULE=on go test -v -race -coverprofile=cover.out -covermode=atomic ./...
	go tool cover -html=cover.out -o cover.html
