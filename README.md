# GoApp-ec2-deployment

Minimal REST API in Go intended for containerized deployment (CI/CD to Docker Hub and EC2). This repository contains a small set of HTTP handlers, models, and utilities.

## Table of contents

- [Project Overview](#project-overview)
- [Architecture & Key Files](#architecture--key-files)
- [Prerequisites](#prerequisites)
- [Environment](#environment)
- [Run locally](#run-locally)
- [Docker](#docker)
- [CI / CD](#ci--cd)
- [API Endpoints](#api-endpoints)
- [Models & Utilities](#models--utilities)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)
- [License](#license)

## Project Overview

This project exposes a few lightweight API endpoints that return JSON data. It is designed to be built inside a container and pushed to Docker Hub as part of a GitHub Actions workflow.

## Architecture & Key Files

- Application entry: [main.go](main.go) — registers handlers and reads runtime environment.
- Module declaration: [go.mod](go.mod).
- HTTP handlers:
  - [`handlers.HelloHandler`](handlers/hello.go) — root greeting ([handlers/hello.go](handlers/hello.go)).
  - [`handlers.ItemsHandler`](handlers/items.go) — returns a sample list of items ([handlers/items.go](handlers/items.go)).
  - [`handlers.GetRandomUser`](handlers/randomuser.go) — fetches data from the external randomuser API ([handlers/randomuser.go](handlers/randomuser.go)).
- Models:
  - [`models.Message`](models/message.go) ([models/message.go](models/message.go))
  - [`models.Item`](models/item.go) ([models/item.go](models/item.go))
  - [`models.UserData`](models/models.go) — structure used for randomuser API responses ([models/models.go](models/models.go))
- Utilities:
  - [`utils.EncodeJSON`](utils/json.go) — helper for JSON encoding and error handling ([utils/json.go](utils/json.go))
- Containerization:
  - [Dockerfile](Dockerfile)
- CI:
  - GitHub Actions workflow: [.github/workflows/cicd.yml](.github/workflows/cicd.yml)

## Prerequisites

- Go 1.22.2 (matches `go.mod`)
- Docker (if building images)
- Git (for cloning & workflow operations)
- Recommended: set GOPROXY if behind restrictive networks

## Environment

Runtime reads environment variables via an `.env` created in CI or local environment.

- PORT — port the server listens on (default in code uses `PORT` env var or falls back to `8080`).
  - The containerization and CI may map ports differently; verify `Dockerfile` EXPOSE and CI run commands.

See [main.go](main.go) for how environment vars are loaded and used.

## Run locally

1. Ensure dependencies are tidy:

   ```bash
   go mod tidy

  Run the server:


go run main.go
Visit endpoints:

http://localhost:8080/ -> handled by handlers.HelloHandler
http://localhost:8080/items -> handled by handlers.ItemsHandler
http://localhost:8080/randomuser -> handled by handlers.GetRandomUser
If you want to run a built binary:


go build -o main [main.go](http://_vscodecontentref_/0)./main
Docker
The repository includes a Dockerfile. Example build and run:


docker build -t youruser/goapp .docker run -e PORT=4040 -p 4040:4040 youruser/goapp
Notes:

The Dockerfile uses golang:1.22.2-alpine.
Confirm that the PORT env and EXPOSE in Dockerfile match your runtime mapping. The app default in code is 8080; the Dockerfile exposes 4040 — ensure you set PORT accordingly when running the container.
CI / CD
This repository contains a GitHub Actions workflow at .github/workflows/cicd.yml. Key points:

On push to branch deploy-to-ec2 it builds a Docker image and pushes to Docker Hub.
Workflow creates .env using a secret for PORT.
Docker Hub creds are required in repository secrets (DOCKER_USERNAME, DOCKER_PASSWORD).
The workflow tags and pushes theewizardorne/theewizardone-goapp:latest. Adjust names as needed.
Open workflow: .github/workflows/cicd.yml

API Endpoints
GET / — greeting JSON (handled by handlers.HelloHandler). Response uses models.Message.
GET /items — returns a JSON array of models.Item (handled by handlers.ItemsHandler).
GET /randomuser — proxied data from randomuser.me, parsed into models.UserData and returns the first user (handled by handlers.GetRandomUser).
Handlers use an internal encoder function (encodeJSON in handlers) and there is a reusable helper at utils.EncodeJSON.

Models & Utilities
models.Item — simple item schema.
models.Message — text message wrapper
models.UserData — types to unmarshal randomuser API responses.
utils.EncodeJSON — centralized JSON encoding + error handling.
