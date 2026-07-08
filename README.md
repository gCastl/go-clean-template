# Go Clean Architecture Template

A clean architecture template for Go projects, providing a well-structured foundation for building scalable and maintainable applications.

## Purpose

This repository serves as a starter template for Go projects following clean architecture principles. It provides:

- **Layered architecture** with clear separation of concerns (domain, application, infrastructure, and presentation layers)
- **Dependency injection** for loose coupling and testability
- **Project structure** that scales as your application grows
- **Build automation** with Task runner for consistent development workflow

This template helps you kickstart new Go projects with best practices already in place, reducing boilerplate setup and establishing a solid foundation for long-term maintenance.

## Usage as a Template

### Using gonew (Recommended)

The easiest way to use this template is with gonew, which automatically copies the template and updates the module name for your new project.

**Install gonew:**

```bash
go install golang.org/x/tools/cmd/gonew@latest
```

**Create a new project from this template:**

```bash
gonew github.com/gCastl/go-clean-template/<template> github.com/yourname/yourproject
```

This command will:
- Clone the template to your desired location
- Automatically update the Go module name from `github.com/gCastl/go-clean-template` to `github.com/yourname/yourproject`
- Replace all references to the old module name throughout the codebase

Replace `github.com/yourname/yourproject` with your actual module path.

Available template:
- rest-api

## Installation

### Prerequisites

Before getting started, you need to install Task:

**Install Task:**

Task is a task runner/build tool that helps automate project setup and common development tasks.

Visit [Task Installation Guide](https://github.com/go-task/task/blob/main/docs/installation.md) for platform-specific instructions.

Quick install options:
- **macOS**: `brew install go-task/tap/go-task`
- **Linux**: `sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b ~/.local/bin`
- **Windows**: `choco install task` or visit the GitHub releases

### Setup Project

Once you have the template in place and Task is installed, run the install task to set up the project:

```bash
task install
```

This will:
- Download and install project dependencies
- Set up the development environment
- Prepare the project for development

### Available Tasks

After installation, you can view all available tasks by running:

```bash
task --list
```

Common tasks include building, testing, and running the application.

## Project Structure

The project follows clean architecture principles with distinct layers:

```
.
├── cmd/              # Application entry points
├── internal/
│   ├── domain/       # Business logic and entities
│   ├── application/  # Use cases and application logic
│   ├── infrastructure/ # External service implementations
│   └── presentation/ # API handlers and response formatting
├── pkg/              # Public packages that can be imported by other projects
└── Taskfile.yaml     # Task runner configuration
```

## Getting Started

1. Install Task (see Prerequisites)
2. Run `task install` to set up the project
3. Explore the project structure to understand the architecture
4. Start building your application in the appropriate layers
