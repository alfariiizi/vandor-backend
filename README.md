# Go Hexagonal Architecture Template

This project provides a template for building Go applications using hexagonal
architecture (also known as ports and adapters architecture). It offers a
structured starting point that emphasizes separation of concerns, testability,
and maintainability.

## Project Structure

```
.
├── cmd/                  # Command-line interface definitions using Cobra
├── config/               # Application configuration
├── internal/             # Private application code
│   ├── delivery/         # Input adapters (how the app receives requests)
│   │   ├── http/         # HTTP handlers
│   │   └── route/        # Route definitions
│   ├── domain/           # Domain layer (business logic core)
│   │   ├── entity/       # Business entities
│   │   └── model/        # Data models
│   ├── infrastructure/   # Infrastructure concerns
│   │   └── database/     # Database connections and initialization
│   ├── repository/       # Data access layer (output ports)
│   └── service/          # Business logic services (use cases)
├── scripts/              # Utility scripts
└── ...
```

## Tech Stack

This project uses the following core technologies and libraries:

- **Go**: Version 1.24.1
- **Echo**: Fast and minimalist web framework (v4.13.3)
- **Uber FX**: Dependency injection framework for modular applications
- **GORM**: Feature-rich ORM with SQLite driver
- **Cobra**: CLI application framework with command nesting and flag support
- **Validator**: Request validation using go-playground/validator
- **UUID**: Google's UUID implementation
- **Zap**: High-performance logging (via Uber FX)
- **GoDotEnv**: Environment variable management

## Features

- **Hexagonal Architecture**: Clean separation between domain logic, application
  services, and external interfaces
- **CLI Interface**: Command-line interface using Cobra
- **Configuration Management**: Environment-based configuration
- **Database Integration**: SQLite database support (easily extendable to other
  databases)
- **Dependency Injection**: Modular application structure with Uber FX
- **Request Validation**: Structured input validation
- **Docker Support**: Containerization for consistent deployment
- **Development Tools**: Air for hot reloading during development

## Prerequisites

- Go 1.21 or later
- Make
- Docker (optional, for containerization)

## Getting Started

### Environment Setup

Copy the example environment file and adjust as needed:

```bash
cp .env.example .env
```

### Building the Application

```bash
make build
```

This will create a binary in the `bin/` directory.

### Running the Application

To run all services:

```bash
./bin/main serve
```

To run only the HTTP server:

```bash
./bin/main http
```

For development with hot reloading:

```bash
./scripts/watch.sh  # Requires the Air tool
```

For hot reloading with activating all delivery, you can run:

```bash
./scripts/watch.sh serve
```

### Available Commands

```
MyApp is a service with multiple modes

Usage:
  myapp [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  http        Run HTTP server
  serve       Run all services

Flags:
  -h, --help   help for myapp
```

## Development

### Project Structure Explanation

This project follows hexagonal architecture principles:

1. **Domain Layer**: Contains the core business logic and entities

   - Located in `internal/domain/`
   - Independent of external concerns

2. **Application Layer**: Implements use cases that orchestrate domain objects

   - Located in `internal/service/`
   - Defines interfaces (ports) that will be implemented by adapters

3. **Adapter Layer**: Connects the application to external systems
   - Input adapters in `internal/delivery/` handle incoming requests
   - Output adapters in `internal/repository/` handle outgoing requests

### Adding a New Feature

To add a new feature:

1. Define any needed entities in `internal/domain/entity/`
2. Create necessary repositories interfaces in `internal/repository/`
3. Implement the repository in a subdirectory of `internal/repository/`
4. Create service logic in `internal/service/`
5. Add HTTP handlers in `internal/delivery/http/`
6. Update routes in `internal/delivery/route/`

### Testing

Run the test suite:

```bash
make test
```

## Docker

Build the Docker image:

```bash
docker build -t myapp .
```

Run the container:

```bash
docker run -p 8000:8000 --env-file .env myapp serve
```

## License

This project is licensed under the MIT License.

## Contributing

Contributions are welcome and appreciated! Here's how you can contribute to this
project:

### Getting Started

1. **Fork the repository** and clone it locally
2. **Create a new branch** for your feature or bug fix
3. **Set up your development environment** using the instructions above

### Making Changes

1. **Follow the existing code style** and project structure
2. **Write tests** for new features or bug fixes
3. **Update documentation** to reflect any changes
4. **Make meaningful commit messages** that clearly explain your changes

### Submitting Changes

1. **Push your changes** to your forked repository
2. **Create a pull request** against the main branch of this repository
3. **Describe your changes** in detail, including the problem they solve
4. **Wait for review** and be open to feedback and suggestions

### Code of Conduct

- Be respectful and inclusive in your interactions with others
- Focus on constructive feedback and discussions
- Help others when possible and share knowledge openly

### Issue Reporting

If you find a bug or have a feature request:

1. Check if the issue already exists in the project's issue tracker
2. If not, create a new issue with a clear title and detailed description
3. Include steps to reproduce bugs and expected behavior
4. Add relevant screenshots or error messages if applicable
