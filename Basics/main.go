package main

import (
	"fmt"
	// "golang.org/x/exp/constraints"
)

// func Add(a, b int) int {
// 	return a + b
// }

// func AddFloat(a, b float64) float64 {
// 	return a + b
// }

// type Num interface {
// 	int | int8 | int16 | float32 | float64
// }

// func Add[T Num](a T, b T) T {
// 	return a + b
// }

type UserID int

func Add[T ~int | float64](a T, b T) T {
	return a + b
}

// func Add[T constraints.Ordered](a T, b T) T {
// 	return a + b
// }

func main() {
	// result := Add(1.0, 2.0)
	// result := AddFloat(2.0, 3.2)
	a := UserID(1)
	b := UserID(2)

	result := Add(a, b)
	fmt.Printf("result: %+v\n", result)
}
