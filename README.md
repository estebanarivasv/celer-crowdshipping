# Celer

### Install steps

1. Create a `.env` file based in the `.env-example` file. This file contains the environment variables for the database where Camunda Engine will run.
2. Run `docker compose up` so that the containers are created.
3. Done! You should have a functional demo now.

### External endpoints

- Golang API base url: http://localhost:5000/
- Golang API swagger: http://localhost:5000/docs/swagger/index.html

- Camunda Welcome page: http://localhost:8080/camunda/app/welcome/default/#!/welcome
- Camunda Cockpit: http://localhost:8080/camunda/app/cockpit/default/#/dashboard

- Angular client: http://localhost:4200/


### Requirements
- Docker and docker-compose

#### If you want to run it locally:
- Go v1.18.10
- node v16.15.0
- npm 8.5.5


### Important resources
- REST API Reference - Camunda Platform 7.17.0 - https://docs.camunda.org/manual/7.17/reference/rest/
