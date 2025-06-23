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

### Delivery Service Target
For demonstration purposes, a URL created with `webhook.site` is used as SMS delivery service target. Requests are sent to this URL by default. It can be changed in `.env` file, modifying `DELIVERY_URL` variable. This might be needed since free webhook.site URLs have 7 days of lifetime and can take up to 100 requests.


If the current URL is broken, you can create a new one and set the response as followed:

```
Status Code: 202

Content-Type: application/json

Content:
{
"message": "Accepted",
"messageId": "123"
}


Timeout: 0
```
Set the new URL in the `.env` file, and it will be enough for SMS Service to run again.

## Architecture
Although the project is a submission, architecture is designed by production principles and maintainability.

Rules of clean architecture are followed by default. The code is written with readability and clarity in mind.

### Stack
- gin-gonic as the web server
- pgx-v5 as PostgreSQL driver
- go-redis as official Redis client

### Directory Structure
To follow the best practices, modern Go project layout is used for the project.

- `/api` holds all the files on public-facing API including router, handlers and request/response structs.
- `/cmd` includes the only binary in the project, sms-service
- `internal` includes all the Go code for the inner logic and implementations, excluding API
- `.env.example` example of `.env` file. Need to be renamed (or copied)