package main

import "fmt"

func main() {
	fruits := []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0])

	var fruitsA = []string{"apple", "grape"}     // slice
	var fruitsB = [2]string{"banana", "melon"}   // array
	var fruitsC = [...]string{"papaya", "grape"} // array

	fmt.Println(fruitsA, fruitsB, fruitsC)

	var buah = []string{"apple", "grape", "banana", "melon"}
	var newFruits = buah[0:2]
	fmt.Println(newFruits)
}
