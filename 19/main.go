package main

import "fmt"

// проходимся по строке, используя два индекса: на начало и на конец. меняем символы местами, преобразовываем к строке

func ReverseString(input string) string {
	runes := []rune(input)
	j := len(runes) - 1
	for i := 0; i < (len(runes)-1)/2; i++ {
		runes[i], runes[j] = runes[j], runes[i]
		j--
	}
	return string(runes)
}

func main() {
	fmt.Println(ReverseString("англосакс"))
	fmt.Println(ReverseString("✓ Hello 世界"))
}
