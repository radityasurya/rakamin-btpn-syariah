package main

import "fmt"

type Employee struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Salary  int    `json:"salary"`
}

func main() {

	// Cara 1
	var newEmployee = Employee{Name: "John", Age: 21, Address: "California, USA", Salary: 1000}
	fmt.Println(newEmployee)

	// Cara 2
	var newEmployee2 = Employee{}
	newEmployee2.Name = "John"
	newEmployee2.Age = 21
	newEmployee2.Address = "California, USA"
	newEmployee2.Salary = 1000
	fmt.Println(newEmployee2)
}
