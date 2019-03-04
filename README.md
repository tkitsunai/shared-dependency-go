## How to Shared Code with dependency

If you want to share code between golang micro services, how do you solve it?

## Project Structure

```
└── shared-dependency
    ├── Makefile
    ├── api1
    │   └── main.go
    ├── api2
    │   └── main.go
    ├── api3
    │   └── main.go
    ├── common
    │   └── config
    │       └── config.go
    ├── go.mod
    └── go.sum
```

`common` directory is writing common code.

## go.mod

Prepare go.mod in the Project Root and define the module name as project root name.

```
module shared-dependency

go 1.12

require (
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/echo/v4 v4.0.0
	github.com/mattn/go-colorable v0.1.1 // indirect
	github.com/mattn/go-isatty v0.0.6 // indirect
	golang.org/x/crypto v0.0.0-20190228161510-8dd112bcdc25 // indirect
)
```

### Caution

- api1 and api2 can depend each other.
  - If you do not want to make it possible, Should be need custom linter.