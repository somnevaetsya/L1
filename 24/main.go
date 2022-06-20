package main

import (
	"fmt"
	"l1/24/point"
	"math"
)

func GetDistance(first, second point.Point) float64 {
	return math.Sqrt(math.Pow(float64(first.GetX()-second.GetX()), 2) + math.Pow(float64(first.GetY()-second.GetY()), 2))
}

func main() {
	firstPoint := new(point.Point)
	firstPoint.SetX(1)
	firstPoint.SetY(25)
	secondPoint := new(point.Point)
	secondPoint.SetX(2)
	secondPoint.SetY(5)
	fmt.Println(GetDistance(*firstPoint, *secondPoint))
}
