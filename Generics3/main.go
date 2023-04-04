package main

import (
	"fmt"
)

type CustomMap[T comparable, V int | string] map[T]V

func main() {

	m := make(CustomMap[int, string])
	m[3] = "3"

	m1 := make(CustomMap[string, string])
	m1["name"] = "rohan"

	fmt.Printf("map1: %+v\nmap2: %+v\n", m, m1)
}

