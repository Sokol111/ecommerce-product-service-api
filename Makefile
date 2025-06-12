.PHONY: update-makefiles gen-go gen-js

OPENAPI_FILE ?= openapi/openapi.yml
VERSION ?= v1.0.0
GO_GEN_DIR ?= api
PACKAGE ?= api
JS_GEN_DIR ?= js-client
NPM_PACKAGE_NAME ?= @sokol111/ecommerce-product-service-api
PROJECT_NAME ?= ecommerce-product-service-api
AUTHOR ?= Sokol111
REPOSITORY_URL ?= https://github.com/Sokol111/ecommerce-product-service-api.git

update-makefiles:
	@echo "Updating includes in Makefile..."
	curl -sSL https://raw.githubusercontent.com/Sokol111/ecommerce-infrastructure/master/makefiles/generate-go-api.mk -o generate-go-api.mk
	curl -sSL https://raw.githubusercontent.com/Sokol111/ecommerce-infrastructure/master/makefiles/generate-js-api.mk -o generate-js-api.mk
	@echo "Done."

gen-go: update-makefiles
	@echo "Generating Go API..."
	make -f generate-go-api.mk generate-go-api \
		OPENAPI_FILE=$(OPENAPI_FILE) \
		GO_GEN_DIR=$(GO_GEN_DIR) \
		PACKAGE=$(PACKAGE)
	@echo "Go API generated successfully."

gen-js: update-makefiles
	@echo "Generating JS API..."
	make -f generate-js-api.mk generate-js-api \
		OPENAPI_FILE=$(OPENAPI_FILE) \
		JS_GEN_DIR=$(JS_GEN_DIR) \
		NPM_PACKAGE_NAME=$(NPM_PACKAGE_NAME) \
		VERSION=$(VERSION) \
		PROJECT_NAME=$(PROJECT_NAME) \
		AUTHOR=$(AUTHOR) \
		REPOSITORY_URL=$(REPOSITORY_URL)
	@echo "JS API generated successfully."
