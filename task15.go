package main

import "strings"

// Проблема была в том, что когда мы возвращаем слайс на какой-то большой массив,
// то массив всё ещё полностью находится в памяти, хотя мы больше его не используем
// Поэтому лучше скопировать интересующую его часть

func createHugeString(n int) string {
	return strings.Repeat("1", n)
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
}
