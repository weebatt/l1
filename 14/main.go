package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v interface{}

	// int
	v = 20
	RuntimeTypeChecker(v)

	// string
	v = "Fake ID"
	RuntimeTypeChecker(v)

	// bool
	v = true
	RuntimeTypeChecker(v)

	// chan int
	v = make(chan int)
	RuntimeTypeChecker(v)

	// chan bool
	v = make(chan bool)

	// chan []int
	v = make(chan []int)
	RuntimeTypeChecker(v)

	// chan map
	v = make(chan map[int]int)
	RuntimeTypeChecker(v)
}

func RuntimeTypeChecker(v interface{}) {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int:
		fmt.Println("int")
	case reflect.String:
		fmt.Println("string")
	case reflect.Bool:
		fmt.Println("bool")
	case reflect.Chan:
		fmt.Println("chan")
	default:
		fmt.Printf("unknown %v\n", v)
	}
}
