## Go Graphql api using gqlgen and gorm

### Prerequisite

Go

### Installations

```
go get github.com/99designs/gqlgen
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```

## update schema

```
go run github.com/99designs/gqlgen generate
```

## Run server

```
go run main.go
```
