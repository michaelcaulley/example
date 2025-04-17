API_TEST_FILES ?= ./...

readme:
	@echo "👀 Open the Makefile to see all commands and comments"

# Initialize the project for local development.
init: down _rebuild migrate
	docker compose up -d

# Stand up all containers and networks.
up:
	docker compose up -d

# Tear down all containers and networks.
down:
	docker compose down

# Run the code generation steps.
generate:
	docker compose run --rm go-runner go generate .
	docker compose restart api

# Create a versioned migration syncing the ent schema.
# Create a migration syncing the ent schema to the database.
migration: _migration hash
_migration:
	@echo "📘 Make sure you are logged in to Atlas before running this command. Run 'make atlas-whoami' to check."
	@read -p "✏️  Migration name (ex. add_user_middle_name_field): " name; \
	docker compose run --rm atlas migrate diff $$name \
                                    --env local || true # continue on error
	@docker compose down atlas-db -v # Apply all migrations to the database.

migrate: up
	docker compose run --rm atlas migrate apply \
		--url "postgres://example_user:password@db:5432/example?sslmode=disable"
	docker compose down atlas-db

# Recalculate the atlas migration hash.
hash:
	docker compose run --rm atlas migrate hash || true # continue on error
	docker compose down atlas-db

# Create a test postgres instance and run the API tests.
api-test:
	docker compose down api-test-db -v
	docker compose up api-test-db -d
	docker compose exec api-test-db sh -c 'until pg_isready -h api-test-db -U postgres; do sleep 1; done'
	docker compose down atlas-db -v
	@docker compose run --rm atlas migrate apply \
		--url "postgres://postgres:password@api-test-db:5432/postgres?sslmode=disable"
	@docker compose exec api-test-db psql -h localhost -U postgres -d postgres -c "SELECT 1 FROM pg_database WHERE datname = 'api_test_template'" | grep -q 1 || docker compose exec api-test-db psql -h localhost -U postgres -d postgres -c "CREATE DATABASE api_test_template TEMPLATE postgres OWNER postgres;"
	@docker compose exec api-test-db psql -h localhost -U postgres -d postgres -c "UPDATE pg_database SET datistemplate = true WHERE datname = 'api_test_template';"
	cd api && ENVIRONMENT=test go test -v $(API_TEST_FILES)

# DO NOT CALL DIRECTLY. This is used by the init command to rebuild the containers.
_rebuild:
	docker rm -f example-api example-db
	docker volume rm -f example_database example_go_cache
	docker compose up --build --force-recreate --no-deps -d --wait
