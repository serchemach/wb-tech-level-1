package main

import (
	"fmt"
)

func main() {
	// Поскольку горутины не имею хэндлов, остановить их извне без
	// содействия изнутри горутины невозможно.

	// Из-за этого существуют два главных способа остановить выполнение горутины:
	// - Использовать какой-то общий ресурс (будь то канал, либо другая структура данных,
	// 	имеющая атомарность) чтобы дать ей сигнал остановиться и прервать выполнение
	// - При возникновении паники в горутине она остановит выполнение (но чтобы наша программа
	// 	не была остановлена вместе с ней, необходимо сделать recovery)

	// Продемонстрируем их на паре примеров.

	// Остановка через общий ресурс (закрытие канала)
	blockChannel := make(chan struct{})
	sharedChannel := make(chan int)
	go func() {
		for {
			_, isOpen := <-sharedChannel
			if !isOpen {
				fmt.Println("Shared resource goroutine died")
				blockChannel <- struct{}{}
				return
			}
		}
	}()

	sharedChannel <- 1
	sharedChannel <- 1
	close(sharedChannel)

	// Остановка через панику
	go func() {
		defer func() {
			fmt.Println("I (goroutine) died")
			if r := recover(); r != nil {
				fmt.Println("But We recovered")
				blockChannel <- struct{}{}
			}
		}()

		arr := []int{1, 2, 3, 4, 5, 6}
		for i := 0; ; i++ {
			_ = arr[i]
		}

	}()

	<-blockChannel
	<-blockChannel
}
