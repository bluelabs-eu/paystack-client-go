# Variables
GENERATOR := openapi-generator 
OUTPUT_DIR := .
OPENAPI_FILE := api/openapi.yaml
YQ := ./yq  # Use the local yq binary

# Install dependencies
install-deps:
	@if [ ! -x $(YQ) ]; then \
		echo "Installing yq..."; \
		curl -L https://github.com/mikefarah/yq/releases/latest/download/yq_darwin_amd64 -o $(YQ); \
		chmod +x $(YQ); \
		echo "yq installed locally as $(YQ)."; \
	else \
		echo "yq is already installed."; \
	fi

# Install OpenAPI Generator CLI
install-generator:
	@if ! command -v $(GENERATOR) &> /dev/null; then \
		echo "Installing openapi-generator-cli via Homebrew..."; \
		brew install openapi-generator; \
	else \
		echo "openapi-generator-cli is already installed."; \
	fi

# Show modifiable paths
show-paths:
	@if [ ! -x $(YQ) ]; then \
		echo "Error: yq is not installed. Please run 'make install-deps'."; \
		exit 1; \
	fi
	@echo "Available paths in the OpenAPI file:"
	@$(YQ) eval '.paths | keys' $(OPENAPI_FILE)

# Modify OpenAPI schema
# Usage: make modify PATH=<path> SCHEMA=<schema_type>
modify: install-deps
	@if [ -z "$(PATH)" ] || [ -z "$(SCHEMA)" ]; then \
		echo "Usage: make modify PATH=<path> SCHEMA=<schema_type>"; \
		exit 1; \
	fi
	chmod +x ./modify_openapi_schema.sh
	YQ=$(YQ) ./modify_openapi_schema.sh $(OPENAPI_FILE) $(PATH) $(SCHEMA)


# Regenerate code
regenerate-code: install-generator
	@echo "Regenerating code from OpenAPI spec..."
	$(GENERATOR) generate -i $(OPENAPI_FILE) -g go -o $(OUTPUT_DIR)
	@echo "Code regeneration complete."