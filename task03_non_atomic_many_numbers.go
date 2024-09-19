package main

import (
	"fmt"
	"sync"
)

func processOneValue(ind int, arr []int32, result *int32) {
	squared := arr[ind] * arr[ind]

	// Без использования атомарных операций или мутехов будет плохо
	// потому что операция инкремента не атомарна
	// Поэтому давайте посмотрим что будет, мне даже интересно

	// С большим количеством чисел результаты действительно варьируются!
	// В моём случае чаще всего 5530, но иногда бывает и 5529
	*result += squared
	fmt.Println(squared)
}

func main() {
	// Если не использовать какой-то механизм синхронизации, то
	// программа завершается сразу
	var (
		wg     sync.WaitGroup
		result int32
	)
	arr := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 1, 1, 1, 1, 1}

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
