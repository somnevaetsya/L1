package main

import (
	"fmt"
	"strings"
)

func contains(input []string, value string) bool {
	for _, item := range input {
		if strings.ToLower(item) == strings.ToLower(value) {
			return true
		}
	}
	return false
}

func main() {
	input := "abCdefA"
	var output []string
	var result = true
	for i := range input {
		// проверяем, есть ли уже в наборе текущий символ, если есть, то меняем флаг и выходим
		if contains(output, string(input[i])) {
			result = false
			break
		} else {
			// если нет, то записываем
			output = append(output, string(input[i]))
		}
	}
	if result == true {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
