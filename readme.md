# Gin Golang Layered Structure Boilerplate

This repository is a boilerplate for creating API service using Gin. This boilerplate using layered structured and
inspired by Ruby on Rails project layout.

# Folder structure

```bash
.
├── cmd
│   ├── migrator
│   │   ├── config.go
│   │   └── main.go
│   └── server
│       └── main.go
├── domains
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   ├── app.go
│   │   └── routes.go
│   ├── config
│   │   └── config.go
│   ├── controller
│   │   ├── dependency.go
│   │   ├── product
│   │   │   ├── controller.go
│   │   │   └── controller_test.go
│   │   ├── purchase
│   │   │   └── controller.go
│   │   └── user
│   │       └── controller.go
│   ├── dto
│   │   ├── login.go
│   │   ├── product.go
│   │   ├── purchase.go
│   │   └── register.go
│   ├── middleware
│   │   ├── auth.go
│   │   ├── header.go
│   │   └── log.go
│   ├── repository
│   │   ├── dependency.go
│   │   ├── product
│   │   │   ├── constants.go
│   │   │   ├── model.go
│   │   │   ├── repository.go
│   │   │   ├── repository_test.go
│   │   │   └── scopes.go
│   │   ├── purchase
│   │   │   ├── model.go
│   │   │   ├── repository.go
│   │   │   └── repository_test.go
│   │   └── user
│   │       ├── model.go
│   │       ├── repository.go
│   │       └── repository_test.go
│   ├── service
│   │   ├── dependency.go
│   │   ├── product
│   │   │   ├── helper.go
│   │   │   ├── service.go
│   │   │   └── service_test.go
│   │   ├── purchase
│   │   │   ├── error.go
│   │   │   └── service.go
│   │   └── user
│   │       ├── error.go
│   │       └── service.go
│   └── testutils
│       ├── db.go
│       └── testdata
│           └── teardown.sql
├── migrations
│   ├── 20250307220522_init_tables.sql
│   └── seeds
│       └── 20250307225245_init_seed.sql
├── mocks
|
├── pkg
│   ├── db
│   │   └── gorm.go
│   ├── delivery
│   │   ├── delivery.go
│   │   └── response.go
│   ├── envloader
│   │   └── envloader.go
│   ├── jsonb
│   │   └── jsonb.go
│   ├── logger
│   │   └── zap.go
│   └── pagination
│       └── pagination.go
└── readme.md
```

## Folder explanation

### `/cmd`
Main applications for this project. I create 2 main applications:
1. `/cmd/migrator` for do migration and seed
2. `/cmd/server` for the api server application

### `/domains`
I keep this as a directory to put all external api call.

### `/internal`
This is where you put all the business related code which is private. This is the code you don't want others importing in their applications or libraries.

### `/migrations`
This is where `.sql` files related to migration and seed

### `/pkg`
This directory for wrapping library and get imported in `/internal` code or OK to be exported.


# Library
These are libraries used:
1. Gin
2. GORM
3. Mockery
4. Goose
5. Sonic
6. Zap


# Todo
1. add example to use redis
2. add example to do external api call