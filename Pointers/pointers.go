package main

import (
	"fmt"
)

func main() {
	var a int = 42
	var b *int = &a
	*b = 23
	str := "gogogo"

	d := 44
	c := &d
	*c = 11

	fmt.Println(a, *b)
	fmt.Println(*c)
	changeValue2(str)
	//fmt.Println(changeValue(&str))
	fmt.Println("\n" + str)
	fmt.Printf("\n %v", &d)
	fmt.Printf("\n %v\n", &c)

}

//passing pointer to the variable
func changeValue(str *string) *string {
	*str = "changed"
	return str
}

//passing value of variable
func changeValue2(str string) {
	str = "changed"
	fmt.Printf("String within second function: %s", str)
}
