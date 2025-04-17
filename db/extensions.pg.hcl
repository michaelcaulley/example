#  extensions.pg.hcl defines the PostgreSQL extensions used in the application.

schema "public" {}

extension "pg_trgm" {
  schema = schema.public
}

extension "ltree" {
  schema = schema.public
}
