# Go Structs and Interfaces Example

This project demonstrates fundamental concepts of Go programming including structs, interfaces, and methods.

## Features

- Custom types using structs (`gasEngine`, `electricEngine`, `owner`)
- Interface implementation (`engine` interface)
- Method receivers
- Anonymous structs
- Embedded structs

## Code Structure

### Types

- `gasEngine`: A struct representing a gas-powered engine with:

  - `mpg` (miles per gallon)
  - `gallons` (fuel capacity)
  - Embedded `owner` struct

- `electricEngine`: A struct representing an electric engine with:

  - `mpkwh` (miles per kilowatt-hour)
  - `kwh` (battery capacity)

- `engine`: An interface that requires implementation of:
  - `milesLeft() uint8`

### Key Functions

- `milesLeft()`: Calculates remaining miles for both engine types
- `canMakeIt()`: Determines if the vehicle can travel a specified distance

## Usage Example

```go
// Create a gas engine
myEngine := gasEngine{mpg: 25, gallons: 10, owner: owner{name: "John Doe"}}

// Create an electric engine
myElectricEngine := electricEngine{mpkwh: 4, kwh: 10}

// Check travel possibility
canMakeIt(myEngine, 200)
canMakeIt(myElectricEngine, 10)
```

## Learning Points

1. Interface implementation in Go is implicit
2. Struct embedding for composition
3. Method receivers with different types
4. Anonymous struct creation and usage
5. Type safety with interfaces
