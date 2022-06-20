package main

import "fmt"

// создание структуры

type Human struct {
	name string
}

// определение методов структуры Human

func (human *Human) Speak(toSpeak string) {
	fmt.Println(toSpeak)
}

// встраиваем Human в Action

type Action struct {
	nameOfAction string
	Human
}

func main() {
	var action Action
	action.Speak("Hello")
}
