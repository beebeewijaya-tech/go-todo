# Go-Todo

Welcome to the Go Todo List Project!

## Introduction

This project is a simple but powerful todo list management tool built with the Go programming language. It allows users to create, read, update and delete tasks.
To get started with this project, follow the instructions in the README file to install the necessary dependencies and run the application. Once you're up and running, you'll be able to create tasks, set their priority, mark them as complete, and view a list of all your tasks.

## Tech Stack

- Go
- Gin
- Sqlx
- JWT
- Postgres
- Makefile

## Installation

```
git clone https://github.com/beebeewijaya-tech/go-todo

cp Makefile.sample Makefile

cp env/sample.config.json env/config.json

go mod tidy

```

Remember change the `Makefile` scripts to insert your PostgreSQL credentials there.

Also remember to add the `config.json` value to the env as well

## Folder Structure

```bash
├───api
├───db
│   ├───migrations
│   └───sql
│       └───mock_sql
├───env
├───middleware
├───token
└───util
```

`api` will represent the gin server for creating the api

`db` will represent SQLX code for database communication

`db/migrations` your migrations file will lives here

`db/sql` your SQL query and mutation will lives here

`env` insert your env here

`middleware` middleware such as authToken, log, csrf will lives here

`token` JWT utils or other token implementation such as PASETO

`util` generic utilities will lives here
