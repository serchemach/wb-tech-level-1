package main

import "fmt"

// В принципе можно и считерить, плюс это работает для любых типов
func cheatSwap(v1 *int, v2 *int) {
	*v1, *v2 = *v2, *v1
}

// Либо старый добрый прикол с ксором
func xorSwap(v1 *int, v2 *int) {
	*v1 ^= *v2
	*v2 ^= *v1
	*v1 ^= *v2
}

// На самом деле ещё можно со сложением, но я забыл как

func main() {
	v1, v2 := 123, 321
	fmt.Printf("Numbers before swap: %d, %d\n", v1, v2)
	cheatSwap(&v1, &v2)
	fmt.Printf("Numbers after swap: %d, %d\n", v1, v2)

	fmt.Printf("Numbers before swap: %d, %d\n", v1, v2)
	xorSwap(&v1, &v2)
	fmt.Printf("Numbers after swap: %d, %d\n", v1, v2)
}
