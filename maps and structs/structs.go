package main

import (
	"fmt"
)

type Developer struct {
	name       string
	language   string
	age        int
	experience int
}

type Base struct {
	someField int
}

type container struct {
	Base
	name     string
	capacity int
}

func main() {
	CppDev := Developer{
		name:       "Jon",
		language:   "Cpp",
		age:        21,
		experience: 2,
	}

	co := container{
		Base: Base{
			someField: 9,
		},
		name:     "Jojo",
		capacity: 78,
	}

	SomeNewDev := CppDev
	SomeNewDev.age = 46
	fmt.Println(CppDev.age)
	fmt.Println(SomeNewDev.age)
	fmt.Println(co.Base.someField)
}
