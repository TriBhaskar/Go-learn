# Go Pointers and Slices Demo

This project demonstrates basic concepts of pointers and slices in Go.

## Pointers

The code shows several pointer operations:

- Creating a new pointer using `new()`
- Dereferencing pointers using `*`
- Getting memory addresses using `&`
- Modifying values through pointers

Example operations demonstrated:

- Creating a pointer to an `int32`
- Modifying the value the pointer points to
- Pointing to different variables
- Showing how changes through pointers affect the original variable

## Slices

The code illustrates slice behavior:

- Creating slices with literal values
- Slice copying (which creates references to the same underlying array)
- Modifying slices and showing how changes affect both the original and copy

## Array Pointer Example

Includes a `square()` function that demonstrates:

- Passing array pointers as function parameters
- Modifying array elements through pointers
- Returning modified arrays
- Memory address printing to show pointer relationships

## Running the Code

To run this example:

```bash
go run main.go
```

The output will show:

- Pointer value changes
- Memory addresses
- Slice modification effects
- Array transformation results
