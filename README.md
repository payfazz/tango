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

If you need to overwrite all previous code you can use the --force, --f flags. The CLI will backup all your code inside domain directory and generate new stubs in the domain directory.

```
$ tango make --force
```

## Features

#### Supported DB

- Postgres

#### Implemented

- Endpoint throttling
- Pre push git hooks
- Migration using fazzdb
- Seeder using fazzdb
- Query logging using fazzdb
- Routing using fazzrouter + go-router
- Middleware using go-middleware
- Cors middleware
- Environment flagging
- Code generator (CRUD with service, command, query and repository)

Deployment:
- Dockerfile
- Jenkinsfile
- Test scripts

#### Planned

- CLI
- Unit test per service
