package main

import "testing"

func TestSquareConcurrency(t *testing.T) {
	t.Run("square one to five", func(t *testing.T) {
		array := [5]int{1, 2, 3, 4, 5}
		got := CompetitiveSquaring(&array)
		want := [5]int{1, 4, 9, 16, 25}

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("square negative", func(t *testing.T) {
		array := [5]int{0, -1, -2, -3, -4}
		got := CompetitiveSquaring(&array)
		want := [5]int{0, 1, 4, 9, 16}

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
