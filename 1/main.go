package main

import "fmt"

type Action struct {
	Human
	name     string
	skill    string
	duration string
}

type BasicHumanAbilities interface {
	Walk() string
	Breath() string
}

type Human struct {
	name   string
	age    int
	weight int
	height int
}

func (h *Human) Walk() string {
	return fmt.Sprintf("Human's name %s (age %d weight %d height %d). %s walks right now", h.name, h.age, h.weight, h.height, h.name)
}

func (h *Human) Breath() string {
	return fmt.Sprintf("Human's name %s (age %d weight %d height %d). %s breaths now", h.name, h.age, h.weight, h.height, h.name)
}

func CallBasicHumanAbilities(bha BasicHumanAbilities) string {
	return bha.Walk() + "\n" + bha.Breath()
}

func main() {
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

	fmt.Println(CallBasicHumanAbilities(action))
	fmt.Println(CallBasicHumanAbilities(human))
}
