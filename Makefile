BIN := astra
BIN_DIR := ./bin
TMP_DIR := ./tmp

DB_DRIVER := mysql
DB_STRING := $(or $(DB_USER),root):$(or $(DB_PASS),)@tcp($(or $(DB_HOST),localhost):$(or $(DB_PORT),3306))/$(or $(DB_NAME),bookstore)?parseTime=true
MIGRATION_DIR := ./internal/config/db/migrations

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

.PHONY: help
help: ## Display usage information
	@echo "Usage:"
	@awk 'BEGIN {FS = ":.*?## "}; /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

.PHONY: tidy
tidy: ## Run go mod tidy
	@echo "Running go mod tidy..."
	@go mod tidy

.PHONY: clean
clean: ## Clean the project
	@echo "Cleaning the project..."
	@rm -rf $(BIN_DIR)
	@rm -rf $(TMP_DIR)

# ==================================================================================== #
# MIGRATIONS
# ==================================================================================== #

.PHONY: migration-up
migration-up: ## Run migration up
	@echo "Running migration up..."
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING="$(DB_STRING)" goose -dir=$(MIGRATION_DIR) up

.PHONY: migration-down
migration-down: ## Run migration down
	@echo "Running migration down..."
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING="$(DB_STRING)" goose -dir=$(MIGRATION_DIR) down

.PHONY: migration-reset
migration-reset: ## Run migration reset
	@echo "Running migration reset..."
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING="$(DB_STRING)" goose -dir=$(MIGRATION_DIR) reset

.PHONY: migration-status
migration-status: ## Check migration status
	@echo "Checking migration status..."
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING="$(DB_STRING)" goose -dir=$(MIGRATION_DIR) status

# ==================================================================================== #
# APPLICATION
# ==================================================================================== #

.PHONY: test
test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

.PHONY: dev
dev: ## Run development environment
	@echo "Running development environment..."
	@air --build.cmd "go build -o $(BIN_DIR)/$(BIN) cmd/main.go" --build.bin "$(BIN_DIR)/$(BIN)"

.PHONY: build
build: ## Build the project
	@echo "Building the project..."
	@go build -o $(BIN_DIR)/$(BIN) ./cmd

.PHONY: run
run: tidy build ## Run the project
	@echo "Running the project..."
	@$(BIN_DIR)/$(BIN)
