# Overview

This project aims to demonstrate parsing/processing of a Quake 3 Arena game, in an event-based application. Generating some reports 
around the events in the game in the end

## Features
- Loads a Quake 3 Arena Game logs
- Parses the logs
- Triggers events for the interested subscribers
- Generates reports

## Considerations
- This project prioritizes memory usage, so we only load one game at a time into memory.
- This project uses parallelism to process the game events.
- This project has raw mocks, not using external libraries for that given its small project size.
- This project uses a dependency injection strategy. This way it's easier to test/evolve.
- The tests can be extended. Some edge case tests were not added due to the deadline.
- This project consumes an already included file and generates a report. No dynamic interaction was added for this MVP.


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

## Examples
[Here](https://gist.github.com/FabioRodrigues/34effb4931ce73eb7c2bcb8dae9cae3a) is an example of a real output of the application
