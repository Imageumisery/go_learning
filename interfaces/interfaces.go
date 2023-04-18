package main

import (
	"fmt"
)

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}
type IntCounter int


type Human struct {
	firstname string
	lastname  string
	age       string
	country   string
}

type Incrementer interface {
	Increment() int
	
}

type DomesticAnimal interface {
	RecieveAffection(from Human)
	GiveAffection(to Human)
}

type Cat struct {
	name string
}

type Dog struct {
	name string
}
type Snake struct {
	name string
}

func (cw ConsoleWriter) Write(data []byte)(int, error)  {
	n, err := fmt.Println(string(data))	
	return n, err
}

func (c Cat) RecieveAffection(f Human) {
	fmt.Printf("The cat named %s has recieved affection from Human named %s\n", c.name, f.firstname)
}

func (c Cat) GiveAffection(to Human) {
	fmt.Printf("The cat named %s has given affection to Human named %s\n", c.name, to.firstname)
}

func (d Dog) RecieveAffection(from Human) {
	fmt.Printf("The dog named %s has recieved affection to Human named %s\n", d.name, from.firstname)
}

func (d Dog) GiveAffection(to Human) {
	fmt.Printf("The dog named %s has given affection to Human named %s\n", d.name, to.firstname)
}

func Pet(animal DomesticAnimal, human Human) {
	animal.GiveAffection(human)
	animal.RecieveAffection(human)
}

func (ic *IntCounter)Increment() int {
	*ic++
	return int(*ic)
	
}

func main() {
	John := Human{
		firstname: "John",
	}
	MyCat := Cat{
		name: "Moki",
	}

	MyDog := Dog {
		name: "Bethoveen",
	}

	/*
	MySnake := Snake {
		name: "NodiraFss",
	} */

	//Pet(MySnake, John) won't compile
	Pet(MyCat, John)
	Pet(MyDog, John)
	
	myInt := IntCounter(0)
	var inc Incrementer = &myInt

	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hi Mom!"))
}
