package main

import "strings"

/*
Issue: Предоставленный фрагмент кода описывает проблему утечек памяти через подстроки
Создается глобальная переменная justString, которой присваивается значение подстроки до 100 элемента, при этом переменная justString
будет ссылаться на тот же блок данных, что и переменная v, это приводит к тому что мы перезапишем заголовок включающий в себя (Pointer,
Length, Capacity) будет ссылаться на тот же массив, но его длинна будет равна размеру среза (например [:100], то есть 100 как дано по
условию). Сама же переменная значение justString после завершения someFunc() и createHugeString() останутся в памяти, а доступа к
оставшейся части исходного массива байтового массива не будет.

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}

*/

/*
Solution: Глобально, чтобы решить эту проблему нужно создать дубликат части среза. Есть несколько подходов представленных ниже Ж)

*/

var justString string

func someFunc() {
	v := createHugeString(1 << 10)

	// First one
	justString = string([]byte(v[:100]))

	// Second one
	justString = (" " + v[:100])[1:]

	// Third one
	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(v[:100])
	justString = sb.String()

	// Fourth one
	strings.Repeat(justString, 10)

	// Fifth one
	strings.Clone(justString)
}

func main() {
	someFunc()
}

func createHugeString(size int) string {
	return strings.Repeat("a", size)
}
