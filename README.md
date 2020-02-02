# (<sub><sup>toy project</sup></sub>) Catalog Site Builder

This is work in progress. An attempt to generalize previous projects into a useful utility.

A tool to build/generate a digital catalog according to a flat JSON schema.

## In a nutshell

1. Ingests data from a Sqlite3, exports the tables as JSON arrays (sqlite3, Bash).
2. Indexes the arrays into dictionary (Bash, jq)
3. Generates a GraphQL schema from the dictionaries (Bash, Node)
4. Generates a Golang GraphQL endpoint (gqlgen, Go)
5. Serves the content (React)
