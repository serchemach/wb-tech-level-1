package main

import (
	"cmp"
	"fmt"
	"sort"
)

// Для начала стоит сказать что в данном задании предполагается
// уникальность элементов в массивах, что хранят множетсва

func sortSlice[T cmp.Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

// Если мы можем упорядочивать элементы множеств (то есть на них определено отношение порядка),
// то довольно эффективным способом нахождения пересечения является их сортировка и
// параллельных обход, что даёт нам O(n + m + nlog(n) + mlog(m))
// Что на самом деле просто O(nlog(n) + mlog(m))
func intersectSetsOrderable[T cmp.Ordered](s1 []T, s2 []T) (intersection []T) {
	// Поскольку множества не упорядоченные, можно не париться об изменении их порядка
	sortSlice(s1)
	sortSlice(s2)
	n, m := len(s1), len(s2)
	// Как обычно больше cap, чтобы не перевыделять память
	intersection = make([]T, 0, min(n, m))
	for i, j := 0, 0; i < n && j < m; {
		switch {
		case s1[i] < s2[j]:
			i++
		case s1[i] > s2[j]:
			j++
		default:
			intersection = append(intersection, s1[i])
			i++
			j++
		}
	}

	// Убираем лишний cap
	temp := make([]T, len(intersection))
	copy(temp, intersection)
	intersection = temp

	return
}

// Но если возможности упорядочивать элементы нет, то придётся использовать более тупой алгоритм,
// который имеет O(n*m)
// (Хотя скорее всего есть лучше алгосы, я просто не шарю так сильно)
func intersectSetsAny[T comparable](s1 []T, s2 []T) (intersection []T) {
	n, m := len(s1), len(s2)
	// Как обычно больше cap, чтобы не перевыделять память
	intersection = make([]T, 0, min(n, m))
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if s1[i] == s2[j] {
				intersection = append(intersection, s1[i])
			}
		}
	}

	// Убираем лишний cap
	temp := make([]T, len(intersection))
	copy(temp, intersection)
	intersection = temp

	return
}

func main() {
	var (
		a1 = []int{1, 7, 2, 9, 19, 34, 62, 3, 4}
		a2 = []int{1, 34, 62, 3, 4, 8, 22}

		s1 = []string{"a", "b", "d", "c", "88"}
		s2 = []string{"9", "b", "a", "0-0", "888888"}
	)

	fmt.Printf("Arrays:\n%#v\n%#v\nIntersection1:\n%#v\nIntersection2:\n%#v\n", a1, a2, intersectSetsOrderable(a1, a2), intersectSetsAny(a1, a2))
	fmt.Printf("Arrays:\n%#v\n%#v\nIntersection1:\n%#v\nIntersection2:\n%#v\n", s1, s2, intersectSetsOrderable(s1, s2), intersectSetsAny(s1, s2))
}
