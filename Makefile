.PHONY: install test

install:
	go install cmd/protoc-gen-gosvc.go

test: install
	protoc --go_out=. --go_opt=Mproto/helloworld.proto=./proto ./proto/*.proto
	protoc --gosvc_out=. --gosvc_opt=Mproto/helloworld.proto=./proto ./proto/*.proto

