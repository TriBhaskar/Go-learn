# Go Basic Functions and Control Statements

This program demonstrates fundamental concepts in Go programming including functions, error handling, and control flow statements.

## Features

### Functions

1. `printMe(string)` - Prints a greeting message with the provided name
2. `intDivision(int, int)` - Performs integer division and returns quotient, remainder, and potential error

### Control Statements

- **If-else** statements for error handling and conditional logic
- **Switch** statement for handling different remainder cases

## Code Examples

### Function with Multiple Return Values

```go
func intDivision(numerator int, denominator int) (int, int, error)
```

- Returns: quotient, remainder, and error
- Includes error handling for division by zero

### Control Flow

The program demonstrates:

- Error checking using if-else statements
- Switch statement with multiple cases
- Default case handling

## Sample Output

When running with inputs (10, 3):

```
Hello, John Doe here!
remainder of intDivision 1
remainder is 1 or 2
```
