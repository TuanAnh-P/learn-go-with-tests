package main

import "testing"

// Test + FunctionName convention
func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	// %q quotes strings in error messages
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
} 