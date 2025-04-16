# Stand up all containers and networks.
up:
	docker compose up -d

# Tear down all containers and networks.
down:
	docker compose down

# Run the code generation steps.
generate:
	docker compose run --rm go-runner go generate .

# Create a migration syncing the ent schema to the database.
ent-migration: up _ent-migration hash
_ent-migration:
	@read -p "Migration name (ex. add_user_middle_name_field): " name; \
	docker compose run --rm atlas migrate diff $$name \
		--dir "file://migrations" \
		--to "ent://internal/ent/schema" \
		--dev-url "postgres://postgres:password@atlas-db:5432/postgres?search_path=public&sslmode=disable"

# Apply all migrations to the database.
migrate: up
	docker compose run --rm atlas migrate apply \
		--url "postgres://example_user:password@db:5432/example?sslmode=disable"
	docker compose down atlas-db

# Recalculate the atlas migration hash.
hash:
	docker compose run --rm atlas migrate hash || true # continue on error
	docker compose down atlas-db
