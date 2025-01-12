# Go Defer and Panic Interaction

This repository demonstrates an uncommon yet important interaction between `defer` statements and `panic` calls in Go.

## Bug Description
The `bug.go` file contains a Go program that uses `defer` to handle potential panics. However, even though a panic occurs, the `defer` function still executes. This behavior might be unexpected to developers unfamiliar with how `defer` works in conjunction with panics.

## Solution
The solution in `bugSolution.go` presents ways to handle panics gracefully. Note that panics are generally indicative of unrecoverable errors, and their handling should be strategic.

## How to Run
1. Clone the repository.
2. Navigate to the repository directory.
3. Run the code using `go run bug.go` and `go run bugSolution.go`.