package main

import (
	"fmt"
)

func main() {
	states := map[string]int{
		"Georgia":      1231415,
		"Pennsylvania": 32772343,
		"LosAngles":    962523523,
		"NewYork":      1343534346,
		"Colorado":     7894566546,
	}

	if field, ok := states["Colorado"]; ok {
		fmt.Println("has that field", field)
	}

	if a := 23; a > 4 {
		fmt.Println("a is more than 4", a)
	}
	fmt.Println("a")
}
