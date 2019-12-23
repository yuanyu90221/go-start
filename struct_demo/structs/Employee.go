package structs

import "fmt"

type Employee struct {
	Name string
	Age int
	Designation string
	Salary int
}

func(emp Employee) ShowDetails() {
	fmt.Println("User Name:", emp.Name)
}