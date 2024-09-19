package main

import (
	"fmt"
	"sync"
)

func incrementAtomicaly(m map[int]int, mu *sync.Mutex, key int) {
	mu.Lock()
	m[key]++
	mu.Unlock()
}

func main() {
	// Как говорится, concurrent write to map IS NOT OKAY
	// Поэтому будем использовать механизм синхронизации
	// Я решил взять старый добрый Mutex

	m := map[int]int{1: 0, 0: 0, 2: 0, 3: 0, 4: 0}
	var (
		mu sync.Mutex
		wg sync.WaitGroup
	)

	for i := range 5 {
		wg.Add(1)
		go func(i int) {
			for _ = range i {
				incrementAtomicaly(m, &mu, i)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Printf("The resulting map:\n%#v\n", m)
}
