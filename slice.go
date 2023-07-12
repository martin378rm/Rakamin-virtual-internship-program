package main

import (
	"fmt"
)

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

	arrays := []int{1}
	fmt.Println(arrays)

	for i, array := range arrays {
		fmt.Printf("Ini index ke = %d dengan nilai = %d\n\n", i, array)
	}

	for i := 0; i < len(arrays); i++ {
		if arrays[i]%2 == 0 {
			fmt.Println("bilangan genap")
		} else {
			fmt.Println("bilangan ganjil")
		}
	}

	arrays = append(arrays, 3)
	fmt.Println(len(arrays))
	fmt.Print(cap(arrays))

	fruitsSliceReference(fruits)
}

func fruitsSliceReference(fruits []string) {
	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]
	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"
	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]

}
