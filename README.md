# Web Application Study

This repository is exclusively for the study of go lang in a web application

foi utilizado desta forma pois o arquivo yaml não esta na raiz
`docker compose -f infra/compose.yaml up -d`
`docker compose -f infra/compose.yaml down` **=>** *Eu deleto a máquina, ou seja perco tudo que foi criado no banco de dados*

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
`docker ps` => descobrir qual é o docker do postgres.
`docker stop <nome do container>` **=>** *desta forma o banco não será deletado.*


