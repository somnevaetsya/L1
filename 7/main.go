package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Counter struct {
	sync.RWMutex
	m map[int]int
}

// мапа - непотокобезопасна, поэтому обложили мьютексом

func Add(c *Counter, key int, value int) {
	c.Lock()
	c.m[key] = value
	c.Unlock()
}

func main() {
	var counter = Counter{m: map[int]int{}}

	go func() {
		for i := 0; i < 10; i++ {
			Add(&counter, i, rand.Intn(10))
		}
	}()

	time.Sleep(time.Second)

	go func() {
		for i := range counter.m {
			counter.Lock()
			fmt.Println(counter.m[i])
			counter.Unlock()
		}
	}()

	time.Sleep(time.Second)
}
