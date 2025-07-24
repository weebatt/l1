package main

import "fmt"

func main() {
	str2reverse := "snow dog sun cat good"
	reversedString := ""

	mainOffset := 0
	helpOffset := 0
	for i := len(str2reverse) - 1; i >= 0; i-- {
		if i == 0 {
			mainOffset++
			reverseString(&str2reverse, &reversedString, mainOffset, helpOffset)
		} else if str2reverse[i] == ' ' {
			reverseString(&str2reverse, &reversedString, mainOffset, helpOffset)
			mainOffset++
			helpOffset = mainOffset
		} else {
			mainOffset++
		}
	}

	fmt.Println(reversedString)
}

func reverseString(str2reverse, reversedString *string, mainOffset, helpOffset int) {
	firstSymbolOfWord := len(*str2reverse) - mainOffset
	lastSymbolOfWord := len(*str2reverse) - helpOffset

	//fmt.Printf("firstSymbolOfWord = %d, lastSymbolOfWord = %d\n", firstSymbolOfWord, lastSymbolOfWord)
	//fmt.Printf("mo = %d | so = %d\n", mainOffset, helpOffset)
	//fmt.Println(str2reverse[firstSymbolOfWord:lastSymbolOfWord], string(str2reverse[firstSymbolOfWord:lastSymbolOfWord]))

	for j := firstSymbolOfWord; j < lastSymbolOfWord; j++ {
		*reversedString += string((*str2reverse)[j])
	}
	*reversedString += " "
}
