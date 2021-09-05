# ayapingping-go

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/dalikewara/ayapingping-go)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/dalikewara/ayapingping-go)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/dalikewara/ayapingping-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/dalikewara/ayapingping-go)](https://goreportcard.com/report/github.com/dalikewara/ayapingping-go)
![GitHub license](https://img.shields.io/github/license/dalikewara/ayapingping-go)


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

![Alt Text](https://lh3.googleusercontent.com/pw/AM-JKLXHIY-P9tKx2cI0sgdLTxzvK5ErAwkToS-3to790cY4UDg2yullDtehGV2LEtYEDU-a1-xa9t_0vjTJJVri45aDNXN7BLxx-eAxOflZltzzrwF2bILJ9bHQWsCnXtCNDC8tMWZMk4tPtDP1iu9OYmD4=w600-h372-no)

### What's next?

Just start working on your project, make changes. If you want to run the example app,
you can use `go run` method:

```bash
go run src/apps/api/main.go
```

or you can use `make start` command:

```bash
make start
```


## Project structure

To implement the concept of Clean Architecture and Domain-Driven Design, and to keep them understandable,
we try to structure the project in such a way.

### src

Main source code of your project.

### src/apps

Application implementation of your project.

- Here you place executable of your applications, what framework used, presenters,
API services, or anything between client and business domain.
  
- In this folder, you also set up dependency injection
for business domain. Example:
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
  
- You **SHOULD NOT** use any of the functions directly in the business domain.

- You can initialize them in `src/apps` or `src/apps/your-app` as 
a dependency injection.

### src/databases

Database adapters used by your project.

- Here you place functions to connect to specified database or anything about database connections
used by your project.

- You **SHOULD NOT** use any of the functions directly in the business domain.

- You can initialize them in `src/apps` or `src/apps/your-app` as
  a dependency injection.

### src/domains

Main business domain of your project.

- Here you place main business models or entities, business repositories,
  business services, business use cases, or anything about business flow requirement.
  
- Try to keep business domain **ISOLATED FROM** anything outside `src/domains`.

- Try to keep business domain **NOT DEPENDS** on any application frameworks.

- Any changes outside `src/domains` **SHOULD NOT** affects existing business domain flow.

### src/helpers

Helpers to help you do some common tasks.

- You **SHOULD NOT** use any of the functions directly in the business domain.

- You can use them in any places except `src/domains`.

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