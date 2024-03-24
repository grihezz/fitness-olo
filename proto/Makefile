LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.33.0
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.19.1

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway

generate:
	make generated-auth-api

generated-auth-api:
	mkdir -p ../auth_service/generated
	protoc --proto_path $(CURDIR)/auth --proto_path $(CURDIR)/vendor.protogen \
	--go_out=$(CURDIR)/../auth_service/generated --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=$(CURDIR)/bin/protoc-gen-go \
	--go-grpc_out=$(CURDIR)/../auth_service/generated --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=$(CURDIR)/bin/protoc-gen-go-grpc \
 	--grpc-gateway_out=$(CURDIR)/../auth_service/generated --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true \
 	--plugin=protoc-gen-grpc-gateway=$(CURDIR)/bin/protoc-gen-grpc-gateway \
	$(CURDIR)/auth/auth.proto

vendor-proto:
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis && \
			mkdir -p vendor.protogen/google/ && \
			mv vendor.protogen/googleapis/google/api vendor.protogen/google && \
			rm -rf vendor.protogen/googleapis ;\
		fi