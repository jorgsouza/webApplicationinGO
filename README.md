# Web Application Study

This repository is exclusively for the study of go lang in a web application

As the compose.yaml file is inside the infra folder and not in the root as usual, it's necessary to include `-f infra/compose.yaml`.  
This way, docker locates the file to upload the container with the correct settings.

`docker compose -f infra/compose.yaml up -d`  
`docker compose -f infra/compose.yaml down` **=>** *Command to delete docker, if you use the command everything will be lost*

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

## docker stop
`docker ps` => list running containers.  
`docker stop <nome do container>` **=>** *Command to stop container, the container will not be deleted*


