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
	*result += squared
	fmt.Println(squared)
	// В результате моих тестов на этом небольшом наборе данных проблем не было, но вообще должно быть плохо на большом (в теории)
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
