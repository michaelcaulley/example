# Example

This project is a simple example of using [ent](https://entgo.io) with [gqlgen](https://gqlgen.com) to create a GraphQL API.

# Getting Started

## Prerequisites

- [Docker](https://www.docker.com/products/docker-desktop)

## Running the project

Run the following command to initialize the project:

```bash
make init
```

Once the initialization is complete, you can visit the GraphQL playground at [http://localhost:8080/graphiql](http://localhost:8080/graphiql).

Try creating a user with the following mutation:

```graphql
mutation createUser {
    createUser(input: { name: "Michael" }) {
        name
    }
}
```

Initialization is only required once, unless you wish to rebuild images. You can take the application up and down with:

```bash
make up
```

and

```bash
make down
```

## Code Generation

To generate ent code, run:

```bash
make generate
```

To generate an ent versioned migration, run:

```bash
make ent-migration
```

## Other Commands

Refer to the [Makefile](Makefile) for other commands.