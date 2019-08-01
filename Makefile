.PHONY: bootstrap
bootstrap:
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/twitchtv/retool
	retool add github.com/golang/mock/mockgen origin/master

.PHONY: proto
proto:
	mkdir -p .gen/proto
	protoc -I/usr/local/include -I. \
	 -I$(GOPATH)/src \
	 -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	 --swagger_out=logtostderr=true:.gen \
	 --grpc-gateway_out=logtostderr=true:.gen \
	 --go_out=plugins=grpc:.gen \
	 proto/jloom6/phishql/phishql.proto

.PHONY: build
build:
	dep ensure
	go build -o ./cmd/api/phishql-api ./cmd/api/main.go
	go build -o ./cmd/proxy/phishql-proxy ./cmd/proxy/main.go

.PHONY: mocks
mocks:
	$(shell go generate `glide novendor`)

.PHONY: test
test:
	go test ./... -coverprofile cover.out; go tool cover -func cover.out
	rm cover.out

.PHONY: run-api
run-api:
	./cmd/api/phishql-api

.PHONY: run-proxy
run-proxy:
	./cmd/api/phishql-proxy
