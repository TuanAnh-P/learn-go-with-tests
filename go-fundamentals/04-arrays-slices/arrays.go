package main

// Sum takes an array of integers and returns their sum
// Array: [5]int - fixed size, size is part of type
func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumSlice takes a slice of integers and returns their sum
// Slice: []int - dynamic size, more commonly used than arrays
func SumSlice(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
} 