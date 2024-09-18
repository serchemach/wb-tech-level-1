package main

import (
	"cmp"
	"fmt"
)

// Очевидно предполагается что массив отсортирован
// В этой вариации мы ищем самый правый элемент, <= данному
func binarySearch[T cmp.Ordered](arr []T, v T) (left int) {
	left, right := 0, len(arr)-1
	var middle int
	for right > left {
		middle = (left+right)/2 + 1
		if arr[middle] <= v {
			left = middle
		} else {
			right = middle - 1
		}
	}
	return
}

func main() {
	var (
		arr1 = []int{1, 2, 3, 4, 4, 4, 4, 5, 5, 5, 6, 20, 123}
		arr2 = []string{"1", "2", "2", "2", "3", "4", "4", "5"}
	)
	fmt.Printf("Arr1: %#v\n", arr1)
	fmt.Printf("Binary search of 4: %#v\n", binarySearch(arr1, 4))
	fmt.Printf("Binary search of 3: %#v\n", binarySearch(arr1, 3))
	fmt.Printf("Binary search of 5: %#v\n", binarySearch(arr1, 5))

	fmt.Printf("Arr2: %#v\n", arr2)
	fmt.Printf("Binary search of \"3\": %#v\n", binarySearch(arr2, "3"))
	fmt.Printf("Binary search of \"2\": %#v\n", binarySearch(arr2, "2"))

}
