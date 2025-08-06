package main

import "testing"

func TestAdd(t *testing.T) {
	// Sub-tests with t.Run() for better test organization
	t.Run("adds two positive numbers", func(t *testing.T) {
		got := Add(2, 2)
		want := 4

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("adds negative numbers", func(t *testing.T) {
		got := Add(-1, -2)
		want := -3

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
} 