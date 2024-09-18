package main

import "fmt"

type Set[T comparable] map[T]struct{}

func (set *Set[T]) Insert(vals []T) {
	for _, val := range vals {
		map[T]struct{}(*set)[val] = struct{}{}
	}
}

func (set *Set[T]) Print() {
	fmt.Printf("{")
	for val, _ := range map[T]struct{}(*set) {
		fmt.Printf(" %v", val)
	}
	fmt.Printf(" }")
}

func main() {
	seq := []string{"cat", "cat", "dog", "cat", "tree"}

	seqSet := make(Set[string])
	seqSet.Insert(seq)
	seqSet.Print()
}
