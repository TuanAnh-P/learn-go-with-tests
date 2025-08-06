package main

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	want := "aaaaa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// Benchmark function: Benchmark + FunctionName
// b.N is the number of iterations Go determines for accurate measurement
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func TestIterationPatterns(t *testing.T) {
	// Demonstrate range over slice
	numbers := []int{1, 2, 3}
	sum := 0
	
	for _, num := range numbers {
		sum += num
	}
	
	if sum != 6 {
		t.Errorf("expected sum 6, got %d", sum)
	}
} 