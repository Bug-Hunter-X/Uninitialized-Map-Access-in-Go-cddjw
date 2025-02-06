# Uninitialized Map Access in Go

This repository demonstrates a common error in Go: accessing elements of an uninitialized map.  Attempting to read or write to a map before initializing it with `make()` will result in a runtime panic.

## The Bug

The `bug.go` file contains a simple Go program that attempts to assign a value to a map without first initializing it.  This will cause the program to panic with a runtime error.

## The Solution

The `bugSolution.go` file shows the correct way to handle map initialization, avoiding the panic.  The map is created using `make()` before any attempts to access or modify its elements.

## How to reproduce

1. Clone this repository.
2. Navigate to the repository directory.
3. Run `go run bug.go` (This will result in a panic).
4. Run `go run bugSolution.go` (This will execute without errors).