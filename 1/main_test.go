package main

import "testing"

func TestBasicAbilities(t *testing.T) {
	human := &Human{
		name:   "Igor",
		age:    20,
		weight: 60,
		height: 170,
	}

	action := &Action{
		Human:    *human,
		name:     "3D modeling",
		skill:    "junior",
		duration: "less than half a year",
	}

	t.Run("basic test by human", func(t *testing.T) {
		got := CallBasicHumanAbilities(human)
		want := "Human's name Igor (age 20 weight 60 height 170). Igor walks right now\nHuman's name Igor (age 20 weight 60 height 170). Igor breaths now"

		assertCorrect(t, got, want)
	})

	t.Run("basic test by action", func(t *testing.T) {
		got := CallBasicHumanAbilities(action)
		want := "Human's name Igor (age 20 weight 60 height 170). Igor walks right now\nHuman's name Igor (age 20 weight 60 height 170). Igor breaths now"

		assertCorrect(t, got, want)
	})

	t.Run("compare results from action and human", func(t *testing.T) {
		got := CallBasicHumanAbilities(action)
		want := CallBasicHumanAbilities(human)

		assertCorrect(t, got, want)
	})
}

func assertCorrect(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
