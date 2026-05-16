MODULE=github.com/I-Frostbyte/Learning-Curve/gRPC/grpc-in-go
PROTO_DIR=protos

run-server:
	go run server/main.go

run-client:
	go run client/main.go

generate-protos:
	@for file in $(PROTO_DIR)/*.proto; do \
		protoc \
			--go_out=. \
			--go_opt=module=$(MODULE) \
			--go-grpc_out=. \
			--go-grpc_opt=module=$(MODULE) \
			$$file; \
	done


# old-generate-protos:
# 	protoc \
# 		--go_out=. \
# 		--go_opt=module=github.com/I-Frostbyte/Learning-Curve/gRPC/grpc-in-go \
# 		--go-grpc_out=. \
# 		--go-grpc_opt=module=github.com/I-Frostbyte/Learning-Curve/gRPC/grpc-in-go \
# 		protos/service.proto