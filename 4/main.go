package main

import (
	"fmt"
	"github.com/google/uuid"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// создаем структуру, в которой будут содеражться каналы
// signal - канал, который будет принимать ctrl + c
// data - канал для передачи данных
// stop - канал для отправки сигнала о прекращении работы

type Channels struct {
	signal chan os.Signal
	data   chan string
	stop   chan int
}

func Run() {
	poolSize := 10
	// создание структуры
	c := Channels{signal: make(chan os.Signal, 1), data: make(chan string), stop: make(chan int)}
	// ловим каналом сигнал о прерывании - ctrl + c
	signal.Notify(c.signal, os.Interrupt)

	// ловим сигнал отмены
	go catchSignal(&c)
	// запускаем наполнение и вызов воркеров
	go sendData(&c)
	go startWorkers(poolSize, &c)

	time.Sleep(time.Second * 5)
}

// ловим сигнал остановки, отправляем в стоп, чтобы в другой функции прекратить запись и выйти из цикла
func catchSignal(c *Channels) {
	for {
		switch <-c.signal {
		case syscall.SIGINT:
			fmt.Printf("Получили CTRL + C\n")
			c.stop <- 0
			close(c.stop)
			close(c.signal)
			return
		}
	}
}

//Создаем воркеры в заданном размере
func startWorkers(workerPoolSize int, c *Channels) {
	// группа ожидания, чтобы дождаться, что все горутины выполнились
	wg := new(sync.WaitGroup)
	for i := 0; i < workerPoolSize; i++ {
		wg.Add(1)
		go work(c.data, wg)
	}
	wg.Wait()
	fmt.Println("Заверешние")
}

//work запускает работу воркера до закрытия канала с данными
func work(in chan string, wg *sync.WaitGroup) {
	// по окончании работы функции говорим, что данная горутина отработала
	defer wg.Done()
	for item := range in {
		fmt.Println("Новые данные:", item)
		time.Sleep(time.Millisecond * 100)
	}
}

//sendData - функция, в которой постоянно отсылаем данные
//при получении значения из stopChan завершаем отправку данных и выходим из цикла
func sendData(c *Channels) {
	for {
		select {
		// если у нас есть, что считать из стопа, то надо останавливаться
		case <-c.stop:
			fmt.Println("Останавливаемся")
			// закрываем канал
			close(c.data)
			return
		default:
			// генерим uuid и пишем в канал
			c.data <- uuid.NewString()
			time.Sleep(time.Millisecond)
		}
	}
}

func main() {
	Run()
}
