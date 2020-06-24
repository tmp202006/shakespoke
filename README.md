# Sakespoke

This repository contains the source code for a simple application that, given a Pokemon name, displays its Shakespearian translation.

## Quick Start

[Make sure you have docker installed](https://docs.docker.com/get-docker/)

```sh
docker-compose up
```

Once the build has completed, the frontend will be available at `http://localhost:3000`, while the backend will respond at `http://localhost:3001`.

## Code Structure

There are two main folders:

- `backend`: contains the Go source code for the backend
- `frontend`: contains the React (Typescript) source code for the web application

## Running Locally / Independetly

If you want to run the project locally, or if only one of the two parts is required, you need to make sure the proper dependencies are installed:

### Backend

- Have Go 1.14 installed: https://golang.org/doc/install
- Have the project properly configured in `GOPATH`: https://golang.org/doc/gopath_code.html
- Fill and source a valid environment file (check `backend/.env.template`), by providing a valid value for `PORT`, `POKEMON_API_URL`, and `SHAKESPEARE_API_URL`, or set them independently.
- Run the start command

```sh
export PORT=3001
export POKEMON_API_URL="https://pokeapi.co/api/v2"
export SHAKESPEARE_API_URL="https://api.funtranslations.com"
go run main.go
```

Note:

- `PORT` is set as `3001` in the Dockerfile
- `POKEMON_API_URL` should be set to `https://pokeapi.co/api/v2`
- `SHAKESPEARE_API_URL` should be set to `https://api.funtranslations.com`

### Frontend

- Have Node 12 installed: https://nodejs.org/it/download/
- Install dependencies
- Run the start command

```sh
npm install
npm start
```

## TEsting

## Missing

Some part of the project should be improved in order to make this a production ready app:

Frontend:

- Health checks for verifying that backend is up and running
- Add unit tests
- Add e2e tests
- Multi-platform
- Handle all status codes from backend

Backend:

- Handle remote services downtimes
- Test handlers by using mocks

Overall

- Improve input validation
- Review code structure
- CI/CD
