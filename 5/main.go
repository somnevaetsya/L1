package main

import (
	"fmt"
	"math/rand"
	"time"
)

// пишем в канал

func writeToChan(in chan int) {
	in <- int(rand.Uint32())
}

// считываем из канала

func readFromChan(in chan int) {
	fmt.Println(<-in)
}

func main() {
	channel := make(chan int)
	second := 5
	now := time.Now()
	for {
		go writeToChan(channel)
		go readFromChan(channel)
		// когда выйдет время, закроем канал
		if time.Since(now) > time.Duration(second)*time.Second {
			close(channel)
			return
		}
	}
}
