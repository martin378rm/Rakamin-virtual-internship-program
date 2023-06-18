package main

import "fmt"

func main () {

	// for i := 1; i <= 5; i++ {
	// 	for j:= 1; j <= i; j++{
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Print("\n")
	// }

	k := 0

	//dengan argumen hanya kondisi

	// for k < 5 {
	// 	fmt.Println("Angka", k)
	// 	k++
	//  }
	 

	// tanpa argumen
	 for {
		fmt.Println("Angka", k)
		k++
		if k == 5 {
		break
		}
	}

	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	for i, fruit := range fruits {
	fmt.Printf("elemen %d : %s\n", i, fruit)
	}


	// penggunaan label untuk keluar dari program
	outerLoop:
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if i == 3 {
					break outerLoop
				}
		fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
		}

	


}