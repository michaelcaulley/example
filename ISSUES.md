1. gql_node.go has an extra "github.com/99designs/gqlgen/graphql" import
   1. run `make generate` to see the issue. I've manually removed it for now.
2. `make migration` does not use a hash when constraint names are longer than what Postgres allows
