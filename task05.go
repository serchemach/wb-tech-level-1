package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var timeToLive int
	flag.IntVar(&timeToLive, "n", 10, "The lifetime of the program in seconds")
	flag.Parse()

	var isEnd atomic.Bool
	go func() {
		time.Sleep(time.Duration(timeToLive) * time.Second)
		isEnd.Store(true)
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	valChannel := make(chan int)
	go func() {
		for {
			val, isOpen := <-valChannel
			if !isOpen {
				fmt.Println("The end of consumer")
				wg.Done()
				return
			}

			fmt.Printf("Recieved value:%d\n", val)
		}
	}()

	for !isEnd.Load() {
		curVal := rand.Intn(2000)
		valChannel <- curVal
	}

	close(valChannel)
	wg.Wait()
	fmt.Println("The end of main")
}
