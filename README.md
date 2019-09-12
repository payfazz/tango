# tango

tango is a skeleton for building golang web services

## Documentation

### Installation

Use the `go` command:

```
$ go get github.com/payfazz/tango
```

### CLI Command

#### Init project

Create new project directory, new directory with the same name as <my-project-name> will be created on your current directory

```
$ tango init <my-project-name>
```

#### Generate project structure

Generate required domain file (using CQS principle) and given base repository

Note: template can be modified in `./make/template` directory

```
$ tango make <path_to_structure; default: ./make/structure.yaml>
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
- code generator (CRUD with service, command, query and repository)

Integration:
- authfazz
- contentfazz

Deployment:
- Dockerfile
- Jenkinsfile
- test scripts

#### Planned

- CLI
- unit test per service
- swagger with go-swagger