package main

import (
	"fmt"
	"runtime"
	"sync"
)
var wg = sync.WaitGroup{}
var m = sync.RWMutex{}
var counter int = 0

func main() {
	wg.Add(1)
	go printMesg("Hi Mom, I'm your programmer son ")
	wg.Wait()
	println("qwesad")
	runtime.GOMAXPROCS(100)
	fmt.Println(runtime.GOMAXPROCS(-1))

	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
	    m.Lock()
		go increment()
	}
	wg.Wait()
}

func printMesg(msg string)  {
println(msg)
	wg.Done()
}

func sayHello()  {
	fmt.Printf("Hello #%v\n", counter)	
	m.RUnlock()
	wg.Done()
}

func increment()  {
	counter++
	m.Unlock()
	wg.Done()
}