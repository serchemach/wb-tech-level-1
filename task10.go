package main

import (
	"fmt"
	// "math"
	"math/rand"
)

func generateRandomTemps(numberOfMeasurements int) (result []float64) {
	result = make([]float64, numberOfMeasurements)

	for i := range numberOfMeasurements {
		result[i] = rand.Float64()*60.0 - 20.0
	}
	return
}

// Eсть небольшая проблема со значениями, которые попадают в 0 группу, потому что
// наши группы симметричны относительно -
// Это можно решить, добавив +0 и -0 группы, чтобы у нас ВСЕ группы были с шагом в 10
// Но я не стал так извращаться
func aggregateTemps(temps []float64) (result map[int][]float64) {
	result = make(map[int][]float64)
	for _, val := range temps {
		rounded := int(val)
		bottomDecade := (rounded / 10) * 10

		vals, ok := result[bottomDecade]
		if !ok {
			result[bottomDecade] = make([]float64, 1, len(temps))
			result[bottomDecade][0] = val
			continue
		}
		result[bottomDecade] = append(vals, val)
	}

	return
}

func main() {
	exampleArr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	exampleAgg := aggregateTemps(exampleArr)
	fmt.Printf("Example array result: %#v\n", exampleAgg)

	testArr := generateRandomTemps(20)
	fmt.Printf("Test arr: %#v\n", testArr)
	testAgg := aggregateTemps(testArr)
	fmt.Printf("Test array result: %#v\n", testAgg)
}
