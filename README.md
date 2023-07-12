# Item Server

The Item Server offers RESTful HTTP APIs that enable the creation, update, retrieval, and deletion of items. For interaction with other services in a Microservices Architecture (MSA), it utilizes gRPC APIs.

---

## Tech stack

Language: Go (1.20)
Framework: Fiber (for HTTP server), gRPC (for RPC server)
Database: MongoDB

---

## How to run

Set environment variables

```shell
export MONGODB_URI="MONGODB_URI"
export MONGODB_DATABASE="DATABASE_NAME"
export AUTH_SERVICE_ENDPOINT="ENDPOINT"
```

Install dependencies

```shell
go mod download
```

Run HTTP server

```shell
go run cmd/server/main/http.go
```

## How to build using docker

Build image for HTTP app

```shell
docker build -f dockerfiles/build-http/Dockerfile .
```
