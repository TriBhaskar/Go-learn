# Go Goroutines and Synchronization Example

This project demonstrates the use of goroutines and different synchronization mechanisms in Go.

## Overview

The code showcases:

- Concurrent execution using goroutines
- Performance testing with multiple goroutines
- Database simulation with artificial delays
- Different synchronization methods

## Examples Included

### Performance Test

- Creates 1000 goroutines performing CPU-intensive calculations
- Uses `WaitGroup` to wait for all goroutines to complete
- Measures total execution time

### Simulated DB Operations (Commented Code)

- Simulates database calls with artificial delays
- Demonstrates thread-safe operations on shared resources
- Shows proper synchronization patterns

## Synchronization Methods Used

### sync.Mutex

- Basic mutual exclusion lock
- Used when only one goroutine should access the shared resource at a time
- Example: Protecting the `results` slice in commented code

### sync.RWMutex

- Reader/Writer mutex that allows multiple readers but only one writer
- More efficient than Mutex when reads are more frequent than writes
- Example: Used in `save()` (writer) and `log()` (reader) functions

### sync.WaitGroup

- Used to wait for a collection of goroutines to finish
- Add operations you want to wait for with `Add(1)`
- Signal completion with `Done()`
- Wait for all operations with `Wait()`
