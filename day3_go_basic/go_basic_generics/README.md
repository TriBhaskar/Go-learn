# Go Generics Example

This project demonstrates the use of generics in Go with practical examples including generic functions and structs.

## Features

### Generic Car Type

- Implements a generic `car` struct that can work with different engine types
- Supports both gas and electric engine variations
- Demonstrates type constraints in struct definitions

### Generic Functions

1. `isEmpty[T any]`

   - Checks if a slice of any type is empty
   - Uses the `any` constraint to accept all types

2. `sumSlice[T int | float32 | float64]`
   - Calculates sum of numeric slices
   - Supports int, float32, and float64 types
   - Uses type constraints to ensure numeric operations

## Example Usage

The main.go file demonstrates:

- Creating and using generic slices
- Working with generic functions for different types
- Instantiating generic car structs with different engine types
- Type-safe operations using generics

## Code Structure

- `gasEngine`: Structure for traditional gas-powered engines
- `electricEngine`: Structure for electric vehicle engines
- `car[T]`: Generic car structure that can use either engine type

## Running the Code

To run the example:

```bash
go run main.go
```
