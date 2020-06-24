# Sakespoke

This repository contains the source code for a simple application that, given a Pokemon name, displays its Shakespearian translation.

## TODO
- Toggle favs
- UX/UI favs
- Fix docker compose

## Quick Start

[Make sure you have docker installed](https://docs.docker.com/get-docker/)

```sh
docker-compose up
```

## Code Structure

There are two main folders:

- `backend`: contains the Go source code for the backend
- `frontend`: contains the React (Typescript) source code for the web application

## Running Locally / Independetly

If you want to run the project locally, or if only one of the two parts is required, you need to make sure the proper dependencies are installed:

### Backend

- Have Go 1.14 installed: https://golang.org/doc/install
- Have the project properly configured in `GOPATH`: https://golang.org/doc/gopath_code.html
- Fill and source a valid environment file (check `backend/.env.template`), by providing a valid value for `PORT`, `POKEMON_API_URL`, and `SHAKESPEARE_API_URL`.
- Run the start command

```sh
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
