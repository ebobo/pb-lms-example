all: lint gen build

lint:
	@buf lint

gen: clean
	@buf generate 

clean:
	@rm -rf pkg/api

build:
	@go mod tidy
	@cd cmd && go build -o ../bin/lms

client:
	@cd cmd/client && go build -o ../../bin/client 

init:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@go install github.com/bufbuild/buf/cmd/buf@latest
