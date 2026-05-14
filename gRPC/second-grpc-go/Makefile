run-server:
	go run server/main.go

run-client:
	go run client/main.go

generate-protos:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		protos/helloworld.proto