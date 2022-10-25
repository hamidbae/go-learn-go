# About the project

this is rest api project to complete 1 month golang coding bootcamp at Hacktive8 x FGA academy

# Built with

- Golang 1.19.1
- Postgresql 14.1
- Gin framework
- Vscode
- AWS EC2
- Github Action

# Getting started

## Prerequisites

- Golang
- Postgresql
- Vscode

## Installation

clone this repo

copy `.env.example` to `.env`

create db

fill `.env` file

run `swag init`

uncomment this code on main.go

```
// err := godotenv.Load(".env")

// if err != nil {
//   log.Fatalf("Error loading .env file")
// }
```

run `go mod tidy`

run `go run main.go`

go to http://localhost:8080/swagger/index.html

now you are able to test endpoint locally

# Documentation

local baseurl : http://localhost:8080/api

local docs url : http://localhost:8080/swagger/index.html

deployed baseurl : http://44.201.153.46/api

deployed docs url : http://44.201.153.46/swagger/index.html

postman documentation : https://documenter.getpostman.com/view/12388903/2s84LStpgT

# Notes For Author

run postgres container

`docker run --name postgres-fga --network go-fga -p 5432:5432 -e POSTGRES_USER=username -e POSTGRES_PASSWORD=password -e POSTGRES_DB=final-project -d postgres`

run container but autoremove when stopped

`docker run --rm --name go-fga-final-project -p 8080:8080 hamidbae/go-fga:latest`

godotenv doesn't required when using --env-file on docker

https://stackoverflow.com/questions/66314534/env-variables-not-coming-through-godotenv-docker
