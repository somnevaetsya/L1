package main

import (
	"fmt"
	"sync"
	"time"
)

//ждем либо данных в канале, либо его закрытия. при закрытии канала отловим этом и вернемся из горутины
func first(ch chan int, wg *sync.WaitGroup) {
	for {
		foo, ok := <-ch
		if !ok {
			println("done")
			wg.Done()
			return
		}
		fmt.Println(foo)
	}
}

// ждем изменение флага. при его измении выходим из горутины

func second(context *bool, ch chan int, wg *sync.WaitGroup) {
	for {
		if *context {
			fmt.Println("stop with context")
			wg.Done()
			return
		} else {
			foo := <-ch
			fmt.Println(foo)
		}
	}
}

func main() {
	ch := make(chan int)
	context := false

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go first(ch, wg)
	go second(&context, ch, wg)
	time.Sleep(time.Second)
	context = true
	ch <- 0
	ch <- 1
	ch <- 2
	close(ch)
	wg.Wait()
}
