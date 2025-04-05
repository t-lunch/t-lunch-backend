PROTO_FILE=./pkg/api/proto/lunch.proto
OUT_API_DIR=./pkg/api/generated

all: protogen gobuild gorun

protogen:
	protoc --proto_path=./pkg/api/proto \
			--go_out $(OUT_API_DIR) --go_opt paths=source_relative \
            --go-grpc_out $(OUT_API_DIR) --go-grpc_opt paths=source_relative \
            --grpc-gateway_out $(OUT_API_DIR) --grpc-gateway_opt paths=source_relative \
            $(PROTO_FILE)

gobuild:
	go build -o ./bin/main.exe ./cmd/main/main.go

gorun:
	./bin/main.exe
