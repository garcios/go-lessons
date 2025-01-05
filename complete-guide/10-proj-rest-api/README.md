# Event Booking Application

## About this application
This is a REST-API implementations for Event booking application which includes the following endpoints:

### Events Manager

- GET /events
  - List events
- GET /events/:id
  - Retrieve an event
- POST /events
  - Create a new event
- PUT /events/:id
  - Update an event
- DELETE 
  - Delete an event

### User Manager

- POST /signup
  - Create a new user
- POST /login
  - Login and validate credentials

This application requires some packages which are described below.

## Gin Framewrok
Gin is a web framework written in Go (Golang). It features a martini-like API with much better performance, up to 40 
times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.

https://gin-gonic.com/docs/

### How to install Gin package
```shell
go get -u github.com/gin-gonic/gin
```

## SQLLite
A sqlite3 driver that conforms to the built-in database/sql interface.
https://github.com/mattn/go-sqlite3

### How to install sqllite3 package
```shell
go get -u github.com/mattn/go-sqlite3
```

## Cryptography
Package bcrypt implements Provos and Mazières's bcrypt adaptive hashing algorithm.
https://pkg.go.dev/golang.org/x/crypto/bcrypt

### How to install crypto package
```shell
go get -u golang.org/x/crypto
```

## JWT package
A Golang implementation of JSON Web Tokens.
https://github.com/golang-jwt/jwt

### How to install jwt package 
```shell
go get -u github.com/golang-jwt/jwt/v5
```

## Post-installation

After downloading the packages, run `go mod tidy` to update the dependencies.

```shell
go mod tidy
```