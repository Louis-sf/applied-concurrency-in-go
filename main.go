package main

import (
	"fmt"
	/*all types in the sync package should
	be passed by pointer to functions*/
	"sync"
)

func main() {
	var wg sync.WaitGroup
	//add 1 because we only wait for 1 goroutine (hello)
	wg.Add(1)
	//pass by reference to avoid copying waitgroup instance
	go hello(&wg)
	wg.Wait()
	goodbye()
}

func hello(wg *sync.WaitGroup) {
	//good practice to defered Done at the begining of the task
	defer wg.Done()
	fmt.Println("Hello, world!")
}

func goodbye() {
	fmt.Println("Goodbye, world!")
}
