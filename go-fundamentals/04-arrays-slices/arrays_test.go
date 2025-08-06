package main

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSumSlice(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	got := SumSlice(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestArrayVsSlice(t *testing.T) {
	// Array: fixed size, size is part of type
	array := [3]int{1, 2, 3}
	if len(array) != 3 {
		t.Errorf("array length should be 3, got %d", len(array))
	}

	// Slice: dynamic size, more flexible
	slice := []int{1, 2, 3, 4, 5}
	if len(slice) != 5 {
		t.Errorf("slice length should be 5, got %d", len(slice))
	}

	// Slice can be resized
	slice = append(slice, 6)
	if len(slice) != 6 {
		t.Errorf("slice length should be 6 after append, got %d", len(slice))
	}
} 