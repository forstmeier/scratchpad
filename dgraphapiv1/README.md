# api

> The API and database :sloth:

## General

Some basic information on getting started with the project.

### Principles

`api` is built around three principles:

- **Automate** _everything_
- **Query** _everything_
- **Reconcile** _everything_

When designing and constructing features, these three principles need to be consulted. For example:

> How can we programmatically extend the functionality provided by the schema-generated queries and mutations?
> How can we design our schema type fields and relationships to provide a superior search experience?
> How can we introduce schema elements that allow item states to be set and tracked over time?

### Setup

Several prerequisites are needed in order to run the `api` package locally. Follow the installation/download instructions for your local operating system.

- install [Git](https://git-scm.com/downloads)
	- [additional installation information](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- install [Go](https://golang.org/doc/install)
- install [Docker](https://docs.docker.com/get-docker/)
- install [Docker Compose](https://docs.docker.com/compose/install/)
- install [Insomnia](https://insomnia.rest/download/core/?&ref=https%3A%2F%2Fgraphql.dgraph.io%2Fdocs%2Fquick-start%2F)

These should be the minimum resources needed to get up and running with `api`! :thumbsup:

## Packages

General notes and instructions on several of the packages included in `api`.

### `bin`

`bin` contains the various helper scripts used to run and interact with the `api` application. See the individual script files for documentation.

### `custom`

`custom` intercepts and processes all Dgraph `@custom` directive requests. This is to allow the additional preprocessing ahead of the main GraphQL logic and to provide a base for local mocking. **No "smarts"** will be built into this package and it will _only be an intermediary_ responsible for invoking external APIs (e.g. Auth0) and internal packages (e.g. `pkg/users`) to fulfill the required logic.

### `demo`

`demo` is responsible for loading demo data into the Dgraph database. Follow the instructions below to get setup and execute all commands in the terminal from the root of the `api` repository.

1. run `chmod -R +x ./bin/` to make all `bin/` scripts executable 
2. run `./bin/run_api` to start **Draph** and the `custom` package
3. run `./bin/load_schema` to insert the **GraphQL** schema into the **Dgraph** database
4. run `./bin/load_demo` to insert demo data into the database
5. launch **Insomnia** and upload the collection file provided by the repo maintainer
6. run `./bin/user_token` or `./bin/app_token` to generate test user and **Auth0 Application** tokens which can be used in the **Insomnia** `X-Auth0-Token` header 

### `token`

`token` fetches either Auth0 Application or test user JWTs for use in testing or demoing. The `./bin/app_token` or `./bin/user_token` scripts can be used to fetch those tokens, respectively.

Currently the user JWT issued from the `util/token` package is configured to `john.forstmeier@gmail.com` in Auth0.

## Notes

### Ports

These are the ports that are currently configured for use across the `api` package.

- Dgraph Zero: `5080` and `6080`
- Dgraph Alpha: `7080`, `8080`, and `9080`
- custom server: `4080`

### Design

Below are a couple of points regarding specific code and design choices.

- `struct` _must_ be used for all request/response objects - specific `struct` objects with appropriate fields and JSON tags should be defined where appropriate
- `map` should be used for Dgraph GraphQL mutation variables - a more flexible `map[string]interface{}` can be used when executing mutations
