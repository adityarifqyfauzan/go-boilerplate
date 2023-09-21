proto:
	protoc --go_out=./app/domain/pb --go_opt=paths=source_relative \
    		--go-grpc_out=./app/domain/pb --go-grpc_opt=paths=source_relative \
            ./proto/*/*.proto