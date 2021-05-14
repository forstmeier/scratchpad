# ui

> The user interface :sloth:

## General

Some basic information for working with the UI.

### Setup

Several prerequisites are needed in order to run the `ui` package locally. Follow the installation/download instructions for your local operating system.

- install [Git](https://git-scm.com/downloads)
	- [additional installation information](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- install [Yarn](https://classic.yarnpkg.com/en/docs/install)
- install dependencies
	- run `yarn install`

### Configurations

There are several configuration options for running the frontend package.

- `yarn serve`: runs the UI with the expectation the backend API is separately available
- `yarn mock`: runs the UI with mock data returned where calls would be made to the API

Going forward there will also likely be options for:

- mock Auth0 client (toggling "live" and mock responses)
- local backend (running Dgraph instance with minimal test data)

## Notes

### Ports

These are the ports that are currently configured for use across the `ui` package.

- ui server: `4000`
	- the Auth0 tenant is configured for this port
