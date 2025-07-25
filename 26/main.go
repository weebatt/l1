package main

import "fmt"

func main() {
	unique := "∂åœ∑´"
	set := make(map[rune]bool)

	fmt.Println(uniqueLetters(unique, &set))
}

func uniqueLetters(unique string, set *map[rune]bool) bool {
	ok := true
	for _, v := range unique {
		if (*set)[v] || (*set)[v+32] {
			ok = false
			return ok
		}

		if v >= 'a' && v <= 'z' {
			(*set)[v] = true
		} else if v >= 'A' && v <= 'Z' {
			(*set)[v+32] = true
		} else {
			(*set)[v] = true
		}
	}
	return ok
}
