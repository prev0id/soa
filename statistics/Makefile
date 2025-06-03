proto:
	protoc \
	  --proto_path=api \
	  --go_out=pkg/stats --go_opt=paths=source_relative \
	  --go-grpc_out=pkg/stats --go-grpc_opt=paths=source_relative \
	  api/stats.proto

proto-client:
	protoc \
	  --proto_path=api \
	  --go_out=../api/pkg/stats --go_opt=paths=source_relative \
	  --go-grpc_out=../api/pkg/stats --go-grpc_opt=paths=source_relative \
	  api/stats.proto
