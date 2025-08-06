# 04: Arrays and Slices

Arrays vs Slices and collection operations.

## Files
- `arrays_test.go` - Tests for Sum functions
- `arrays.go` - Sum function implementations
- `go.mod` - Module file

## Run
```bash
go test          # Run tests
go test -cover   # Run tests with coverage
```

## Concepts
- Arrays: `[5]int` (fixed size, size is part of type)
- Slices: `[]int` (dynamic size, more commonly used)
- `len()` function to get length
- `append()` to add elements to slices
- Test coverage with `go test -cover`

## Key Differences
```go
// Array: fixed size
array := [3]int{1, 2, 3}

// Slice: dynamic size
slice := []int{1, 2, 3}
slice = append(slice, 4)  // Can grow
``` 