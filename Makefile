proto-compile: proto-compile-service proto-compile-model

proto-compile-service:
	 protoc --proto_path=shared/proto --go_out=paths=source_relative:./shared/service --go-grpc_out=paths=source_relative:./shared/service shared/proto/*.service.proto

proto-compile-model:
	protoc --proto_path=shared/proto --go_out=paths=source_relative:./shared/model shared/proto/*.model.proto
