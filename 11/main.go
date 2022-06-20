package main

import "fmt"

// содержит ли слайс строк данную величину

func contains(input []string, value string) bool {
	for _, item := range input {
		if item == value {
			return true
		}
	}
	return false
}

func main() {
	firstSet := []string{"a", "b", "c", "d", "e", "f"}
	secondSet := []string{"a", "d", "b", "j", "k", "o"}
	var output []string
	for i := range firstSet {
		if contains(secondSet, firstSet[i]) {
			// если содержит, то добавляем в выходной массив
			output = append(output, firstSet[i])
		}
	}
	fmt.Println(output)
}
