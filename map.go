package main

import (
	"fmt"
)

func main() {

	chicken := map[string]int{}
	chicken["januari"] = 50
	chicken["februari"] = 40
	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])
	fmt.Println(chicken)

	var _ = map[string]int{}
	var _ = make(map[string]int)
	var _ = *new(map[string]int)

	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}
	for key, val := range chicken2 {
		fmt.Println(key, " \t:", val)
	}

	fmt.Println("\n\n")
	delete(chicken2, "januari")

	for key, val := range chicken2 {
		fmt.Println(key, " \t:", val)
	}
	fmt.Printf("chicken: %v\n", chicken)

	var chicken5 = map[string]int{"januari": 50, "februari": 40}
	value, isExist := chicken5["januari"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	fmt.Println("\n\n ==== Kombinasi Slice dan Map ====")
	sliceCombineMap()
}

func sliceCombineMap() {

	// deklarasi dengan memberikan tipe pada tiap elemen
	// var chickens = []map[string]string{
	// 	map[string]string{"name": "chicken blue", "gender": "male"},
	// 	map[string]string{"name": "chicken red", "gender": "male"},
	// 	map[string]string{"name": "chicken yellow", "gender": "female"},
	// }

	// deklarasi tanpa memberikan tipe pada tiap elemen
	var chickens = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}
}
