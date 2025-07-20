package main

import "fmt"

type Set struct {
	m map[any]any
}

func NewSet() *Set {
	return &Set{
		m: make(map[any]any),
	}
}

func (s *Set) Add(x any) {
	s.m[x] = struct{}{}
}

func main() {
	seq := []string{"cat", "cat", "dog", "cat", "tree"}
	//seq := []int{1, 1, 2, 3, 5}

	set := NewSet()

	for _, v := range seq {
		set.Add(v)
	}

	fmt.Println(set)
}
