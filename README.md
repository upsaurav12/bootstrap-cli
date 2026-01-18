#  BootstrapCLI

**BootstrapCLI** is a CLI tool that scaffolds production-ready Golang projects — no dependency headaches, no manual setup.  
Just run a command and get a fully configured project with linters, routers, and structure ready to code.

* * *
##  Installation

Install globally using `go install`:

`go install github.com/upsaurav12/bootstrap@latest`

Once installed, confirm the installation:

`bootstrap --help`

* * *

## Quick Start 💨

```
bootstrap new myapp --type=rest --router=gin --port=8080 --db=postgres
```

```
cd myapp && make tidy
```

Make sure that the docker and docker compose is installed in your local machine
```
docker compose up -d
```

Before running ``make run`` make sure that you have running 'db' in docker, running with the same credentials as in .env file.
```
make run
```


## Example Project Structure 

```
 myapp/
    ├── cmd
    │   └── main.go
    ├── docker-compose.yml
    ├── go.mod
    ├── go.sum
    ├── internal
    │   ├── config
    │   │   ├── config.go
    │   │   └── config_test.go
    │   ├── db
    │   │   └── database.go
    │   ├── handler
    │   │   └── user_handler.go
    │   ├── model
    │   │   ├── registory.go
    │   │   └── user_model.go
    │   ├── repository
    │   │   └── user_repo.go
    │   ├── server
    │   │   ├── routes.go
    │   │   └── server.go
    │   └── service
    │       └── user_service.go
    ├── Makefile
    ├── project.yaml
    └── README.md

```

##  CLI Options
### `new`

Creates a new project with the specified configuration options.

### Flags

| Flag         | Description                                     | Example              |
|--------------|-------------------------------------------------|----------------------|
| `--type`     | Type of project (e.g., `rest`)                  | `--type=rest`        |
| `--router`   | Router framework (`gin`, `chi`, `echo`, `fiber`) | `--router=gin`       |
| `--port`     | Application port                                | `--port=8080`        |
| `--db`       | Database integration                            | `--db=postgres`      |
| `--entities` | Add entities                                    | `--entities=user`    |


### `apply`
Create a new project using yaml file configurations
| Flag         | Description                                      | Example             |
|--------------|------------------------------------------------- |----------------------|
| `--yaml`     | unique file name for the yaml file               | `--yaml=project.yaml`|

**BootstrapCLI** automates all that.  
You focus on business logic — it handles the rest.

> Note: This is my first OSS project, I want to make a CLI tool(maybe webUI) which is not just generator tool which only generate
> go code, but it will help developers to follow best practices, and assist during the project development. In future versions of the project i will add AI which will help developer to assist during their development and help in debugging + fixing error. I am adding AI not to generate code in there project but for assisting purpose only.

* * *

##  Roadmap

*    Add CLI command that let users to write their project description, to generate the project automatically without using flags.
*    Command such as ``explain``, ``error`` , ``upgrade`` for the tool to make it progressive CLI tool.
*    Add support for ``auth``, ``logging`` , ``observability`` and so on if it make sense.
*    Add functionality in which users can switch to other options, for example postgres -> mongodb.    

* * *

## Contributing

Contributions, feedback, and ideas are welcome!  
Feel free to open an issue or PR on [GitHub](https://github.com/upsaurav12/bootstrap).

Hope you like this project.

* * *

##  License

Licensed under the **MIT License** © 2026 [Saurav Upadhyay](https://github.com/upsaurav12)

* * *
