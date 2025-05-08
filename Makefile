# Variables
GENERATOR_IMAGE := openapitools/openapi-generator-cli:v7.13.0
OUTPUT_DIR := .
OPENAPI_FILE := api/openapi.yaml
YQ_VERSION := 4.44.3
YQ_IMAGE := mikefarah/yq:$(YQ_VERSION)

# Docker options
DOCKER_RUN_OPTS := --rm -v ${PWD}:${PWD} -w ${PWD}

# Function to run yq via Docker
define yq
	docker run $(DOCKER_RUN_OPTS) $(YQ_IMAGE) $(1)
endef

# Function to run openapi-generator via Docker
define generator
	docker run $(DOCKER_RUN_OPTS) $(GENERATOR_IMAGE) $(1)
endef

# Show modifiable paths
show-paths:
	@echo "Available paths in the OpenAPI file:"
	@$(call yq,eval '.paths | keys' $(OPENAPI_FILE))

# Modify OpenAPI schema
# Usage: make modify API_PATH=<path> SCHEMA=<media-type>
modify:
	@if [ -z "$(API_PATH)" ] || [ -z "$(SCHEMA)" ]; then \
		echo "Error: Both API_PATH and SCHEMA are required"; \
		exit 1; \
	fi
	@echo "Validating schema '$(SCHEMA)' exists for POST $(API_PATH)..."
	@$(call yq,eval '.paths["$(API_PATH)"].post.requestBody.content["$(SCHEMA)"]' $(OPENAPI_FILE)) > /dev/null || { \
		echo "Error: Schema type '$(SCHEMA)' not found in the requestBody."; \
		exit 1; \
	}
	@echo "Keeping only schema '$(SCHEMA)' in POST $(API_PATH)..."
	@$(call yq,eval -i '.paths["$(API_PATH)"].post.requestBody.content = {"$(SCHEMA)": .paths["$(API_PATH)"].post.requestBody.content["$(SCHEMA)"]}' $(OPENAPI_FILE))
	@echo "OpenAPI file updated."

# Regenerate code via Docker
regenerate-code:
	@echo "Regenerating code from OpenAPI spec using Docker..."
	@$(call generator,generate -i $(OPENAPI_FILE) -g go -o $(OUTPUT_DIR))
	@echo "Code regeneration complete."
