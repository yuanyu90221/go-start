package main

import "fmt"

func main() {
	input := "Hello"
	printReverse(&input)
}

func printReverse(str *string) {
	fmt.Print(*str)
	// if *str !=""{
	// 	return
	// }
	// printReverse(&(str+1))
	// fmt.Print(*str);
}