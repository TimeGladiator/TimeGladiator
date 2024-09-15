# TimeGladiator

TimeGladiator is a tool designed for self-time management, particularly useful for freelancers. It helps you track, manage, and optimize your time effectively.

## Features

- Track time spent on various tasks
- Generate reports for better time management
- Integrate with other tools and services
- Build and deploy easily using Docker

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.22 or later)
- [Docker](https://docs.docker.com/get-docker/)
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/TimeGladiator.git
    cd TimeGladiator
    ```

2. Build the project:
    ```sh
    make build
    ```

3. Run the project:
    ```sh
    ./bin/main
    ```

### Makefile Commands

The `Makefile` provides several useful commands to help you manage the project. Below are some of the key commands:

- **Help**: Display help information about make commands
    ```sh
    make help
    ```

- **Version**: Display the version of the API server
    ```sh
    make version
    ```

- **All**: Execute all the targets (lint, format, test, build, build-docker, build-packto)
    ```sh
    make all
    ```

- **Build Docker**: Build the API server as a Docker image
    ```sh
    make build-docker
    ```

- **Push Docker**: Push the Docker image to a repository
    ```sh
    make push-docker
    ```

- **Run Docker**: Run the Docker image
    ```sh
    make run-docker
    ```

- **Stop Docker**: Stop the running Docker image
    ```sh
    make stop-docker
    ```

- **Clean Docker**: Remove the Docker image
    ```sh
    make clean-docker
    ```

- **Build Packto**: Build the API server using Paketo buildpacks
    ```sh
    make build-packto
    ```

- **Run Packto**: Run the Docker image built with Paketo
    ```sh
    make run-packto
    ```

- **Lint**: Run golint on all Go packages
    ```sh
    make lint
    ```

- **Format**: Run gofmt on all Go packages
    ```sh
    make fmt
    ```

- **Run**: Run the API server
    ```sh
    make run
    ```

- **Build**: Build the API server
    ```sh
    make build
    ```

- **Test**: Run the tests
    ```sh
    make test
    ```

- **Reset Dependencies**: Reset the Go modules
    ```sh
    make deps-reset
    ```

- **Upgrade Dependencies**: Upgrade the Go modules
    ```sh
    make deps-upgrade
    ```

- **Tidy Dependencies**: Tidy up the Go modules
    ```sh
    make tidy
    ```

- **Clean Dependencies Cache**: Clean the Go modules cache
    ```sh
    make deps-cleancache
    ```

### Usage

After building the project, you can start using TimeGladiator to manage your time effectively. The tool provides various features to help you track and optimize your time.

### Contributing

We welcome contributions to improve TimeGladiator. Please fork the repository and submit pull requests.

### License

This project is licensed under the MIT License - see the LICENSE file for details.

---

Feel free to customize this README to better fit your project's needs.