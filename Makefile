BUILD_FILENAME = ova-travel-api
LOCAL_BIN:=$(CURDIR)/bin

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	go get ./...
	go build -o $(BUILD_FILENAME) ./cmd/ova-travel-api/main.go

.PHONY: generate_proto
generate_proto:
	mkdir -p pkg/ova-travel-api
	protoc \
			--go_out=pkg/ova-travel-api --go_opt=paths=import \
			--go-grpc_out=pkg/ova-travel-api --go-grpc_opt=paths=import \
			api/ova-travel-api.proto
	mv pkg/ova-travel-api/github.com/ozonva/ova-travel-api/pkg/ova-travel-api/* pkg/ova-travel-api/
	rm -rf pkg/ova-travel-api/github.com

.PHONY: deps
deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
	ls go.mod || go mod init gitlab.com/siriusfreak/lecture-6-demo
	GOBIN=$(LOCAL_BIN) go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	GOBIN=$(LOCAL_BIN) go get -u github.com/golang/protobuf/proto
	GOBIN=$(LOCAL_BIN) go get -u github.com/golang/protobuf/protoc-gen-go
	GOBIN=$(LOCAL_BIN) go get -u google.golang.org/grpc
	GOBIN=$(LOCAL_BIN) go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger