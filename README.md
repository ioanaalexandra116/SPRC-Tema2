## General Description
This is an app for storing meteorological data and for providing information based on geographical specifications.

## Architecture
- REST API (Golang)
- Database (PostgreSQL)
- web-based GUI management application (pgAdmin 4)

## Pre-requisites
-   [Golang](https://golang.org/dl/)
-   [Docker](https://docs.docker.com/engine/install/)

## Docker services
- postgres (port 5432)
- pgadmin (port 8001: available [here](http://localhost:8001/))
    - connect with credentials admin@admin.com (Email Adress) and admin (Password) then add a new server and connect to your database using the data from your .env file
- go-api (port 6000: you can send requests to http://localhost:6000/api)

You need to have a .env file with the following content:
```bash
DB_HOST="postgres" # db container name
DB_PORT="5432"
DB_USER="admindb" # if you would like to use another username, change ownership related statements in ./init_scripts/create_database.sh
DB_PASSWORD="<your-database-password>"
DB_NAME="<your-database-name>"
```

## How to run

```bash
go mod init main # generate go.mod file
go mod tidy # generate go.sum file
source .env # export env variables
docker compose -f docker-compose.yml up # run the app in a container
docker compose -f docker-compose.yml down # shuts down the container but keeps the volumes
docker compose -f docker-compose.yml down --volumes # shuts down the container and deletes the volumes (the database will be empty)
```

## Features (what you can do)
- add a country
- display all countries
- modify/delete an existing country
- add a city to an existing country
- display all cities
- display all cities from a country
- modify/delete an existing city
- add a temperature in Celsius degrees to an existing city
- display all temperatures
- display temperatures based on:
    - longitude
    - latitude
    - date (from ... until ...)
    - city
    - country
- modify/delete an existing temperature
