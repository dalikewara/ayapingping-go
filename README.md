# ayapingping-go

**ayapingping-go** generates standard project structure for building
applications in Golang that follow Clean Architecture and Domain-Driven Design concept.

## Getting started

### Installation

You can use the `go get` method:

```bash
go get github.com/dalikewara/ayapingping-go
```

### Usage

To generate new project, just simply run `ayapingping-go` command:

```bash
ayapingping-go
```

then enter your project name & your go module path. After you confirm your inputs,
the **ayapingping-go** generator will setup the project for you.

### What's next?

Just start working on your project, make changes. If you want to run the example app,
you can use `go run` method:

```bash
go run app/api/main.go
```

or you can use `make start` command:

```bash
make start
```


## Project structure

To implement the concept of Clean Architecture and Domain-Driven Design, and to keep them understandable,
we try to structure the project in such a way.

### app

Main applications of the project. Here you place executables of your application, what frameworks used, presenters,
services, and anything between client and business domain. In this folder, you also make dependecy injection
for business domain, like; initialize database connection for business repository, initialize repository for business use case, etc.

### config

Main configurations of the project. Here you place functions to set up configurations or anything about configurations of
your application.

### database

Main databases adapter of the project. Here you place functions to connect to the database or anything about database connections
of your application.

### domain

Main business domains of the project. Here you place main business models, business logics, business helpers, etc
of your application. Business domain **should be isolated** from the outside world. Business domain **should knows nothing** about
anything outside. Business domain **should not depend on** any frameworks.

### Make your own

You're free to create your own style to suit your requirements as long as
still follows the main architecture concept. You can create folders like; `infra` to place
anything about infrastructure such as Dockerfile, `dev` to place anything to run your
application in development mode, `migration` to place your database migrations, etc.

## Release

### Changelog

Read at [CHANGELOG.md](https://github.com/dalikewara/ayapingping-go/blob/master/CHANGELOG.md)

### Credits

Copyright &copy; 2021 [Dali Kewara](https://www.dalikewara.com)

### License

[GNU General Public License v3](https://github.com/dalikewara/ayapingping-go/blob/master/LICENSE)