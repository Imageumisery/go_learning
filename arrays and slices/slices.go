package main

import (
	"fmt"
)

func main() {

	someSlice := []int{3, 4, 5, 6, 7, 8, 9}
	a := someSlice[2:7]
	b := someSlice[1:3]
	fmt.Printf("slice: %v\n", someSlice)
	someSlice = append(someSlice, []int{4, 6, 8, 123, 56, 23}...)
	numbers := [...]int{45, 23, 51}
	makeSlice := make([]int, 10, 100)

	numbers[0] = 0

	fmt.Println(cap(makeSlice))
	fmt.Println(len(makeSlice))
	fmt.Printf("slice: %v\n", someSlice)
	fmt.Printf("slice1: %v\n", someSlice[1:14])
	fmt.Printf("len: %v\n", len(someSlice))
	fmt.Printf("copy of slice(a): %v\n", a)
	fmt.Printf("b: %v\n", b)
}
