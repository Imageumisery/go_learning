package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Err: ", err)
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}
