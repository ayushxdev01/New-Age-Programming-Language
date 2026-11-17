# Go Inventory Management Module

This project implements the Unit 1 Go Language practical assessment.

## Features

- `Item` struct with ID, Name, Quantity, and Price.
- Slice-based inventory storage.
- Add, update, remove, and display operations.
- Input validation and error handling.
- Concurrent inventory summary using a goroutine and channel.
- Unit tests for core inventory functions.

## Run the program

```bash
go run .
```

## Run unit tests

```bash
go test -v
```

## Build the project

```bash
go build
```

The summary reports the number of different items, total quantity, and total inventory value.
