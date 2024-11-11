# Overview

The goal of this project is to demonstrate parsing / processing of a Quake 3 Arena game, generating some reports 
around the events in the game

## Features
- Loads a Quake 3 Arena Game logs
- Parse the logs
- Trigger events for the interested subscribers
- Generates reports

## Considerations
- This project prioritizes memory usage, so we only load one game at a time in memory
- This projects uses parallelism for processing the game events
- This project has raw mocks, not using external libraries for that given its small project size
- This projects uses dependency injection strategy. This way it's easier to test/evolve


## Makefile Commands

You can use the following commands defined in the Makefile:

- **Run Tests:**
  `make test`
  Runs all the tests in the project using `go test`.

- **Run Application:**
  `make run`
  Runs the application using `go run`.

- **Check Code:**
  `make check`
  Performs a static analysis of the code using `go vet`.

- **Build Application:**
  `make build`
  Compiles the application using `go build`.

### Examples

To run the tests, execute:
`make test`

To run the application, execute:
`make run`

To check the code, execute:
`make check`

To build the application, execute:
`make build`

## Entry Point

The entry point of the application is located in the `cmd/main` directory. This is where the main function resides and where the application starts execution.

## Parallelism

The project utilizes parallelism for most cases where it makes sense to do so, ensuring better performance and efficiency.

## Getting Started

To get started with this project, clone the repository and navigate to the project directory. Use the provided Makefile commands to test, run, check, or build the application as needed.

```sh
git clone https://github.com/FabioRodrigues/cw-q3arena.git
cd cw-q3arena
make run
```