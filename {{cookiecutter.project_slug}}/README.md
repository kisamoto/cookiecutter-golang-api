# {{cookiecutter.project_name}}

A Go powered API 

## Stack

- Uses no API framework
  - `gorilla/mux` for routing 
  - Std Lib HTTP Handler functions
- PostgreSQL for database persistence 
  - No ORM
  - DB Migrations managed by [FlywayDB](https://github.com/flyway/flyway)
- Caching layer backed by Redis
- Inter-API notification layer abstracted using [go-notify](https://github.com/bitly/go-notify)
  - Notifications exist only in the **app instance**
- Dependency management via [Glide](https://github.com/Masterminds/glide)
- Deployed via Docker 
- Web server provided by [Caddy](https://github.com/mholt/caddy)
  - HTTPS via Let's Encrypt
  - `gzip` enabled
  - log to `stdout` (for Docker)
- Logging powered by [Zap](https://github.com/uber-go/zap)

## Configuration Options

**Configuration Filename:** `{{cookiecutter.project_slug}}.[yaml|json]`

**Environment Variable Prefix:** `{{cookiecutter.project_slug}}_`

## Database Migrations

_More info can be found at {{cookiecutter.repo}}/db/README.md_

## Binary Layout

Binary | Purpose
-------|----------
`api` | Main entry point for API. Building and running this binary with the relevant configuration will run the API

## Package Layout

Package | Purpose
--------|----------
`{{cookiecutter.project_slug}}` | Main package to define structs (models) and interfaces to be implemented elsewhere. 
`env` | Environment configuration and setup
`handlers` | HTTP Handler functions to control access to API
`mock` | Mock implementations of service interfaces from `pkg/{{cookiecutter.project_slug}}`
`postgresql` | DB implementations of service interfaces from `pkg/{{cookiecutter.project_slug}}` backed by PostgreSQL

## Models

Some models and services are provided out-of-the-box as they are fairly commonplace in my API implementations.

Model | Purpose
------|---------
`AccountCredentials` | Holds account credentials performing password validation and verification
`AccountProfile` | Holds user preferences; profile information etc.
