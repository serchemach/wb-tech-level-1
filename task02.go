package main

import (
	"fmt"
	"sync"
)

func processOneValue(ind int, arr []int) {
	squared := arr[ind] * arr[ind]
	fmt.Println(squared)
}

func main() {
	// Если не использовать какой-то механизм синхронизации, то
	// программа завершается сразу
	var wg sync.WaitGroup

	arr := []int{2, 4, 6, 8, 10}
	for ind, _ := range arr {
		wg.Add(1)

		go func() {
			defer wg.Done()
			processOneValue(ind, arr)
		}()
	}

	wg.Wait()
}
