# Smux-Server

This README introduces a Go backend application template created with chi and wire. It follows the repository pattern, offering a clean and decoupled architecture. As the project evolves, I'll continue to make updates and enhancements. Your feedback is greatly appreciated!

## Prerequisite

Before running the project, make sure you have the following:

- [Go](https://golang.org/) installed on your machine.

## How to run the app

1. Clone the repository:

   ```sh
   git clone https://github.com/prcryx/smux-server.git
   ```

2. Navigate to the project directory

   ```sh
   cd sumx-server
   ```

3. Install dependencies:
   ```sh
   go mod download
   ```
4. Install Google's wire for genererating dependency injection codes:
   ```sh
   go install github.com/google/wire/cmd/wire@latest
   ```
5. create .env file:

   ```sh
   SERVER_PORT=8080
   OS_NAME=windows
   ARCH_NAME=amd64
   ```

6. Run the App:
   ```sh
   make run
   ```

## Directory Setup

```bash
│   .env
│   go.mod
│   go.sum
│   Makefile
│
├───build
├───cmd
│   └───app
│           main.go
│
├───config
│       config.go
│
├───di
│   ├───container
│   │       controllers_registry.go
│   │
│   └───wire
│           controllers.di.go
│           datasources.di.go
│           repositories.di.go
│           usecases.di.go
│           wire.go
│           wire_gen.go
│
└───internals
    ├───application
    │   ├───app
    │   │       app.go
    │   │
    │   ├───controllers
    │   │   ├───auth
    │   │   └───blog
    │   │           blog_controller.go
    │   │           blog_router.go
    │   │
    │   ├───routes
    │   │   │   all_routes.go
    │   │   │   root.go
    │   │   │
    │   │   └───middlewares
    │   │           cors.go
    │   │
    │   └───server
    │           server.go
    │
    ├───common
    │   ├───constants
    │   │   │   cors_constants.go
    │   │   │   headers_constants.go
    │   │   │
    │   │   └───routerconst
    │   │           route_constants.go
    │   │
    │   ├───err
    │   │       error_constants.go
    │   │       error_response.go
    │   │
    │   └───utils
    │           json.util.go
    │
    ├───data
    │   ├───datasources
    │   │       test_datasource.go
    │   │
    │   └───repositories_impl
    │           post_repository_impl.go
    │
    ├───db
    │       test.db.go
    │
    └───domain
        ├───entities
        │       author.entity.go
        │       blog.entity.go
        │
        ├───repositories
        │       blog_repository.go
        │
        ├───types
        │       app_route.types.go
        │       server.types.go
        │
        ├───usecases
        │       blog_usecase.go
        │
        └───valobj
                appversion.valobj.go

```
