package main

import (
	"fmt"
)

const (
	a = iota
	_
	b
	c
)

const (
	b1 = 2 << 2
	b2 = 2 << 3
	b3 = 2 << 4
	b4 = 3 << 4
	b5 = 5 << 2
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeAfrica

	canSeeSanFrancisco
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	//const a int64 = 78
	/*fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(b4)
	fmt.Println(b5)
	*/
	callBytes()
}

func callBytes() {
	var roles byte = isAdmin | canSeeSanFrancisco | canSeeEurope
	fmt.Printf("%b", roles)
	fmt.Printf("\n Is Admin %v", isAdmin&roles == isAdmin)
	fmt.Println("\n", isAdmin)
	fmt.Println(isHeadquarters)
	fmt.Println(canSeeAfrica)
	fmt.Println(canSeeSanFrancisco)
	fmt.Println(canSeeEurope)
	fmt.Println(canSeeNorthAmerica)
}
