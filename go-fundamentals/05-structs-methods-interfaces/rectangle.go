package main

import "math"

// Rectangle represents a rectangle with width and height
// Struct: like a class in other languages, but simpler
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle represents a circle with radius
type Circle struct {
	Radius float64
}

// Shape interface defines methods that shapes must implement
// Interface: defines behavior, not implementation
// Go uses "duck typing": if it has Area() method, it implements Shape
// No explicit 'implements' keyword like Java - it's implicit!
type Shape interface {
	Area() float64
}

// Area calculates the area of a rectangle (function)
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

// Area method for Rectangle (receiver syntax)
// Method: function with receiver (r Rectangle)
// This automatically makes Rectangle implement Shape interface
// Because Rectangle has Area() float64 method, it "implements" Shape
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area method for Circle
// This automatically makes Circle implement Shape interface
// Because Circle has Area() float64 method, it "implements" Shape
// Both Rectangle and Circle can be used wherever Shape is expected
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
} 