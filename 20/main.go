package main

import (
	"fmt"
	"strings"
)

// делим текущую строку на слайс строк с помощью разделителя пробел

func ReverseWords(input string) string {
	buf := strings.Split(input, " ")
	var output string
	for i := len(buf) - 1; i >= 0; i-- {
		output += buf[i] + " "
	}
	return output[:len(output)-1]
}

func main() {
	fmt.Println(ReverseWords("snow dog sun 123"))
}
