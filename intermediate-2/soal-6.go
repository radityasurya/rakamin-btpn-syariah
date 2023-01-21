package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Salary  int    `json:"salary"`
}

func main() {

	jsonString := `{"name" : "John", "age" : 21, "address" : "California, USA", "salary" : 1000}`

	var result Employee
	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	fmt.Println("Nama : ", result.Name+"\n"+"Alamat : "+result.Address)
}
