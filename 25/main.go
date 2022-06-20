package main

import (
	"fmt"
	"time"
)

// сравниваем разницу между текущим и засеченным ранее, при входе
func Sleep(second int) {
	now := time.Now().Unix()
	for {
		if time.Now().Unix()-now > int64(second) {
			return
		}
	}
}

func main() {
	fmt.Println("start")
	Sleep(2)
	fmt.Println("end")
}
