# ayapingping-go

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/dalikewara/ayapingping-go/v3)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/dalikewara/ayapingping-go)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/dalikewara/ayapingping-go)
![GitHub license](https://img.shields.io/github/license/dalikewara/ayapingping-go)

**ayapingping-go** generates standard project structure to build applications in Golang that follow Clean
Architecture and Domain-Driven Design concept.

## Getting started

### Installation

You can use the `go install` method:

```bash
go install github.com/dalikewara/ayapingping-go/v3@latest
```

or, you can also use the `go get` method (DEPRECATED since `go1.17`):

```bash
go get github.com/dalikewara/ayapingping-go/v3
```

### Usage

To generate new project, just simply run `ayapingping-go` command:

```bash
ayapingping-go
```

then enter your project name & your go module path. After you confirm your inputs, the **ayapingping-go** generator will
setup the project for you.

![Alt Text](https://lh3.googleusercontent.com/pw/AM-JKLXHIY-P9tKx2cI0sgdLTxzvK5ErAwkToS-3to790cY4UDg2yullDtehGV2LEtYEDU-a1-xa9t_0vjTJJVri45aDNXN7BLxx-eAxOflZltzzrwF2bILJ9bHQWsCnXtCNDC8tMWZMk4tPtDP1iu9OYmD4=w600-h372-no)

### What's next?

Just start working on your project, make changes.

## Project structure

To implement the concept of Clean Architecture and Domain-Driven Design, and to keep them understandable, we try to
structure the project like this:

### main.go

- In this file, you initialize dependencies, injections, and anything required to start and run your application
- You can use command `go run main.go` or `make start` to run your application

### src/service

- In this package, you put your application handlers, servers or clients. Example:
  - `src/service/rest/`, to handle REST requests from client
  - `src/service/grpc/`, to handle gRPC requests from client
  - `src/service/cron/`, to handle CRON
  - etc...
- Also in this package, you put presenters or anything to be done between client and your application.

### src/config

- In this package, you put any functions to set up constant variables, env variables, messages or anything about configurations
- For constant `const` configuration, you may use the variable directly from anywhere

### src/adapter

- In this package, you put any functions to connect to external frameworks or connections like database, external client, etc

### src/entity

- Entity is your main business model
- Here you define your main object models or properties for your business, like database model, DTO (Data Transfer Object), etc.
- Keep this package as simple as possible. Don't code anything that is not related to the model itself

### src/library

- In this package, you create custom helper functions
- Provide reusable helpers for your application
- You may call any functions inside this package from anywhere

### src/repository

- Repository is a place where you communicate with the real external data resource, like database, cloud service, external service, etc.
- It's always better to keep your repository as simple as possible, don't add too much logic here
- If you have to, you can separate the operations into smaller methods
- You should always call your repository methods inside the use case package
- You may use your `src/library` functions directly in this package
- Any changes outside this package should not affect your repositories (except changes for business entity)
- If you need config variables, database frameworks, or external clients, pass/inject them as a dependency
- You can use your own style as long as it doesn't break the main idea

### src/usecase

- Use case is your main business logic
- You should always call your repository methods in this package
- You may use your `src/library` functions directly in this package
- Any changes outside this package should not affect your use cases (except changes for business entity and repository)
- If you need config variables, external clients, or repositories, pass/inject them as a dependency
- You can use your own style as long as it doesn't break the main idea

### infra

- In this place, you put any infrastructure configurations or scripts to help you deploy your project in a server or vm
- It is always **BETTER TO** create folders based on what environment used, example:
  - `infra/dev`, for development
  - `infra/staging`, for staging
  - `infra/prod`, for production
  - etc...

### Make your own

You're free to create your own style to suit your requirements as long as still follow the main architecture concept.
You can create folders like; `migration` to place your database migrations, `tmp` to place temporary files, etc.

## Release

### Changelog

Read at [CHANGELOG.md](https://github.com/dalikewara/ayapingping-go/blob/master/CHANGELOG.md)

### Credits

Copyright &copy; 2021 [Dali Kewara](https://www.dalikewara.com)

### License

[MIT License](https://github.com/dalikewara/ayapingping-go/blob/master/LICENSE)
