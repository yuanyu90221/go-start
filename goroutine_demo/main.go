package main

import (
	"fmt"
	"time"
)

func heavy() {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("Heavy")
	}
}

func superHeavy() {
	for {
		time.Sleep(time.Second * 2)
		fmt.Println("Super heavy")
	}
}
func main() {
	go heavy()
	go superHeavy()
	fmt.Println("Fin")
	// time.Sleep(time.Second * 5)
	select {}
}
