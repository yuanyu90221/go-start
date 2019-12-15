package main

import "fmt"

func Anything(anything interface{}) {
	fmt.Println(anything)
}
func main() {
	fmt.Println("test")
	Anything(2.44)
	Anything("name")
	Anything(struct{}{})

	mymap := make(map[string]interface{})

	mymap["name"] = "json"
	mymap["age"] = 10
	fmt.Println(mymap)
}
