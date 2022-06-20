package main

import "fmt"

var input = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

// вычисляем целочисленный остаток от деления на 10, умножаем на 10 и группируем в мапу

func SpreadValues(input []float64, output *map[int][]float64) {
	for i := range input {
		(*output)[int(input[i]/10)*10] = append((*output)[int(input[i]/10)*10], input[i])
	}
}

func main() {
	output := make(map[int][]float64)
	SpreadValues(input, &output)
	fmt.Println(output)
}
