-include generate-api.mk

update-makefiles:
	@echo "Updating includes in Makefile..."
	curl -sSL https://raw.githubusercontent.com/Sokol111/ecommerce-makefiles/master/generate-api.mk -o generate-api.mk
	@echo "Done."
