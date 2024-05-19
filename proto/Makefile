LOCAL_BIN:=$(CURDIR)/generated/bin

# example for windows WSL: make -f Makefile install-deps

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
	make generate-olo

generate-olo:
	mkdir -p ../olo_service/generated
	protoc --proto_path $(CURDIR)/olo --proto_path $(CURDIR)/generated/vendor.protogen \
	--go_out=$(CURDIR)/../olo_service/generated --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=$(CURDIR)/generated/bin/protoc-gen-go \
	--go-grpc_out=$(CURDIR)/../olo_service/generated --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=$(CURDIR)/generated/bin/protoc-gen-go-grpc \
	--grpc-gateway_out=$(CURDIR)/../olo_service/generated --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true \
	--plugin=protoc-gen-grpc-gateway=$(CURDIR)/generated/bin/protoc-gen-grpc-gateway \
	$(CURDIR)/olo/olo.proto

generated-migration_for_service_olo:
	migrate -path $(CURDIR)/../olo_service/schema -database 'mysql://root:password@tcp(localhost:5055)/sso?x-tls-insecure-skip-verify=true' up

generated-auth-api:
	mkdir -p ../auth_service/generated
	protoc --proto_path $(CURDIR)/auth --proto_path $(CURDIR)/generated/vendor.protogen \
	--go_out=$(CURDIR)/../auth_service/generated --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=$(CURDIR)/generated/bin/protoc-gen-go \
	--go-grpc_out=$(CURDIR)/../auth_service/generated --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=$(CURDIR)/generated/bin/protoc-gen-go-grpc \
 	--grpc-gateway_out=$(CURDIR)/../auth_service/generated --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true \
 	--plugin=protoc-gen-grpc-gateway=$(CURDIR)/generated/bin/protoc-gen-grpc-gateway \
	$(CURDIR)/auth/auth.proto

#http://localhost:6060/pkg/?m=all - docs for all packages
docs:
	godoc -http=:6060

vendor-proto:
		@if [ ! -d generated/vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis generated/vendor.protogen/googleapis && \
			mkdir -p generated/vendor.protogen/google/ && \
			mv generated/vendor.protogen/googleapis/google/api generated/vendor.protogen/google && \
			rm -rf generated/vendor.protogen/googleapis ;\
		fi

test-auth-service:
	docker compose -f ../docker-compose.yml up db -d --wait
	go test -v ../auth_service/tests/...
	docker compose down