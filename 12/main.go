package main

import "fmt"

var input = []string{"cat", "cat", "dog", "cat", "tree"}

func contains(input []string, value string) bool {
	for _, item := range input {
		if item == value {
			return true
		}
	}
	return false
}

func main() {
	var output []string
	for i := range input {
		if !contains(output, input[i]) {
			// если уже содержит, то не добавляем
			output = append(output, input[i])
		}
	}
	fmt.Println(output)
}
