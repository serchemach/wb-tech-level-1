package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var numberOfWorkers int
	flag.IntVar(&numberOfWorkers, "n", 10, "The number of workers")
	flag.Parse()

	valChannel := make(chan int)
	var wg sync.WaitGroup

	for i := range numberOfWorkers {
		wg.Add(1)
		go func(workerNum int) {
			for {
				val, isOpen := <-valChannel
				if !isOpen {
					fmt.Printf("CLOSED!!!!! WORKER #%d\n", workerNum)
					wg.Done()
					return
				}
				fmt.Printf("WORKER #%d: %d\n", workerNum, val)
			}
		}(i + 1)
	}

	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)
	isEnd := false

	go func() {
		<-sigChannel
		isEnd = true
	}()

	for !isEnd {
		newVal := rand.Intn(2000)
		valChannel <- newVal
	}

	fmt.Println("We got a sigterm")
	close(valChannel)

	// Wait until all goroutines are closed
	wg.Wait()
	fmt.Println("We ended stuff")
}
