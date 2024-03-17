# Changelogs

## 2024

- **v4.4.0 - v4.4.2** (2024-03-17)
  - Fix `main_v4.sh` permission
  - Fix create `main_v4_latest.sh` permission
  - Rename `_baseStructure` to `_base_structure`
  - Now using [ayapingping-sh](https://github.com/dalikewara/ayapingping-sh) main shell script
  - Add command `exportFeature`, `exportDomain` and `exportCommon`
  - Add support for param `externals` in the `features/yourFeature/dependency.json` file
  - Change some `_base_structure` scripts

- **v4.2.0 - v4.3.9** (2024-03-11)
  - Add command `importDomain` and `importCommon`
  - Update & fix some scripts

- **v4.1.0 - v4.2.0** (2024-03-09)
  - Rename `structure` directory to `_baseStructure`
  - Update some description in `README.md`
  - Rename `features/{featureName}/commons` directory to `features/{featureName}/utilities` to avoid confusion

- **v4.0.0 - v4.0.3** (2024-03-02)
  - Major project structure changes
  - Transition to Feature-Driven Design
  - Introduction of `importFeature` command

- **v3.0.2** (2024-01-26)
  - Refactoring of certain structures

- **v3.0.1** (2024-01-07)
  - Correction and adjustment of rules to simplify potential implementations

## 2023

- **v3.0.0** (2023-09-17)
  - Significant modifications to the project structure
  - Adjustment of rules to simplify potential implementations

## 2022

- **v2.1.0** (2022-09-14)
  - Refactoring code implementations to `panic("implement me")` for simplicity

- **v2.0.2** (2022-09-12)
  - License change to MIT License

- **v2.0.0 - v2.0.1** (2022-09-11)
  - Major changes to the project structure

- **v1.3.0 - v1.3.1**
  - Script updates
  - Replacement of `domain/example` with `domain/user`
  - Replacement of `infra/dev` with `infra/local`

- **v1.2.0 - v1.2.2**
  - Project structure updates
    - Move `HTTP/API/App/Route` handler inside the domain directory
    - Rename `helpers` directory to `libraries`
    - Addition of `infra` directory
    - Removal of all empty folders
  - Updates to `README.md`
  - Code updates and refactoring

## 2021

- **v1.1.4 - v1.1.6**
  - Update of `go` version to `1.17`
  - Code updates

- **v1.1.0 - v1.1.3**
  - Project structure changes
  - Code updates

- **v1.0.0 - v1.1.0**
  - Initial release
