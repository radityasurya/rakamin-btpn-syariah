package main

import "fmt"

func main() {
	var fruits1 = []string{"Mango", "Grape", "Apple", "Watermelon", "Orange"}
	fmt.Println(fruits1)

	fruits1 = append(fruits1[:5], "Avocado")
	fmt.Println(fruits1)

	// fruits1 = append(fruits1[:3], "Rambutan")
	// fmt.Println(fruits1)
}
