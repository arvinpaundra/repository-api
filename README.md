# Repository Scientific API

An API for scientific repositories like [EPrints](https://www.eprints.org/uk), [DSpace](https://dspace.lyrasis.org), etc. The output of this application are to upload and search student or lecturer or somebody's paper, thesis, etc. This application could be a reference especially for students to write a scientific papers. Last but not least, this application also can generate a library identity card with pdf output.

## Table of Contents

- [Usage](#usage)
- [Tech Stack](#Tech-Stack)
- [API Documentation](#API-Documentation)
- [Testing](#Testing)

## Usage

### Prerequisite

- Golang `migrate` installed

This project using golang migrate to perform database migration or versioning. Alternatively, to perform database migration could use builtin migration with gorm `AutoMigrate` since this project using gorm. Choose wisely!

- `Wkhtmltopdf` installed

For generated pdf, I prefer to using wkhtmltopdf which is a CLI apps for generate pdf from html file. So, make sure to install this binary too.

### Run application

After fullfil the prerequisite above, clone or download this repo, then you can follow the steps below :

- Download all dependencies

```sh
go mod download
```

- Perform migrate database using golang `migrate`

```sh
migrate -database "mysql://uname:pass@tcp(localhost:3306)/db_name" -path ./migrations up
```

- Run with no worries

```sh
go run ./app/main.go
```

## Tech Stack

- [Golang](https://go.dev/) - The Protagonist
- [Echo](https://echo.labstack.com/) - Web Framework
- [Gorm](https://gorm.io/) - ORM
- [Migrate](https://github.com/golang-migrate/migrate) - Tool for database migration
- [go-redis](https://github.com/go-redis/redis) - Type-safe Redis client
- [Swagger](https://swagger.io) - API Docs
- [Testify](https://github.com/stretchr/testify) - Testing library
- [Mockery](https://github.com/vektra/mockery) - Mock generator

## API Documentation

This project using [Swagger](https://swagger.io) for API Documentation or API Specification and always update the docs when each feature is done. For the documentation, you can visit this link. [API Docs](https://example.com)

## Testing

Do run the test with command snippet below :

```sh
go test -v -cover ./services/...
```

For more information about test coverage and more readable (generated into html), then do run this command :

```sh
go test -v -coverprofile=cover.out ./services/... && go tool cover -html=cover.out
```
