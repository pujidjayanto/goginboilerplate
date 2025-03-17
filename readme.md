# Gin Golang Layered Structure Boilerplate

This repository is a boilerplate for creating API service using Gin. This boilerplate using layered structured and
inspired by Ruby on Rails project layout.

# Directory structure

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
│   │   └── user
│   │       └── controller.go
│   ├── dto
│   │   ├── login.go
│   │   └── register.go
│   ├── middleware
│   ├── repository
│   │   ├── dependency.go
│   │   └── user
│   │       ├── model.go
│   │       ├── repository.go
│   │       └── repository_test.go
│   ├── service
│   │   ├── dependency.go
│   │   └── user
│   │       ├── error.go
│   │       └── service.go
│   └── testutils
|
├── migrations
│   ├── 20250307220522_init_tables.sql
│   └── seeds
│       └── 20250307225245_init_seed.sql
├── mocks
|
├── pkg
|
└── readme.md
```

## Directory explanation

### `/cmd`
Main applications for this project. I create 2 main applications:
1. `/cmd/migrator` for do migration and seed.
2. `/cmd/server` for the api server application.

### `/internal`
This is where you put all the business related code which is private. This is the code you don't want others importing in their applications or libraries.

### `/internal/app`
This directory for routes and application server setup.

### `internal/config`
This directory for loading the environment variable. If you have new variable, register it there.

### `internal/controller`
This directory for controller or handler related code. I use sub-directory approach (a folder for each controller) instead of file approach (a file like `user_controller.go` or `user_handler.go`). It is because i got inspired from rails where one controller file has a lot of methods behind one controller that's not defined explicitly. If we do that in go with file approach, we will have a lots of files in this controller directory or have one controller with very long code. Keep those in a folder is more nicer.

If you want to add a new controller, you must create a directory, and register it in `internal/controller/dependency.go` so it will be available for your routes.

### `internal/service`
This directory for service or use-case related code. Using same approach with controller. If you want to add a new service, you must create a directory, and register it in `internal/service/dependency.go` so it will be available for your controller.

### `internal/repository`
This directory for repository or model related code. Using same approach with controller, and service. If you want to add a new repository, you must create a directory, and register it in `internal/repository/dependency.go` so it will be available for your service.

### `internal/dto`
This directory for defining a request and response.

### `internal/middleware`
This directory related to middleware.

### `internal/testutils`
This directory for defining test helpers method that you will be used for creating test files.

### `/migrations`
This is where `.sql` files related to migration and seed.

### `/pkg`
This directory for wrapping library and get imported in `/internal` code or OK to be exported. Not only for wrapping a library, but if you have piece of code or your own library that is OK to be exported, you can put it here.

# Migration
This boilerplate include a migration application in `/cmd/migrator`. You can run `go run ./cmd/migrator -help` to see available commands to run the migrator.
```bash
Options:
  -create-migration string
    	Create a new migration file
  -create-seed string
    	Create a new seed file
  -migrate
    	Run database migrations
  -seed
    	Run database seeds

Examples:
  Create a new migration:
    go run ./cmd/migrator -create-migration init_tables
  Create a new seed:
    go run ./cmd/migrator -create-seed seed_init
  Run migrations up:
    go run ./cmd/migrator -migrate up
  Run migrations down:
    go run ./cmd/migrator -migrate down
  Run seeds:
    go run ./cmd/migrator -seed
```

Migration app uses and produce `.sql` files in `/migrations`. I separate between migration and seeder and have an history table for each, so we know what migrations and seed that already applied.

# Up and Running
1. Prepare 2 database (if development). 1 for the main database and 1 for test database
2. Make `.env` file. You can see the `.env.example`.
3. Create a migration files and seed if needed, and migrate it with `go run ./cmd/migrator -migrate up` and then `go run ./cmd/migrator -seed` to seed data to main database.
4. Run `go mod download` or `go mod tidy` for sync the library
5. Run the application with `go run ./cmd/server`.


# Testing
As mentioned before testing use a test database for unit testing (repository). But use mock provided by [mockery](https://vektra.github.io/mockery/) to test service and controller layer. Make sure you have it installed in your local and run `mockery` to update the mock files in `./mocks` directory.

Run the test by `go test ./... -v` for verbose log.


# Use as Template
You can use this as template for your gin golang project by clicking `Use this template` in github and make sure to change the module name in `go.mod` and in all code.

I prepare 3 domain which are user, products, and purchases just for demo purposes. But you can delete them if no longer needed.


# Library
These are libraries used:
1. Gin for http handlers.
2. GORM for database operation.
3. Mockery for testing.
4. Goose for migration.
5. [Sonic](https://github.com/bytedance/sonic) for handling json type.
6. Zap for logging purposes.
7. [Decimal](https://github.com/shopspring/decimal) for handling decimal type.
8. Testify for testing.


# Contributing
All pull requests are welcome


# License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details