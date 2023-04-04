package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

//input:=[1,2,3], (n)=>n*2
//output:[2,4,6]

// func MapValues(value []int, mapFunc func(int) int) []int {
// 	var newValues []int
// 	for _, v := range value {
// 		newValue := mapFunc(v)
// 		newValues = append(newValues, newValue)
// 	}
// 	return newValues
// }

func MapValues[T constraints.Ordered](value []T, mapFunc func(T) T) []T {
	var newValues []T
	for _, v := range value {
		newValue := mapFunc(v)
		newValues = append(newValues, newValue)
	}
	return newValues
}

func main() {
	result := MapValues([]float32{1.2, 2.3, 3.5}, func(n float32) float32 {
		return n * 2.5
	})
	fmt.Printf("result: %+v\n", result)
}
