OAPI_GEN := $(HOME)/go/bin/oapi-codegen
OPENAPI_FILE := openapi/openapi.yml

.PHONY: all client server clean install-tools

all: install-tools server client

install-tools:
	@which oapi-codegen || go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

server:
	@echo "Generating Go server..."
	rm -rf go-server
	mkdir -p go-server
	$(OAPI_GEN) -generate gin-server,strict-server -package server \
		-o go-server/server.gen.go $(OPENAPI_FILE)

client:
	@echo "Generating Go client..."
	rm -rf go-client
	mkdir -p go-client
	$(OAPI_GEN) -generate client -package client \
		-o go-client/client.gen.go $(OPENAPI_FILE)

clean:
	rm -rf go-client go-server