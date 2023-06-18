package main

import (
	"fmt"
	"math"
)

func main() {
	// desimal
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	// float
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.1f\n", decimalNumber)

	var floor = math.Ceil(decimalNumber)
	fmt.Println(floor)

	// boolean
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	// string
	var message string = "Halo"
	fmt.Printf("message: %s \n", message)
	fmt.Println(message[0])
	char := string(message[0])
	fmt.Println(char)
}
