# Go Strings, Runes, and Bytes Demo

This code demonstrates various aspects of string handling in Go, including working with runes, string concatenation, and using the `strings.Builder`.

## Concepts Covered

### 1. Runes in Go

- Converting string to rune slice: `[]rune("Bhèskar")`
- Individual rune access and type demonstration
- Rune literal declaration: `rune = 'B'`
- Unicode character support (demonstrated with 'è')

### 2. String Iteration

- Using `range` to iterate over runes in a string
- Accessing index and value of each character
- Proper handling of Unicode characters

### 3. String Concatenation

- Basic string concatenation using `+=` operator
- More efficient concatenation using `strings.Builder`
- Working with string slices

### 4. Important Notes

- Strings are immutable in Go
- `strings.Builder` provides better performance for multiple string concatenations
- Runes represent Unicode code points in Go

## Code Examples

### Converting String to Runes

```go
var myString = []rune("Bhèskar")
```

### String Concatenation Comparison

```go
// Using += operator
var catStr = ""
for i := range strSlice {
    catStr += strSlice[i]
}

// Using strings.Builder (more efficient)
var strBuilder strings.Builder
for i := range strSlice {
    strBuilder.WriteString(strSlice[i])
}
```

## Output Examples

- Rune values of "Bhèskar"
- Individual rune type demonstration
- Concatenated strings using different methods
