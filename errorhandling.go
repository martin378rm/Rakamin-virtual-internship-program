package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var input string
	fmt.Print("Type some number: ")
	fmt.Scanln(&input)
	var number int
	var err error
	number, err = strconv.Atoi(input)

	// jika tidak ada error atau error == kosong
	if err == nil {
		fmt.Println(number, "is number")
	} else {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error())
	}

	// panic == dengan throw
	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)
	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		panic(err.Error())
	}
	fmt.Println("end")

	// recover
	defer catch()
	var names string
	fmt.Print("Type your name: ")
	fmt.Scanln(&names)
	if valid, err := validate(names); valid {
		fmt.Println("halo", names)
	} else {
		panic(err.Error())
	}
	fmt.Println("end")

}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}
