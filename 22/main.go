package main

import (
	"fmt"
	"math/big"
)

// используем библиотеку math/big

func Add(a string, b string) string {
	aInt := new(big.Float)
	aInt.SetString(a)

	bInt := new(big.Float)
	bInt.SetString(b)
	return big.NewFloat(0).Add(aInt, bInt).String()
}

func Subtraction(a string, b string) string {
	aInt := new(big.Float)
	aInt.SetString(a)

	bInt := new(big.Float)
	bInt.SetString(b)
	return big.NewFloat(0).Sub(aInt, bInt).String()
}

func Multiply(a string, b string) string {
	aInt := new(big.Float)
	aInt.SetString(a)

	bInt := new(big.Float)
	bInt.SetString(b)
	return big.NewFloat(0).Mul(aInt, bInt).String()
}

func Division(a string, b string) string {
	aInt := new(big.Float)
	aInt.SetString(a)

	bInt := new(big.Float)
	bInt.SetString(b)
	return big.NewFloat(0).Quo(aInt, bInt).String()
}

func main() {
	a := "2188824200011112223123"
	b := "31287147291749127371203712"

	fmt.Println("Сложение:", Add(a, b))
	fmt.Println("Вычитание:", Subtraction(a, b))
	fmt.Println("Умножение:", Multiply(a, b))
	fmt.Println("Деление:", Division(b, a))
}
