package main

import (
	"fmt"
	"math"
	"sync"
)

// массив входных чисел
var mass = [5]float64{2, 4, 6, 8, 10}

func main() {
	// создаем группу ожидания, которая будет дожидаться выполнения всех горутин
	var mutex sync.Mutex
	wgroup := new(sync.WaitGroup)
	var result float64
	for item := range mass {
		// добавляем еще одну горутину в список ожидания
		wgroup.Add(1)
		go func(i int, wgroup *sync.WaitGroup, mu *sync.Mutex) {
			// после выполнения горутины засчитаем выполнение горутины
			defer wgroup.Done()
			// вводим мьютекс, чтобы избежать гонок при сложении текущего вычисления с результатом
			mu.Lock()
			result += math.Pow(mass[i], 2)
			mu.Unlock()
		}(item, wgroup, &mutex)
	}
	wgroup.Wait()
	fmt.Println(result)
}
