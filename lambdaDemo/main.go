package main

import (
	"fmt"
	"sync"
)

func main() {
	// wait groups
	wg := &sync.WaitGroup{}
	// add done , and wait

	// mutexes critical section

	// mut := &sync.Mutex
	wg.Add(2)
	go func() {
		fmt.Println("Hello")
		wg.Done()
	}()
	go func() {
		fmt.Println("World")
		wg.Done()
	}()
	fmt.Println("Fin")
	wg.Wait()
	fmt.Println("Exit")
	// select {}
}
