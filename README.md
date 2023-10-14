# Web Application Study

This repository is exclusively for the study of go lang in a web application

As the compose.yaml file is inside the infra folder and not in the root as usual, it's necessary to include `-f infra/compose.yaml`.  
This way, docker locates the file to upload the container with the correct settings.

`docker compose -f infra/compose.yaml up -d`  
`docker compose -f infra/compose.yaml down` **=>** *Command to delete docker, if you use the command everything will be lost*

## .env && compose.yaml
Created an .env file to add the password:  
`POSTGRES_PASSWORD=pass`


***Leaving the yaml file like this:***   

```yaml
services:
database:
image: "postgres:16.0-alpine3.18"
environment:
POSTGRES_PASSWORD: ${POSTGRES_PASSWORD}
ports:
- "5432:5432"
```

## DataBase
### access PostgresSQL
`psql --host=localhost --username=postgres --port=5432`

### create database
`CREATE DATABASE store`

### create a table
`CREATE TABLE products (
    id serial primary key,
    name varchar,
    description varchar,
    price decimal,
    quantity integer
);`
### insert data
`INSERT INTO products (name,description,price,quantity) VALUES ('Glasses', 'Hyper Technology Black Glasses', 999.99, 2)`

## Stopping Docker Container
`docker ps` => list running containers.  
`docker stop <nome do container>` **=>** *Command to stop container, the container will not be deleted*


# Run
`go run main.go`

***NOTE:*** The go application has two .go files inside the infra folder, starDatabase.go and stopDatabase.go which are responsible for raising the container when the application UP or stopping the container when the application is DOWN


