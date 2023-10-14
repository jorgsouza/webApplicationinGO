# Web Application Study

This repository is exclusively for the study of go lang in a web application

As the compose.yaml file is inside the infra folder and not in the root as usual, it's necessary to include `-f infra/compose.yaml`.  
This way, docker locates the file to upload the container with the correct settings.


*=>** *The command is used to start Docker Compose services defined in a specific compose file.* **<=**
```shell
docker compose -f infra/compose.yaml up -d  

docker compose -f infra/compose.yaml down
```

## .env && compose.yaml
Created an .env file to add the password:  
```env
POSTGRES_PASSWORD=pass
```



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
```shell
psql --host=localhost --username=postgres --port=5432
```

### create database
```sql
CREATE DATABASE store
```

### create a table
```sql
CREATE TABLE products (
    id serial primary key,
    name varchar,
    description varchar,
    price decimal,
    quantity integer
);

```
### insert data
```sql
INSERT INTO products (name,description,price,quantity) VALUES ('Glasses', 'Hyper Technology Black Glasses', 999.99, 2)
```

## Stopping Docker Container

**=>** The docker ps command is used to list all Docker containers running on the host **<=**
```shell
docker ps
```
**=>** If you want to shut down the container, just use the docker stop command. It takes the ID or name of the container as an argument.**<=**
```shell
docker stop <nome do container>
```



# Run
```go
go run main.go
```


***NOTE:*** The go application has two .go files inside the infra folder, starDatabase.go and stopDatabase.go which are responsible for raising the container when the application UP or stopping the container when the application is DOWN


