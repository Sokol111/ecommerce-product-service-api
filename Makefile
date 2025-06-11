-include generate-api.mk

.PHONY: update-makefiles

update-makefiles:
	@echo "Updating includes in Makefile..."
	curl -sSL https://raw.githubusercontent.com/Sokol111/ecommerce-makefiles/master/generate-api.mk -o generate-api.mk
	@echo "Done."

generate-go-api: update-makefiles
	@echo "Generating Go API..."
	make -f generate-api.mk install-tools
	make -f generate-api.mk types
	make -f generate-api.mk server
	make -f generate-api.mk client
	@echo "Go API generated successfully."

generate-js-api: update-makefiles
	@echo "Generating JS API..."
	make -f generate-api.mk install-tools
	make -f generate-api.mk js-generate
	make -f generate-api.mk js-package
	make -f generate-api.mk js-tsconfig
	make -f generate-api.mk js-build
	@echo "JS API generated successfully."