
GOPATH:=$(shell go env GOPATH)

.PHONY: proto test docker


proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/hello/hello.proto

build: proto
	mkdir -p ./bin
	go build -o ./bin/microsample-srv main.go plugin.go
	go build -o ./bin/client ./client/main.go

test:
	go test -v ./... -cover

docker:
	docker build . -t microsample-srv:latest
