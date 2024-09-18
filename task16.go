package main

import (
	"cmp"
	"fmt"
)

func hoarePartition[T cmp.Ordered](arr []T, pivot T) int {
	i, j := 0, len(arr)-1
	for {
		for ; arr[i] < pivot; i++ {
		}
		for ; arr[j] > pivot; j-- {
		}

		if i >= j {
			return j
		}

		arr[i], arr[j] = arr[j], arr[i]
		if arr[j] == pivot {
			i++
		} else if arr[j] == pivot {
			j--
		}
	}
}

func quickSort[T cmp.Ordered](arr []T) {
	n := len(arr)
	if n < 2 {
		return
	}

	// Не будем ничего мутить, возьмём за пивот первый элемент
	pivot := arr[0]

	point := hoarePartition(arr, pivot)

	quickSort(arr[:point])
	quickSort(arr[point+1:])
}

func main() {
	var (
		arr1 = []int{1, 5, 12, 324, 214, 2, 4, 124, 2164, 214, 21, 4}
		arr2 = []string{"12312", "absc", "823123", "g", "55555"}
	)

	fmt.Printf("The array before sorting: %#v\n", arr1)
	quickSort(arr1)
	fmt.Printf("The array after sorting: %#v\n", arr1)

	fmt.Printf("The array before sorting: %#v\n", arr2)
	quickSort(arr2)
	fmt.Printf("The array after sorting: %#v\n", arr2)

}
