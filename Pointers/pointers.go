package main

import (
	"fmt"
)

func main() {
	var a int = 42
	var b *int = &a
	*b = 23

	fmt.Println(a, *b)
}
