package main

import (
	"fmt"
	"sync"
)

/*
Как я понял в посте https://go.dev/blog/loopvar-preview, как раз говорится, что с версии 1.22 баг
в Go, связанный с замыканиями в горутинах, исправлен... мол область видимости переменных цикла
в горутинах теперь корректно работает, и не нужно использовать отдельную переменную для каждого
итерации цикла, как это было до версии 1.22.

Возможно я что то понимаю не так и корректно будет или передавать индекс как параметр, или объявлять
переменную цикла в теле цикла
*/

func main() {
	array := [5]int{2, 4, 6, 8, 10}
	array = CompetitiveSquaring(&array)
	for i, value := range array {
		fmt.Printf("i_%d: %d\n", i, value)
	}
}

func CompetitiveSquaring(array *[5]int) [5]int {
	var wg sync.WaitGroup
	wg.Add(len(array))

	// Объявление переменной цикла в теле циклаы
	for i := range len(array) {
		i := i
		go func() {
			array[i] *= array[i]
			wg.Done()
		}()
	}

	// Передача индекса как параметра в горутину
	for i := range len(array) {
		go func(index int) {
			array[index] *= array[index]
			wg.Done()
		}(i)
	}

	// for i := range len(array) {
	// 	go func() {
	// 		array[i] *= array[i]
	// 		wg.Done()
	// 	}()
	// }

	wg.Wait()

	return *array
}
