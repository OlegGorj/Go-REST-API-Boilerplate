# Go-REST-API-Boilerplate
Boilerplate designed to get you up and running with a project structure optimized for developing REST API services in Go

It promotes the best practices that follow the [SOLID principles](https://en.wikipedia.org/wiki/SOLID)
and [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

The Boilerplate provides the following features right out of the box:

* RESTful endpoints in the widely accepted format
* Standard CRUD operations of a database table
* JWT-based authentication
* Environment dependent application configuration management
* Structured logging with contextual information
* Error handling with proper error response generation
* Database migration
* Data validation
* Full test coverage
* Live reloading during development

The kit uses the following Go packages which can be easily replaced with your own favorite ones:

* Routing: [ozzo-routing](https://github.com/go-ozzo/ozzo-routing)
* Database access: [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx)
* Database migration: [golang-migrate](https://github.com/golang-migrate/migrate)
* Data validation: [ozzo-validation](https://github.com/go-ozzo/ozzo-validation)
* Logging: [zap](https://github.com/uber-go/zap)
* JWT: [jwt-go](https://github.com/dgrijalva/jwt-go)


## Getting Started

Implementation of API concists of a RESTful API server running at `http://127.0.0.1:8080`.
It provides the following endpoints:

* `GET /health`: a health check
* `POST /v1/login`: authenticates a user and generates a JWT
* `GET /v1/projects`: returns a paginated list of the projects
* `GET /v1/projects/:id`: returns the detailed information of an project
* `POST /v1/projects`: creates a new project
* `PUT /v1/projects/:id`: updates an existing project
* `DELETE /v1/projects/:id`: deletes an project


## Project Layout




