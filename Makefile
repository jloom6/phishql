.PHONY: bootstrap
bootstrap:
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/twitchtv/retool
	retool add github.com/golang/mock/mockgen origin/master

.PHONY: bootstrap-db
bootstrap-db:
	docker exec -it phishql-mysql mysql -u root -pkingofprussia -e "GRANT ALL PRIVILEGES ON *.* TO 'wilson'@'%'"
	docker exec -it phishql-mysql mysql -u root -pkingofprussia -e "DROP DATABASE IF EXISTS phish"
	docker exec -it phishql-mysql mysql -u root -pkingofprussia -e "CREATE DATABASE phish"
	cat fixtures/init.sql | docker exec -i phishql-mysql mysql -u root -pkingofprussia phish

.PHONY: proto
proto:
	mkdir -p .gen/proto
	protoc -I/usr/local/include -I. \
	 -I$(GOPATH)/src \
	 -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	 --swagger_out=logtostderr=true:. \
	 --grpc-gateway_out=logtostderr=true:.gen \
	 --go_out=plugins=grpc:.gen \
	 proto/jloom6/phishql/phishql.proto

.PHONY: build
build:
	dep ensure
	make proto
	GOOS=linux go build -o ./cmd/api/phishql-api ./cmd/api/main.go
	GOOS=linux go build -o ./cmd/proxy/phishql-proxy ./cmd/proxy/main.go
	docker build -f cmd/api/Dockerfile -t jloom6/phishql-api .
	docker build -f cmd/proxy/Dockerfile -t jloom6/phishql-proxy .
	docker build -f cmd/migration/Dockerfile -t jloom6/phishql-migration .

.PHONY: mocks
mocks:
	go generate ./...

.PHONY: test
test:
	make mocks
	go test ./... -coverprofile cover.out; go tool cover -func cover.out
	@rm cover.out

.PHONY: run-api
run-api:
	docker run -p 9090:9090 --name=phishql-api -e "PHISHQL_MYSQL_HOST=$$(docker-machine ip default)" jloom6/phishql-api

.PHONY: run-db
run-db:
	docker run -p 3306:3306 --name phishql-mysql -e MYSQL_ROOT_PASSWORD=kingofprussia -e MYSQL_USER=wilson -e MYSQL_PASSWORD=wilson -e MYSQL_DATABASE=phish mysql

.PHONY: run-proxy
run-proxy:
	docker run -p 8080:8080 --name=phishql-proxy -e "PHISHQL_API_ENDPOINT=$$(docker-machine ip default):9090" jloom6/phishql-proxy

.PHONY: run-all
run-all:
	docker-compose up

.PHONY: clean
clean:
	-docker stop phishql-proxy phishql-api phishql-migration phishql-mysql
	-docker rm phishql-proxy phishql-api phishql-migration phishql-mysql
	-docker image rm jloom6/phishql-migration jloom6/phishql-proxy jloom6/phishql-api
	rm -rf .gen vendor

.PHONY: run-hard
run-hard:
	make clean
	make build
	make run-all

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	make fmt
	golint ./... | grep -v vendor/ || :
