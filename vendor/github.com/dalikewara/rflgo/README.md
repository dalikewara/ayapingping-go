# rflgo

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/dalikewara/rflgo)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/dalikewara/rflgo)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/dalikewara/rflgo)
![GitHub license](https://img.shields.io/github/license/dalikewara/rflgo)

**rflgo** composes new value into **destination** type based on the **source** data value. It will reflects all values
from the **source**, but keeps the structure, types and the properties of the **destination**. Both must have the same kind of type,
for example, if `dest int` then the source must be `source int`.

## Getting started

### Installation

You can use the `go get` method:

```bash
go get github.com/dalikewara/rflgo
```

### Todos

- Add support for this kind of type: `map`, `chan`, `array`, `func`

### Usage

Imagine you have **source** `s []*userSource` data like this:

```go
type roleSource struct {
    Permission string
    CreatedAt  time.Time
}
type userSource struct {
    Id        int
    Name      string
    Roles     *[]roleSource
    CreatedAt time.Time
}
r := &[]roleSource{
    {
        Permission: "create",
        CreatedAt:  time.Now(),
    },
}
s := []*userSource{
    {
        Id:        1,
        Name:      "johndoe",
        Roles:     r,
        CreatedAt: time.Now(),
    },
    {
        Id:        2,
        Name:      "dalikewara",
        Roles:     r,
        CreatedAt: time.Now(),
    },
}
```

```json
[{"Id":1,"Name":"johndoe","Roles":[{"Permission":"create","CreatedAt":"2022-09-06T18:14:33.620313918+07:00"}],"CreatedAt":"2022-09-06T18:14:33.620314256+07:00"},{"Id":2,"Name":"dalikewara","Roles":[{"Permission":"create","CreatedAt":"2022-09-06T18:14:33.620313918+07:00"}],"CreatedAt":"2022-09-06T18:14:33.620314303+07:00"}]
```

and you want to compose the **source** data into this `d []*userDest`:

```go
type roleDest struct {
    Permission string
}
type userDest struct {
    Name  string
    Roles *[]roleDest
}
var d []*userDest
```

you can compose it with `rflgo.Compose()`:

```go
err := rflgo.Compose(&d, s)
if err != nil {
    panic(err)
}
```

then the `d` will be the same as shown bellow:

```go
type roleDest struct {
    Permission string
}
type userDest struct {
    Name  string
    Roles *[]roleDest
}
d := []*userDest{
    {
        Name: "johndoe",
        Roles: &[]roleDest{
            {
                Permission: "create",
            },
        },
    },
    {
        Name: "dalikewara",
        Roles: &[]roleDest{
            {
                Permission: "create",
            },
        },
    },
}
```

```json
[{"Name":"johndoe","Roles":[{"Permission":"create"}]},{"Name":"dalikewara","Roles":[{"Permission":"create"}]}]
```

## Release

### Changelog

Read at [CHANGELOG.md](https://github.com/dalikewara/rflgo/blob/master/CHANGELOG.md)

### Credits

Copyright &copy; 2021 [Dali Kewara](https://www.dalikewara.com)

### License

[GNU General Public License v3](https://github.com/dalikewara/rflgo/blob/master/LICENSE)