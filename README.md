# ayapingping-go

> **ayapingping-go** has been rebranded to **[uwais](https://github.com/dalikewara/uwais)**. This repository will stay up forever and won’t be deleted, so you can still use it if needed. However, there won’t be any more updates here since all future updates will be worked on in the new **[uwais](https://github.com/dalikewara/uwais)** repository.

[![reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/dalikewara/ayapingping-go/v4)
![version](https://img.shields.io/github/go-mod/go-version/dalikewara/ayapingping-go.svg?style=flat)

![build](https://img.shields.io/circleci/project/github/dalikewara/ayapingping-go.svg?style=flat)
![language](https://img.shields.io/github/languages/top/dalikewara/ayapingping-go.svg?style=flat)
![issue](https://img.shields.io/github/issues/dalikewara/ayapingping-go.svg?style=flat)
![last_commit](https://img.shields.io/github/last-commit/dalikewara/ayapingping-go.svg?style=flat)
![github_tag](https://img.shields.io/github/v/tag/dalikewara/ayapingping-go.svg?style=flat)
![github_license](https://img.shields.io/github/license/dalikewara/ayapingping-go.svg?style=flat)

**ayapingping-go** generates standard project structure to build applications in Golang that follow Clean
Architecture and Feature-Driven Design concept.

> Python variant: [ayapingping-py](https://github.com/dalikewara/ayapingping-py)

> TypeScript variant: [ayapingping-ts](https://github.com/dalikewara/ayapingping-ts)

## Requirements

- Golang>=1.19
- Operating systems supporting `/bin/sh` with **POSIX** standards ([WHY?](https://github.com/dalikewara/ayapingping-sh)).
  **Linux** and **macOS** should have no issues here as they support it by default. For **Windows** users, consider using WSL instead

## Getting started

### Installation

You can use the `go install` method:

```bash
go install github.com/dalikewara/ayapingping-go/v4@latest
```

or, you can also use the `go get` method (DEPRECATED since `go1.17`):

```bash
go get github.com/dalikewara/ayapingping-go/v4
```

### Usage

To generate a new project, simply run the `ayapingping-go` command:

```bash
ayapingping-go
```

Then enter your project name and your Go module path, the **ayapingping-go** generator will set up the project for you.

![Alt Text](https://lh3.googleusercontent.com/drive-viewer/AKGpihZVKfRP1YbgPEilKjEypqE84gyuFpsONb8qqVY2qrnZsAkBo68gqR1UioKlq0G2gW_kCZqFVIPYA7kbRJBrRqb-vl3OnA=w840-h939)

### What's next?

Simply start working on your project and make changes.

## Project Structure

To implement the concept of Clean Architecture and ~~Domain-Driven Design~~ Feature-Driven Design, and to keep them understandable, we structure the project like this:

### main.go

- In this file, you initialize dependencies, injections, and anything required to start and run your application
- You can use the command `go run main.go` or `make start` to run your application

### domain

- The **Domain** represents your primary business model or entity
- Define your main object models or properties for your business here, including database models, DTOs (Data Transfer Objects), etc
- Keep this package as straightforward as possible. Avoid including any code that is not directly related to the model itself
- If a **Feature** imports anything from this location, and you want the **Feature** to be accessible through the `importFeature` or `exportFeature` command
  without the risk of missing package errors, **DON'T FORGET** to include them in the `features/yourFeature/dependency.json` file

### features

- A **Feature** encapsulates your main business feature, logic, or service
- Here, you include everything necessary to ensure the proper functioning of the feature
- Please prioritize **Feature-Driven Design**, ensuring that features can be easily adapted and seamlessly integrated and imported into different projects
- If another **Feature** imports anything from this location (the current **Feature**), and you want the current **Feature** to be
  accessible through the `importFeature` or `exportFeature` command without the risk of missing package errors, **DON'T FORGET** to include them in the `dependency.json` file
- The `dependency.json` is **OPTIONAL**, and **ONLY USEFUL WHEN** you use the `importFeature` or `exportFeature` command. It serves to define
  the **Feature** dependencies and avoids possible missing package errors
- A standard **Feature** comprises the following parts: `delivery`, `repositories`, `usecases` and `utility`
  - **delivery**
    - Hosts feature handlers like HTTP handlers, gRPC handlers, cron jobs, or anything serving between the client and your application or feature
    - For config variables, external clients, or use cases, pass or inject them as dependencies
  - **repositories**
    - Handles communication with external data resources like databases, cloud services, or external services
    - Keep your repositories as simple as possible; avoid adding excessive logic
    - If necessary, separate operations into smaller methods
    - Changes outside the `repositories` should not affect them (except changes for business domain/model/entity)
    - For config variables, database frameworks, or external clients, pass or inject them as dependencies
  - **usecases**
    - Contains the main feature logic
    - Changes outside the `usecases` should not impact them (except changes for business domain/model/entity and repositories)
    - For config variables, external clients, or repositories, pass or inject them as dependencies
  - **utility**
    - Accommodates functions tailored to help with common tasks specifically for the **Feature**—treat them as helpers
- Feel free to adopt your own style as long as it aligns with the core concept

### common

- In this place, you can implement various functions to assist you in performing common tasks—consider them as helpers
- Common functions can be directly called from any location
- If a **Domain** or **Feature** imports anything from this location, and you want the **Feature** to be accessible through
  the `importFeature` or `exportFeature` command without the risk of missing package errors, **DON'T FORGET** to include
  them in the `features/yourFeature/dependency.json` file

### infra

- This is the location to house infrastructure configurations or scripts to facilitate the deployment of your project on a server or VM

### Make It Your Own

Feel free to create your own style to suit your requirements, as long as you still follow the main architecture concept.
You can create folders such as `migration` to store your database migrations, `tmp` for temporary files, etc.

## Importing Features from Another Project

To seamlessly incorporate or import features from another project, use the `importFeature` command:

```bash
ayapingping-go importFeature [feature1,feature2,...] from [/local/project or https://example.com/user/project.git or git@example.com:user/project.git]
```

For example:

```bash
ayapingping-go importFeature exampleFeature from /path/to/your/project
```

```bash
ayapingping-go importFeature exampleFeature1,exampleFeature2 from git@github.com:username/project.git
```

### Feature dependency

If your feature relies on external packages, it's crucial to address dependencies properly during the import process.
Failure to import necessary dependencies may result in missing packages. To prevent this, please document your feature
dependencies in the `dependency.json` file. Supported dependencies are limited to the following directories: `domain`, `common`, and `features`.
Ensure that your feature dependencies strictly adhere to these directories, avoiding reliance on other locations.
You can also include any external packages to `externals` param to install them automatically using the `go get` method.

Example `dependency.json` file:

```json
{
  "domains": [
    "domain1.go",
    "domain2.go"
  ],
  "features": [
    "feature1",
    "feature2"
  ],
  "commons": [
    "commonFunction1.go",
    "commonFunction2.go"
  ],
  "externals": [
    "github.com/go-sql-driver/mysql",
    "github.com/jmoiron/sqlx"
  ]
}
```

## Other Commands

There are several commands similar to `importFeature` above, such as `importDomain`, `importCommon`, `exportFeature`, `exportDomain`, etc.
They function in the same way, for example:

```bash
ayapingping-go importDomain example.go from /path/to/your/project
```

```bash
ayapingping-go importCommon commonFunction1.go from https://example.com/user/project.git
```

For `export` command, the behavior is similar to the `import` command, but now uses `export` as the prefix and `to` instead of
`from` when pointing to the source, for example:

```bash
ayapingping-go exportFeature exampleFeature to /path/to/your/project
```

For more detail and explanation, please visit [ayapingping-sh](https://github.com/dalikewara/ayapingping-sh)

## Release

### Changelog

Read at [CHANGELOG.md](https://github.com/dalikewara/ayapingping-go/blob/master/CHANGELOG.md)

### Credits

Copyright &copy; 2021 [Dali Kewara](https://www.dalikewara.com)

### License

[MIT License](https://github.com/dalikewara/ayapingping-go/blob/master/LICENSE)
