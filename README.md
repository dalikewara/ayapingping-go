# ayapingping-go

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/dalikewara/ayapingping-go)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/dalikewara/ayapingping-go)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/dalikewara/ayapingping-go)
![GitHub license](https://img.shields.io/github/license/dalikewara/ayapingping-go)

**ayapingping-go** generates standard project structure for building applications in Golang that follow Clean
Architecture and Domain-Driven Design concept.

## Getting started

### Installation

You can use the `go install` method:

```bash
go install github.com/dalikewara/ayapingping-go@latest
```

or, you can also use the `go get` method (DEPRECATED since `go1.17`):

```bash
go get github.com/dalikewara/ayapingping-go
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

Just start working on your project, make changes. If you want to run the example app, you can use `go run` method:

```bash
go run src/apps/api/ginGonic/ginGonic.go
```

or you can use `make start` command:

```bash
make start
```

## Project structure

To implement the concept of Clean Architecture and Domain-Driven Design, and to keep them understandable, we try to
structure the project in such a way.

### src

Main source code of your project.

### src/apps

Application implementation of your project.

- Here you place executable of your applications, define what framework used, and initialize application handlers from
  domains.

- In this folder, you also set up dependency injection for business domain. Example:
    - Initialize configurations.
    - Initialize database connections for business repositories.
    - Initialize business repositories for business services.
    - Initialize business services for business use cases.
    - etc...

- Application implementation can be one of the following, or you can use all of them:
    - REST API (`src/apps/api`).
    - Cron (`src/apps/cron`).
    - Web (`src/apps/web`).
    - etc... (`src/apps/your-app`).

### src/configs

Configuration setup used by your project.

- Here you place functions to set up configurations or anything about configurations used by your project.

- You may use any of the functions directly in the business domain, but it always **BETTER TO NOT** use it directly in
  there. You can pass the configs as a dependency injection for your business domains.

- You can initialize them in `src/apps` or `src/apps/your-app` as a dependency injection.

### src/databases

Database adapters used by your project.

- Here you place functions to connect to specified database or anything about database connections used by your project.

- You **SHOULD NOT** use any of the functions directly in the business domain.

- You can initialize them in `src/apps` or `src/apps/your-app` as a dependency injection.

### src/domains

Main business domain of your project.

- Here you place main business models or entities, business repositories, business services, business use cases, or
  anything about business flow requirement.

- Any changes outside `src/domains` **SHOULD NOT** affect existing business domain flow.

### src/domains/{my-domain}/delivery

Delivery handlers used by your specified business domains.

- Here you place handlers for your specified business domains, to handle presenters or anything to be done between client and the business domains.
  Example:
    - `src/domains/{my-domain}/delivery/http`, to handle http requests from client.
    - `src/domains/{my-domain}/delivery/grpc`, to handle gRPC requests from client.
    - etc...

### src/libraries

Helpers (custom functions) to help you do some common tasks.

- Here you place custom functions to help you do some common tasks for your applications.

- You may use any of the functions directly in the business domain. But, **BE CAREFUL**, any changes
  in this folder **SHOULD NOT** affect existing business domain flow.

- You may use them in any place.

### infra

Infrastructure configuration of your projects.

- Here you place any infrastructure configurations or scripts to help you deploy your project in a server or vm.
- It is always **BETTER TO** create folders based on what environment used, example:
    - `infra/dev`, for development
    - `infra/rc`, for release candidate
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

[GNU General Public License v3](https://github.com/dalikewara/ayapingping-go/blob/master/LICENSE)