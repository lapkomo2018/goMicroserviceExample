.PHONY: docs proto clean

# Command to generate documentation
docs:
	swag init --parseDependency --parseInternal -g cmd/app/main.go

PROTO_DIR := .\api\grpc
PROTO_OUT_DIR := .\pkg\grpc\example

clean:
	@if exist $(PROTO_OUT_DIR) (rd /s /q $(PROTO_OUT_DIR))
	@mkdir $(PROTO_OUT_DIR)

proto: clean
	protoc --go_out=$(PROTO_OUT_DIR) --go-grpc_out=$(PROTO_OUT_DIR) --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative $(PROTO_DIR)\example.proto
	@move /Y $(PROTO_OUT_DIR)\api\grpc\example.pb.go $(PROTO_OUT_DIR)\service.pb.go
	@move /Y $(PROTO_OUT_DIR)\api\grpc\example_grpc.pb.go $(PROTO_OUT_DIR)\service_grpc.pb.go
	@rd /s /q $(PROTO_OUT_DIR)\api
