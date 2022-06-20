package main

import (
	"fmt"
	"math/rand"
	"time"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	//возможен выход за пределы строки, если ее длина меньше, чем 100
	if len(v) <= 100 {
		fmt.Println("выход за пределы массива")
	} else {
		justString = v[:100]
		fmt.Println(justString)
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func createHugeString(i int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, i)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	someFunc()
}
