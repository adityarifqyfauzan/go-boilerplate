# Go Boilerplate

I made this repository to make it easier for me to start a new project using Go. There are several common packages that I usually use for development.

## Features

- Web Framework uses Gin Gonic [(Read Documentation)](https://gin-gonic.com/docs)
- ORM uses GORM [(Read Documentation)](https://gorm.io/docs)
- Dependency Injection uses Uber Dig [(Read Documentation)](https://pkg.go.dev/go.uber.org/dig)
- Logging uses Logrus [(Read Documentation)](https://github.com/sirupsen/logrus)

## Requirements

- Go with version [(1.21.0)](https://go.dev/) or greater

## Directory Hierarcy

```
├── app
│   ├── domain
│   │   ├── model
│   │   ├── pb
│   │   │   └── proto
│   │   │       └── *
│   │   └── repository
│   ├── handler
│   ├── middleware
│   ├── request
│   ├── response
│   ├── service
│   └── usecase
├── cmd
│   ├── grpc
│   └── rest
├── config
├── container
├── migrations
├── pkg
│   └── *
├── proto
│   └── *
└── router

```
