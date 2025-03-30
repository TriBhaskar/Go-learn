# Go Basic Examples: Arrays, Slices, Maps, and Loops

This repository contains basic Go examples demonstrating fundamental concepts including arrays, slices, maps, and different types of loops.

## Contents

### Arrays

- Basic array declaration and initialization
- Fixed-size arrays with different initialization methods
- Array memory address demonstration
- Array indexing and slicing

### Slices

- Slice declaration and initialization
- Dynamic size manipulation using append
- Length and capacity demonstration
- Slice creation using make()
- Performance comparison of preallocated vs non-preallocated slices

### Maps

- Map declaration using make()
- Map initialization with literal values
- Key-value operations
- Checking for key existence
- Deleting map entries

### Loops

- For loop
- While loop (using for with condition)
- Range-based loops
  - Iterating over arrays
  - Iterating over maps

## Performance Testing

The repository includes a performance test comparing slice operations:

- Tests slice performance with and without preallocation
- Demonstrates the importance of capacity planning
- Uses Go's testing package for benchmarking

## Running the Code

To run the main examples:

```bash
go run main.go
```

To run the performance tests:

```bash
go test -v
```
