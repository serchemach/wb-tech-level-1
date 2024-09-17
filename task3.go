package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func processOneValue(ind int, arr []int32, result *int32) {
	squared := arr[ind] * arr[ind]
	fmt.Println(squared)

	// Без использования атомарных операций или мутехов будет плохо
	// потому что операция инкремента не атомарна
	atomic.AddInt32(result, squared)
}

func main() {
	// Если не использовать какой-то механизм синхронизации, то
	// программа завершается сразу
	var (
		wg     sync.WaitGroup
		result int32
	)
	arr := []int32{2, 4, 6, 8, 10}

	for ind, _ := range arr {
		wg.Add(1)

		go func() {
			defer wg.Done()
			processOneValue(ind, arr, &result)
		}()
	}

	wg.Wait()
	fmt.Printf("The result is %d\n", result)
}
