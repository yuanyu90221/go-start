package main

import (
	"fmt"
)

func main() {
	// var m1 string
	// m1 := "my name"
	// m2 := "name"
	// fmt.Println(strings.Contains(m1, m2))
	// m3 := strings.ToUpper(m1)
	// // fmt.Println(strings.ReplaceAll(m1, "m", "NO"))
	// fmt.Println(strings.Split(m1, " "))
	// fmt.Println(m3)
	todo()
	pointer()
	m1, m2 := 2, 3
	fmt.Println(m1, m2)
	swap(&m1, &m2)
	fmt.Println(m1, m2)
	structMain()
}

// func main() {
// 	m1 :=2
// 	m2 :=3
// 	fmt.Println(m1 + m2)
// }

func todo() {
	// var arr []int
	arr := []int{1, 2, 3, 4}
	arr2 := []string{"hi", "my", "name"}
	arr2 = append(arr2, "test", "1", "2")
	fmt.Println(arr, arr2)
}
func swap(m1, m2 *int) {
	var temp int
	temp = *m2
	*m2 = *m1
	*m1 = temp
}
func pointer() {
	m1 := 2
	ptr := &m1
	fmt.Println(*ptr)
}

// structure
type Car struct {
	Name    string
	Age     int
	ModelNo int
}

func (c Car) Print() {
	fmt.Println(c)
}
func (c Car) Drive() {
	fmt.Println("driving")
}
func (c Car) GetName() string {
	return c.Name
}
func structMain() {
	// c := Car{"chevy", 1, 2}
	// var c1 Car
	c := Car{
		Name:    "Cherry",
		Age:     1,
		ModelNo: 2,
	}
	// fmt.Println(c)
	c.Print()
	c.Drive()
	fmt.Println(c.GetName())
}
