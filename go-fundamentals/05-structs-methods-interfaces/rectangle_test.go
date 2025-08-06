package main

import "testing"

func TestArea(t *testing.T) {
	rectangle := Rectangle{12, 6}
	got := Area(rectangle)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestAreaMethods(t *testing.T) {
	// Table-driven test for multiple shapes
	// This demonstrates polymorphism - different types implementing same interface
	areaTests := []struct {
		name    string
		shape   Shape  // Interface type - can hold Rectangle or Circle
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()  // Polymorphic call - same method, different implementations
			if got != tt.hasArea {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
			}
		})
	}
}

func TestInterfaceImplementation(t *testing.T) {
	// This test shows that both Rectangle and Circle implement Shape
	// Go's implicit interface implementation in action
	var shapes []Shape  // Slice of interface type
	
	// Can add Rectangle to Shape slice - Rectangle implements Shape
	shapes = append(shapes, Rectangle{5, 10})
	
	// Can add Circle to Shape slice - Circle implements Shape  
	shapes = append(shapes, Circle{5})
	
	// Both can be treated as Shape interface - polymorphism
	for _, shape := range shapes {
		area := shape.Area() // Polymorphic call - calls appropriate implementation
		if area <= 0 {
			t.Errorf("Area should be positive, got %f", area)
		}
	}
} 