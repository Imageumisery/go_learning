package main

import (
	"fmt"
)

func main() {
	IPerson := Person{
		name: "Bob",
		age:  17,
	}
	fmt.Println("Person name: ", IPerson.name)
	IPerson.changeName()

	s := sum(1, 2, 3, 4, 5)
	fmt.Println("result of sum:", s)
	divideResult, err := divide(4.2, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("divideResult: ", divideResult)
}

type Person struct {
	name string
	age  int
}

func (IPerson *Person) changeName() {
	IPerson.name = "1234"
	fmt.Println("the name is changed succesfully: ", IPerson.name)
}
func sum(values ...int) int {
	fmt.Println("The sum is:", values)
	var result int
	for _, v := range values {
		result += v
	}
	return result
}

func divide(a, b float64) (float64, error) {
	if b == 0 || a == 0 {
		return 0.0, fmt.Errorf("Cannot divide to zero")
	}
	return a / b, nil
}
