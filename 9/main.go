package main

import (
	"fmt"
)

func toChannel(mass ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range mass {
			out <- n
		}
		close(out)
	}()
	return out
}

func toPow(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	channel := toChannel(2, 3, 5, 6)
	out := toPow(channel)

	for n := range out {
		fmt.Println(n)
	}
}
