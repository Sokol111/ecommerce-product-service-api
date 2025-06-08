OAPI_GEN := $(HOME)/go/bin/oapi-codegen
OPENAPI_FILE := openapi/openapi.yml
GEN_DIR := gen

.PHONY: all client server types clean install-tools

all: install-tools clean types server client

install-tools:
	@which oapi-codegen >/dev/null || go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

types:
	@echo "Generating types (models)..."
	mkdir -p $(GEN_DIR)
	$(OAPI_GEN) -generate types -package gen -o $(GEN_DIR)/models.gen.go $(OPENAPI_FILE)

server:
	@echo "Generating Go server..."
	mkdir -p $(GEN_DIR)
	$(OAPI_GEN) -generate gin-server,strict-server -package gen -o $(GEN_DIR)/server.gen.go $(OPENAPI_FILE)

client:
	@echo "Generating Go client..."
	mkdir -p $(GEN_DIR)
	$(OAPI_GEN) -generate client -package gen -o $(GEN_DIR)/client.gen.go $(OPENAPI_FILE)

clean:
	rm -rf $(GEN_DIR)