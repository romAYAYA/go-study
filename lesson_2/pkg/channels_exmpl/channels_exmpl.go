package channelsexmpl

import "fmt"

func ShowExmpl() {
	var ch chan int // переменная типа канал, принимающая int
	ch = make(chan int)

	go func(ch chan<- int) { // канал только на отправку
		ch <- 100
	}(ch)

	val := <-ch // операция чтения или записи блокирует программу
	fmt.Println(val)
}

// OUTPUT 100
