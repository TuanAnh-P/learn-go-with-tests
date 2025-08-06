# 05: Structs, Methods & Interfaces

Object-oriented concepts in Go.

## Files
- `rectangle_test.go` - Tests for shapes and methods
- `rectangle.go` - Structs, methods, and interfaces
- `go.mod` - Module file

## Run
```bash
go test
```

## Concepts
- **Structs**: `type Rectangle struct { Width, Height float64 }`
- **Methods**: `func (r Rectangle) Area() float64`
- **Interfaces**: `type Shape interface { Area() float64 }`
- **Table-driven tests**: Testing multiple scenarios
- **Receiver syntax**: `(r Rectangle)` for methods

## Key Differences
```go
// Struct (like class in other languages)
type Rectangle struct {
    Width  float64
    Height float64
}

// Method (function on struct)
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Interface (polymorphism)
type Shape interface {
    Area() float64
}
``` 