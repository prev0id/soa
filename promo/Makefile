proto-loyalty:
	protoc \
	  --proto_path=api \
	  --go_out=pkg/loyalty --go_opt=paths=source_relative \
	  --go-grpc_out=pkg/loyalty --go-grpc_opt=paths=source_relative \
	  api/loyalty.proto

proto-events:
	protoc \
	  --proto_path=api \
	  --go_out=pkg/events --go_opt=paths=source_relative \
	  --go-grpc_out=pkg/events --go-grpc_opt=paths=source_relative \
	  api/events.proto


proto-loyalty-client:
	protoc \
	  --proto_path=api \
	  --go_out=../api/pkg/loyalty --go_opt=paths=source_relative \
	  --go-grpc_out=../api/pkg/loyalty --go-grpc_opt=paths=source_relative \
	  api/loyalty.proto
