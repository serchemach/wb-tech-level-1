package main

import (
	"fmt"
	"sync"
)

func consume(input chan int, output chan int, wg *sync.WaitGroup) {
	for val := range input {
		output <- val * val
	}
	fmt.Println("Closed worker")
	wg.Done()
}

func printChan(output chan int, wg *sync.WaitGroup) {
	for val := range output {
		fmt.Printf("Value from the conveyer: %d\n", val)
	}
	fmt.Println("Closed printer")
	wg.Done()
}

func main() {
	inputChannel := make(chan int)
	outputChannel := make(chan int)
	inputArr := []int{1, 2, 3, 5, 6, 7, 8, 9, 0, 1, 3, 4, 6, 7}
	var wgWorker sync.WaitGroup
	var wgPrinter sync.WaitGroup

	numWorkers := 10
	numPrinters := 5

	wgWorker.Add(numWorkers)
	for _ = range numWorkers {
		go consume(inputChannel, outputChannel, &wgWorker)
	}

	wgPrinter.Add(numPrinters)
	for _ = range numPrinters {
		go printChan(outputChannel, &wgPrinter)
	}

	for _, val := range inputArr {
		inputChannel <- val
	}
	close(inputChannel)

	wgWorker.Wait()
	close(outputChannel)

	wgPrinter.Wait()
}
