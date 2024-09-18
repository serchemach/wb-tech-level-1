package main

import (
	"fmt"
	"slices"
)

// Ну вообще, с го 1.21 даже встроенная функция есть для этого, но давайте напишем свою, почему бы и нет
func deleteFromSlice[T any](s []T, i int) []T {
	return append(s[:i], s[i+1:]...)
}

func main() {
	slice := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Printf("Slice before deleting the 0 element:\n%#v\n", slice)
	slice = deleteFromSlice(slice, 0)
	fmt.Printf("Slice after deleting the 0 element:\n%#v\n", slice)

	// Встроенный метод
	fmt.Printf("Slice before deleting the 2 element:\n%#v\n", slice)
	slice = slices.Delete(slice, 2, 3)
	fmt.Printf("Slice after deleting the 2 element:\n%#v\n", slice)
}
