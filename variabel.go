package main

import "fmt"

func main() {

	firstName := "Martin"
	var lastName string = "Rivaldo"
	fmt.Println(firstName, lastName)
	fmt.Printf("halo %s %s!\n", firstName, lastName)

	var fullname = "Martin Rivaldo Manurung"
	fmt.Println(fullname)

	name := new(string)
	fmt.Println(*name)

}
