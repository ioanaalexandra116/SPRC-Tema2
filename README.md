## General Description
This is an app for storing meteorological data and for providing information based on geographical specifications.

## Architecture:
- REST API (Golang)
- Database (PostgreSQL)
- web-based GUI management application (pgAdmin 4)

## Pre-requisites
-   [Golang](https://golang.org/dl/)
-   [Docker](https://docs.docker.com/engine/install/)

You need to have a .env file with the following content:
```bash
DB_HOST="<your-database-host>"
DB_PORT="<your-database-port>"
DB_USER="<your-database-user>"
DB_PASSWORD="<your-database-password>"
DB_NAME="<your-database-name>"
```

## How to run

```bash
source .env #export env variables
docker compose -f docker-compose.yml up # run the app in a container
docker compose -f docker-compose.yml up # shuts down the container but keeps the volumes
docker compose -f docker-compose.yml up --volumes # shuts down the container and deletes the volumes (the database will be empty)
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
