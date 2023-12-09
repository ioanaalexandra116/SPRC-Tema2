## General Description
This is an app for storing meteorological data and for providing information based on geographical specifications.

## Architecture
- Database (PostgreSQL)
- web-based GUI management application (pgAdmin 4)
- REST API (Golang)

## Docker services
- postgres (port 5432)
- pgadmin (port 8001: available [here](http://localhost:8001/))
    - connect with credentials admin@admin.com (Email Adress) and admin (Password) then add a new server and connect to your database using the data from the .env file (host: postgres, maintenance database:  postgres, username: admindb, password: password)
- go-api (port 6000: you can send requests to http://localhost:6000/api)

## Pre-requisites
-   [Docker](https://docs.docker.com/engine/install/)

## How to run

```bash
docker compose -f docker-compose.yml up # run the app in a container
docker compose -f docker-compose.yml down # shuts down the container but keeps the volumes
docker compose -f docker-compose.yml down --volumes # shuts down the container and deletes the volumes (the database will be empty)
docker images | grep 'go-api' | awk '{print $1}' | xargs docker image rm # remove the api image
```
For convenience, I have included the file .env

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
