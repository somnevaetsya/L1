package main

import "fmt"

func main() {
	a := 10
	b := 20
	fmt.Println("a:", a, "b:", b)
	b, a = a, b
	fmt.Println("a:", a, "b:", b)
}
