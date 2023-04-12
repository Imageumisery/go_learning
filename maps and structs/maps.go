package main

import (
	"fmt"
)

//return order of the map is not guaranteed
func main() {
	states := map[string]int{
		"Georgia":      1231415,
		"Pennsylvania": 32772343,
		"LosAngles":    962523523,
		"NewYork":      1343534346,
		"Colorado":     7894566546,
	}
	fmt.Println("a")
	fmt.Println(states)
	delete(states, "Colorado")
	states["NewYork"] = 1
	states["Keek"] = 89765
	fmt.Println(states)
}
