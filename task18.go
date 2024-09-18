package main

import (
	"fmt"
	"sync"
)

// Давайте сделаем вид, что мы в другом пакете и поля не видно...
type Counter struct {
	val uint64
	mu  sync.Mutex
}

func (cnt *Counter) Increment() {
	cnt.mu.Lock()
	cnt.val++
	cnt.mu.Unlock()
}

func (cnt *Counter) Get() uint64 {
	return cnt.val
}

func main() {
	var (
		cnt Counter
		wg  sync.WaitGroup
	)
	numOfWorkers := 5

	wg.Add(numOfWorkers)
	for i := range numOfWorkers {
		go func(i int, cnt *Counter, wg *sync.WaitGroup) {
			for _ = range i {
				cnt.Increment()
			}
			wg.Done()
		}(i, &cnt, &wg)
	}

	wg.Wait()
	fmt.Printf("The resulting value of the counter is: %d\n", cnt.Get())
}
