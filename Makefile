.PHONY: install test

install:
	go install cmd/protoc-gen-gosvc/protoc-gen-gosvc.go

example: install
	protoc --go_out=. --go_opt=Mproto/helloworld.proto=./proto ./proto/*.proto
	protoc --gosvc_out=. --gosvc_opt=Mproto/helloworld.proto=./proto ./proto/*.proto
	go run cmd/example/example.go

