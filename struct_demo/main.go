package main

import (
	"fmt"
	"./structs"
)
func UpdateEmployee(empDetails *structs.Employee, replaceName string) {
	empDetails.Name = replaceName
}

func main() {
	var newEmployee =  structs.Employee{Name:"Json", Age:11, Designation:"sty", Salary:500000 }
	fmt.Println(newEmployee)
	UpdateEmployee(&newEmployee, "Jefery");
	fmt.Println("new",newEmployee.Name)
	newEmployee.ShowDetails()
	var empEmployee = new(structs.Employee)
	fmt.Println("emp:",empEmployee.Name)
	empEmployee.ShowDetails()
}
