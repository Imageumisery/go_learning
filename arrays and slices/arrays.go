package main

import (
	"fmt"
)

func main() {
	var nums [7]int = [7]int{34, 56, 78}
	nums[2] = 5

	someSlice := []int{3, 4, 5}

	numbers := [...]int{45, 23, 51}
	b := &numbers
	//просто присваивание тут означает копирование содержимого
	//чтобы присвоить адрес переменной надо пользоваться специальным оператором &

	/*fmt.Printf("numbers: %v\n", numbers)
	fmt.Printf("Capacity: %v\n", cap(nums))
	fmt.Printf("length: %v\n", len(nums))
	fmt.Printf("nums: %v\n", nums)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("num: %v\n", numbers[0])
	fmt.Printf("numbers: %v\n", numbers)
	fmt.Printf("num: %v\n", numbers[0])
	b[0] = 2
	*/
	fmt.Printf("b: %v\n", b)
	//slices
	a := someSlice
	a[0] = 0

	fmt.Printf("copy of slice(a): %v\n", a)
	fmt.Printf("slice: %v\n", someSlice)

}
