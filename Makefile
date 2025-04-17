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
ent-migration: up _ent-migration hash

# Apply all migrations to the database.
migrate: up
	docker compose run --rm atlas migrate apply \
		--url "postgres://example_user:password@db:5432/example?sslmode=disable"
	docker compose down atlas-db

# Recalculate the atlas migration hash.
hash:
	docker compose run --rm atlas migrate hash || true # continue on error
	docker compose down atlas-db

# DO NOT CALL DIRECTLY. This is used by the ent-migration command.
_ent-migration:
	@read -p "Migration name (ex. add_user_middle_name_field): " name; \
	docker compose run --rm atlas migrate diff $$name \
		--dir "file://migrations" \
		--to "ent://internal/ent/schema" \
		--dev-url "postgres://postgres:password@atlas-db:5432/postgres?search_path=public&sslmode=disable"

# DO NOT CALL DIRECTLY. This is used by the init command to rebuild the containers.
_rebuild:
	docker rm -f example-api example-db
	docker volume rm -f example_database example_go_cache
	docker compose up --build --force-recreate --no-deps -d --wait