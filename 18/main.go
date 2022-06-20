package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	value int
}

// так как будем увеличивать в конкурентной среде, используем мьютекс для безопасности

func (c *Counter) Inc() {
	c.Lock()
	c.value++
	c.Unlock()
}

func (c *Counter) GetValue() int {
	return c.value
}

func main() {
	counter := &Counter{
		value: 0,
	}
	wg := sync.WaitGroup{}
	for w := 0; w < 10; w++ {
		wg.Add(1)
		go func(c *Counter) {
			defer wg.Done()
			counter.Inc()
		}(counter)
	}
	wg.Wait()
	fmt.Println(counter.GetValue())
}
