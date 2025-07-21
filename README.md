# Go-Service

A robust backend service built with Go, implementing Domain-Driven Design (DDD)
and Hexagonal Architecture patterns. This service provides a scalable foundation
for building enterprise-grade applications with clean separation of concerns and
dependency injection.

## Architecture Overview

This project follows the principles of **Domain-Driven Design (DDD)** combined
with **Hexagonal Architecture** (also known as Ports and Adapters pattern). This
architectural approach ensures that your business logic remains independent of
external frameworks and infrastructure concerns.

### Core Architecture Components

**Core Layer** - The heart of your application containing pure business logic:

- **Models**: Domain entities and value objects that represent your business
  concepts
- **Services**: Business logic grouped by domain (e.g., auth, user management)
  with one struct per service group and focused methods
- **Repositories**: Data access interfaces generated using Ent.go for type-safe
  database operations
- **Use Cases**: Application-specific business rules, implemented as
  single-responsibility structs with one method each

**Infrastructure Layer** - External concerns and technical implementations:

- **Database**: PostgreSQL integration with connection pooling and transaction
  management
- **Redis**: Caching and session management capabilities
- **External APIs**: Third-party service integrations

**Delivery Layer** - How your application communicates with the outside world:

- **HTTP**: RESTful API endpoints using Chi router and Huma for API
  documentation
- **Cron**: Scheduled background tasks using gocron
- **CLI**: Command-line interface built with Cobra for administrative tasks

### Dependency Management

The project uses **Uber FX** for dependency injection, which provides a robust
and testable way to manage dependencies throughout your application. This
approach ensures loose coupling between components and makes your code more
maintainable and testable.

## Technology Stack

### Core Framework & Language

- **Go 1.24.1**: Latest Go version with improved performance and language
  features
- **Uber FX**: Dependency injection container for managing application lifecycle
  and dependencies

### Database & ORM

- **PostgreSQL**: Primary database with robust ACID compliance and advanced
  features
- **Ent.go**: Type-safe ORM with code generation for database operations
- **Atlas**: Database schema migration tool for version control of your database
  structure
- **pgx/v5**: High-performance PostgreSQL driver with connection pooling

### Web Framework & API

- **Chi Router**: Lightweight and fast HTTP router with middleware support
- **Huma v2**: OpenAPI 3.0 specification and automatic API documentation
  generation
- **CORS**: Cross-Origin Resource Sharing middleware for web API security

### Authentication & Security

- **JWX v2**: JSON Web Token implementation for secure authentication
- **bcrypt**: Password hashing with configurable cost for security

### Caching & Background Processing

- **Redis**: In-memory data structure store for caching and session management
- **Gocron**: Cron job scheduler for background task processing

### Development & Build Tools

- **Task**: Modern task runner as an alternative to Make
- **Docker**: Containerization for consistent deployment environments
- **Cobra**: CLI framework for building command-line applications

## Project Structure

```plaintext
go-service/
├── bin
├── cmd
│   ├── app
│   ├── service-generator
│   └── usecase-generator
├── config
├── database
│   ├── migrate
│   │   └── migrations
│   └── schema
├── docker
├── internal
│   ├── core
│   │   ├── model
│   │   ├── repository
│   │   ├── service
│   │   │   ├── auth
│   │   │   ├── system
│   │   │   └── user
│   │   └── usecase
│   ├── cron
│   ├── delivery
│   │   ├── cmd
│   │   └── http
│   │       ├── api
│   │       │   └── middleware
│   │       ├── method
│   │       ├── route
│   │       │   └── system
│   │       └── server
│   ├── infrastructure
│   │   ├── database
│   │   └── redis
│   ├── types
│   └── utils
├── scripts
├── seeder
├── storage
│   ├── logs
│   └── public
```

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed on your development
machine:

- **Go 1.24.1 or later**: Download from
  [golang.org](https://golang.org/download/)
- **Docker & Docker Compose**: For running databases and containerized services
- **Task**: Install via `go install github.com/go-task/task/v3/cmd/task@latest`
- **PostgreSQL**: Either locally installed or via Docker
- **Redis**: Either locally installed or via Docker

### Installation & Setup

1. **Clone the repository**:

   ```bash
   git clone https://github.com/alfariiizi/vandor.git
   cd go-service
   ```

2. **Install dependencies**:

   ```bash
   go mod download
   ```

3. **Set up environment variables**: Copy `.env.example` file to `.env` and fill up the env file

4. **Start required services**: If using Docker for databases:

   ```bash
   docker-compose up -d postgres redis
   ```

5. **Generate repository code**:

   ```bash
   task generate:repo
   ```

6. **Run database migrations**:

   ```bash
   task migrate:up
   ```

## Development Workflow

### Running the Application

**Development Mode** (with hot reload):

```bash
task run:dev
```

This command will generate repositories, start file watching, and automatically
restart the server when code changes are detected.

**Production Build**:

```bash
task build
./bin/main
```

### Database Operations

**Check migration status**:

```bash
task migrate:status
```

**Create a new migration**:

```bash
task migrate:diff NAME=add_user_table
```

**Apply pending migrations**:

```bash
task migrate:up
```

### Code Generation

This project includes powerful code generation tools to maintain consistency and
reduce boilerplate:

**Generate repository code** (after modifying Ent schemas):

```bash
task generate:repo
```

**Generate a new use case**:

```bash
task generate:usecase name=CreateUser
```

**Generate a new service** (grouped by domain):

```bash
task generate:service group=auth name=LoginUser
```

### Testing

**Run all tests**:

```bash
task test
```

**Run tests with coverage**:

```bash
go test -cover ./...
```

## Docker Deployment

The project includes Docker support for both the application and database
migrations:

### Building Docker Images

**Build application image**:

```bash
./build.sh app
```

**Build migration image**:

```bash
./build.sh migrate
```

**Build both images**:

```bash
./build.sh all
```

### Docker Images

- **alfariiizi/go-app:latest**: Main application server
- **alfariiizi/go-migrate:latest**: Database migration runner

## API Documentation

Once the application is running, you can access the automatically generated API
documentation at:

- **Swagger UI**: `http://localhost:8080/docs`
- **OpenAPI Spec**: `http://localhost:8080/openapi.json`

The API documentation is automatically generated from your code using Huma v2,
ensuring it stays synchronized with your implementation.

## Architecture Benefits

### Domain-Driven Design (DDD)

- **Clear Business Logic**: Core business rules are isolated and easily testable
- **Ubiquitous Language**: Code reflects the business domain language
- **Bounded Contexts**: Different parts of the system have clear boundaries

### Hexagonal Architecture

- **Technology Independence**: Business logic doesn't depend on frameworks
- **Testability**: Easy to test business logic in isolation
- **Flexibility**: Easy to swap out infrastructure components

### Dependency Injection with Uber FX

- **Loose Coupling**: Components depend on interfaces, not concrete
  implementations
- **Lifecycle Management**: Automatic startup and shutdown of components
- **Testing**: Easy to mock dependencies for unit testing

## Contributing

When contributing to this project, please follow these guidelines:

1. **Code Organization**: Place new code in the appropriate architectural layer
2. **Testing**: Write unit tests for new functionality
3. **Documentation**: Update this README if you add new features or change the
   architecture
4. **Migrations**: Always create migrations for database schema changes
5. **Code Generation**: Use the provided generators for consistency

## Development Scripts

The project includes several helpful scripts in the `scripts/` directory:

- `ent-tools.sh`: Ent.go code generation utilities
- `atlas-tool.sh`: Database migration management
- `watch.sh`: File watching for development mode

## Performance Considerations

- **Connection Pooling**: PostgreSQL connections are pooled for optimal
  performance
- **Redis Caching**: Implement caching strategies for frequently accessed data
- **Structured Logging**: Use structured logging for better observability
- **Graceful Shutdown**: The application handles shutdown signals gracefully

## Security Features

- **JWT Authentication**: Secure token-based authentication
- **Password Hashing**: Bcrypt for secure password storage
- **CORS Configuration**: Proper cross-origin request handling
- **Input Validation**: Request validation using Huma v2

## Monitoring & Observability

Consider integrating the following for production deployments:

- **Structured Logging**: Already configured with appropriate log levels
- **Metrics Collection**: Add Prometheus metrics for monitoring
- **Distributed Tracing**: Consider adding OpenTelemetry for request tracing
- **Health Checks**: Implement health check endpoints for load balancers

This README provides a comprehensive guide to understanding and working with the
Go-Service project. The architecture decisions made here prioritize
maintainability, testability, and scalability while following established
patterns in the Go community.
