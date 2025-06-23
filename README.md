# SMS Service
## Introduction
SMS Service stores messages and sends them periodically to a specific SMS delivery service. This project is intended to be a submission, therefore it should not be considered for production.

The repository is set to work seamlessly with Docker. Variables defined in the environment file is by default for Docker Compose.

**Note for review:** This project is written by me, Rojhat Sinan Balka, without any help from AI (except Goland's auto-completion)

## Running
To start the service, 

Rename `.env.example` to `.env` and simply run:
```shell
    make up
```

If you made changes in the environment file, run build first:
```shell
  make build
  make up
```

Made changes, left out things, not sure which version is working? run:
```shell
    make all
```
This will completely save you from mess.

Documentation for API, Swagger(OAPI) and Postman files can be found in `/docs` directory.


## Architecture
Although the project is a submission, architecture is designed by production principles and maintainability.

## Stack
- gin-gonic as the web server
- pgx-v5 as PostgreSQL driver
- go-redis as official Redis client

## Directory Structure
To follow the best practices, modern Go project layout is used for the project.

- `/api` holds all the files on public-facing API including router, handlers and request/response structs.
- `/cmd` includes the only binary in the project, sms-service
- `internal` includes all the Go code for the inner logic and implementations, excluding API
- `.env.example` example of `.env` file. Need to be renamed (or copied)