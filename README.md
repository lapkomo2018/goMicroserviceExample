# Go Microservice Example

This project provides a boilerplate for a Go microservice with gRPC and Docker Compose. It includes a basic setup to run a Go server with Rest, gRPC and a MySQL database using Docker Compose.

## Features

- Go microservice with Rest, gRPC
- Docker Compose setup for server and MySQL database
- Basic configuration for server and database connection

## Prerequisites

- Docker
- Docker Compose
- Go (for development)

## Project Structure

```
├── api
│ └── proto
│   └── service.proto
├── cmd
│ └── app
│   └── main.go
├── config
│ ├── config.go
│ └── config.yaml
├── docs
│ ├── docs.go
│ ├── swagger.json
│ └── swagger.yaml
├── internal
│ ├── core
│ │ └── struct.go
│ ├── service
│ │ ├── service.go
│ │ └── struct.go
│ ├── storage
│ │ └── mysql
│ │   ├── storage.go
│ │   └── struct.go
│ └── transport
│   ├── grpc
│   │ ├── v1
│   │ │ └── handler.go
│   │ └── server.go
│   └── rest
│     ├── v1
│     │ ├── handler.go
│     │ └── struct.go
│     └── server.go
├── Makefile
├── Dockerfile
├── docker-compose.yml
├── .env
├── go.mod
├── go.sum
└── README.md
```

### Important Files

#### `api/proto/service.proto`
- **Description**: This file contains the gRPC service definitions. It defines the RPC methods and the messages used in communication between the client and server.

#### `cmd/app/main.go`
- **Description**: The entry point for the application. It sets up and starts the Go microservice, initializing configurations, and launching the server.

#### `config/config.go`
- **Description**: Contains the code for loading and parsing the configuration file (`config.yaml`). It provides functionality to read and manage application settings.

#### `config/config.yaml`
- **Description**: Configuration file for the microservice. It includes settings such as database connection strings and validation rules.

#### `docs`
- **Description**: Contains automatically generated swagger docs.

#### `internal/core/struct.go`
- **Description**: Example of core data structures and types used throughout the application. It may include shared models or utility functions.

#### `internal/service/service.go`
- **Description**: Contains the services. This file using for creating all services in one struct.

#### `internal/service/struct.go`
- **Description**: Defines data structures specific to the service layer. These structures are used within service methods and handlers.

#### `internal/storage/mysql/storage.go`
- **Description**: Contains the implementation for interacting with the MySQL database. It includes functions for database operations and management.

#### `internal/storage/mysql/struct.go`
- **Description**: Defines data structures and models related to MySQL storage. This file may include schema definitions and database-related utilities.

#### `internal/transport/grpc/v1/handler.go`
- **Description**: Handles gRPC requests and responses. It contains the implementation of the gRPC service methods defined in `service.proto`.

#### `internal/transport/grpc/server.go`
- **Description**: Sets up and configures the gRPC server, including registering services and handling incoming gRPC calls.

#### `internal/transport/rest/v1/handler.go`
- **Description**: Handles REST API requests and responses. It defines the HTTP handlers and routes for RESTful communication.

#### `internal/transport/rest/v1/struct.go`
- **Description**: Defines data structures used in REST API handlers, including request and response models.

#### `internal/transport/rest/server.go`
- **Description**: Sets up and configures the REST server, including route registration and middleware using handler.

#### `.env`
- **Description**: Contains environment variables used to configure the application. This file is typically used to store sensitive information and configuration settings required by the application at runtime. It is not recommended to include in version control to keep sensitive information secure.

#### `Makefile`
- **Description**: Contains build automation commands for the project. It simplifies common tasks such as generating documentation and compiling Protocol Buffers (protobuf) files.

#### `Dockerfile`
- **Description**: Contains instructions for building the Docker image of the Go microservice. It specifies the base image, dependencies, and build steps.

#### `docker-compose.yml`
- **Description**: Docker Compose configuration file. It defines and manages the services, including the Go microservice and MySQL database, and their dependencies.

#### `go.mod` and `go.sum`
- **Description**: Go module files. `go.mod` specifies module dependencies and versions, while `go.sum` ensures the integrity of module downloads.

#### `README.md`
- **Description**: Documentation file for the project. It provides an overview, setup instructions, and usage guidelines.

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/lapkomo2018/goMicroserviceExample.git
cd your-repo
```

### 2. Build and Run the Services

To build and start the services, use Docker Compose:

```bash
docker-compose up -d
```

This command will build the Docker image for the Go microservice and start both the microservice and MySQL database containers.

### 3. Verify the Services

Check if the services are running:

```bash
docker-compose ps
```

You should see both `server` and `mysql_db` running.

### 4. Access the Microservice

The Go microservice will be available at:
- `http://localhost:8080` Rest
- `http://localhost:8081` gRPC

## Configuration

The configuration for the microservice is stored in `config.yaml`. Update the configuration file as needed.

### Example `config.yaml`

```yaml
restServer:
  port: 8080
  bodyLimit: 25165824
  allowedOrigins:
    - "*"

grpcServer:
  port: 8081

hash:
  salt: "123123"

validator:
  messageRegexp: "^.{1,30}$"
```

### Example `.env`

```dotenv
DB="root:strongpass@tcp(localhost:3306)/main?parseTime=true"
```

## Development

To develop the microservice, make changes to the Go code and rebuild the Docker image:

```bash
docker-compose build
```

Then restart the services:

```bash
docker-compose up -d
```

## Stopping the Services

To stop and remove the containers, use:

```bash
docker-compose down
```

## Generating Swagger Docs

To automatically generate docs, use:

```bash
swag init --parseDependency --parseInternal -g cmd/app/main.go
```

## Troubleshooting

- **Build context errors**: Ensure that all files in your project have correct permissions and are not being ignored by `.dockerignore`.
- **Connection issues**: Verify that the configuration for the database connection in `.env` is correct.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any questions or issues, your welcome.
