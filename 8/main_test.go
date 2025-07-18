package main

import "testing"

func TestInversingSpecificBit(t *testing.T) {
	t.Run("positive number", func(t *testing.T) {
		var n int64 = 5
		i := 0

		got := InversingSpecificBit(i, n)
		var want int64 = 4

		if want != got {
			t.Errorf("want: %d, got: %d", want, got)
		}
	})

	t.Run("negative number", func(t *testing.T) {
		var n int64 = -5
		i := 0

		got := InversingSpecificBit(i, n)
		var want int64 = -6

		if want != got {
			t.Errorf("want: %d, got: %d", want, got)
		}
	})

	t.Run("zero number", func(t *testing.T) {
		var n int64 = 0
		i := 0

		got := InversingSpecificBit(i, n)
		var want int64 = 1

		if want != got {
			t.Errorf("want: %d, got: %d", want, got)
		}
	})
}
