protoc:
	protoc --go_out=./app/domain/pb --go_opt=paths=source_relative \
    		--go-grpc_out=./app/domain/pb --go-grpc_opt=paths=source_relative \
            ./proto/*/*.proto

build-rest:
	go build ./cmd/rest

build-grpc:
	go build ./cmd/grpc

start-rest:
	./rest

start-grpc:
	./grpc

conf:
	cp config-example.yaml cmd/rest/config.yaml && cp config-example.yaml cmd/grpc/config.yaml