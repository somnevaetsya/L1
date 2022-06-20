package main

import "fmt"

//обрезаем массив на два: до позиции и после позиции. соединяем обратно

func DeleteElement(input []int, pos int) []int {
	return append(input[:pos], input[pos+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(DeleteElement(slice, 3))
}
