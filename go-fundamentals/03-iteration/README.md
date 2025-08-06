# 03: Iteration

For loops and benchmarking.

## Files

- `iteration_test.go` - Tests and benchmarks for Repeat function
- `iteration.go` - Repeat function implementation
- `go.mod` - Module file

## Run

```bash
go test          # Run tests
go test -bench=. # Run benchmarks
```

## Concepts

- `for` loops: `for i := 0; i < count; i++`
- `range` keyword: `for index, value := range slice`
- Benchmarking with `go test -bench=.`
- Using standard library: `strings.Repeat`
- Performance optimization
- Import statements: `import "strings"`

## Iteration Patterns

```go
// Traditional for loop
for i := 0; i < len(slice); i++ { }

// Range over slice (like Python's for item in list)
for index, value := range slice { }

// Range over slice (value only)
for _, value := range slice { }

// Range over map (like Python's for key, value in dict)
for key, value := range map { }

// Range over string (character by character)
for index, char := range "hello" { }
```
