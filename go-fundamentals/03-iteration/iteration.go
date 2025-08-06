package main

// Import standard library package
import "strings"

// Repeat takes a character and repeats it count times
// Uses strings.Repeat for better performance
func Repeat(character string, count int) string {
	return strings.Repeat(character, count)
}

// Examples of different iteration patterns in Go
func IterationExamples() {
	// 1. Traditional for loop with index
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(numbers); i++ {
		_ = numbers[i] // access by index
	}

	// 2. Range over slice (like Python's for item in list)
	for index, value := range numbers {
		_ = index  // index available
		_ = value  // value available
	}

	// 3. Range over slice (value only)
	for _, value := range numbers {
		_ = value // only value, index ignored
	}

	// 4. Range over map (like Python's for key, value in dict)
	person := map[string]string{"name": "Alice", "age": "30"}
	for key, value := range person {
		_ = key   // key available
		_ = value // value available
	}

	// 5. Range over string (character by character)
	for index, char := range "hello" {
		_ = index // index available
		_ = char  // rune (Unicode character)
	}
} 