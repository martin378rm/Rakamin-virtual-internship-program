package main

import "fmt"
func main () {
	const name = "martin"

	fmt.Println(name)

	// deklarasi multi konstanta
	const (
		// metode type inference
		square = "kotak"
		// metode manifest typing
		isToday bool = true
		numeric uint8 = 1
		floatNum = 2.2
	 )

	 fmt.Println(square,isToday,numeric,floatNum)
}