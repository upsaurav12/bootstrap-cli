#  BootstrapCLI

**BootstrapCLI** is a CLI tool that scaffolds production-ready Golang projects — no dependency headaches, no manual setup.  
Just run a command and get a fully configured project with linters, routers, and structure ready to code.

> Note: This is my first OSS project, I want to make a CLI tool(maybe webUI) which is not just generator tool which only generate go code, but it will help developers to follow best practices, and assist during the project development. In future versions of the project i will add AI which will help developer to assist during their development and help in debugging + fixing error. I am adding AI not to generate code in there project but for assisting purpose only.

* * *
##  Installation

Install globally using `go install`:

`go install github.com/upsaurav12/bootstrap@latest`

Once installed, confirm the installation:

`bootstrap --help`

* * *

## Quick Start 💨

#### 1. Create a New Project

```
bootstrap new myapp --type=rest --router=gin --port=8080 --db=postgres
```
- This command scaffolds a production-ready Go project with:
- Standard project structure
- Database configuration
- Router setup
- Makefile and tooling

#### 2. Prepare the Project

```
cd myapp && make tidy
```

#### 3. Start Required Services (Database)
```
docker compose up -d
```

Before running ``make run`` make sure that you have running 'db' in docker, running with the same credentials as in .env file.

#### 4. Run the Application

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

* * *

##  Roadmap

*    Add CLI command that let users to write their project description, to generate the project automatically without using flags.
*    Command such as ``explain``, ``error`` , ``upgrade`` for the tool to make it progressive CLI tool.
*    Add support for ``auth``, ``logging`` , ``observability`` and so on if it make sense.
*    Add functionality in which users can switch to other options, for example postgres -> mongodb.    

* * *

## Contributing

Contributions, feedback, and ideas are welcome.

If you would like to contribute:

- Open an issue to report bugs or suggest features

- Submit a pull request with clear context and rationale

Please ensure that:

- Code changes are well-structured and documented

- New features include appropriate tests where applicable

- Existing functionality is not broken

You can get started by visiting the project repository on GitHub:
https://github.com/upsaurav12/bootstrap

If you find this project useful, consider giving it a star ⭐ — it helps with visibility and continued development.

* * *

##  License

This project is licensed under the MIT License.

© 2026 [Saurav Upadhyay](https://github.com/upsaurav12)

See the LICENSE file for full license details.
* * *
