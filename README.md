# tango

tango is a skeleton for building golang web services

## Documentation

### Installation

Use the `go` command:

```
$ go get github.com/payfazz/tango
```

### CLI Command

Create new project directory, new directory with the same name as <my-project-name> will be created on your current directory

```
$ tango init <my-project-name>
```

## Features

#### Supported DB

- Postgres

#### Implemented

- simple access control list (by endpoint)
- throttling
- pre push git hooks
- migration using fazzdb
- seeder using fazzdb
- query logging using fazzdb
- routing using fazzrouter
- middleware using go-middleware
- cors middleware
- environment flagging

Integration:
- authfazz
- contentfazz

Deployment:
- Dockerfile
- Jenkinsfile
- test scripts

#### Planned

- CLI
- code generator (CRUD with service, command, query and repository)
- unit test per service
- swagger with go-swagger