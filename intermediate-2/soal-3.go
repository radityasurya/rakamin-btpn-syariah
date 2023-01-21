package main

import "fmt"

func main() {
	var fruits1 = []string{"Mango", "Grape", "Apple", "Watermelon", "Orange"}

	fruits1 = fruits1[2:4]
	fmt.Println(fruits1)

	// nama_array[start_index:last_index_before]

	// Contoh 2
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
