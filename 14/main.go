package main

import "fmt"

func GetTypeOf(input interface{}) {
	fmt.Printf("Current var = %T\n", input)
}

func main() {
	a1 := "hello world"

	a2 := 10

	a3 := 1.55

	a4 := true

	a5 := []string{"foo", "bar", "baz"}

	a6 := map[int]string{100: "Ana", 101: "Lisa", 102: "Rob"}

	GetTypeOf(a1)
	GetTypeOf(a2)
	GetTypeOf(a3)
	GetTypeOf(a4)
	GetTypeOf(a5)
	GetTypeOf(a6)
}
