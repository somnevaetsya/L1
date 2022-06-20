package main

import (
	"fmt"
	"math"
	"sync"
)

// массив числел

var mass = [5]float64{2, 4, 6, 8, 10}

func main() {
	// создаем группу ожидания, которая будет дожидаться выполнения всех горутин
	wgroup := new(sync.WaitGroup)
	for item := range mass {
		// добавляем еще одну горутину в список ожидания
		wgroup.Add(1)
		go func(i int, wgroup *sync.WaitGroup) {
			// после выполнения горутины засчитаем выполнение горутины
			defer wgroup.Done()
			fmt.Println(math.Pow(mass[i], 2))
		}(item, wgroup)
	}
	wgroup.Wait()
}
