package main

import "fmt"

func main() {
	name := "martin"

	fmt.Println(name)

	// deklarasi multi konstanta
	const (
		// ]type inference
		square   = "kotak"
		floatNum = 2.2

		//  manifest typing
		isToday bool  = true
		numeric uint8 = 1
	)

	fmt.Println(square, isToday, numeric, floatNum)

	const satu, dua = 1, 2
	fmt.Println(satu)
	fmt.Println(dua)

	const three, four string = "tiga", "empat"
	fmt.Println(three, four)

	Martin := new(Person)

	*&Martin.name = "Martin"
	*&Martin.age = 21

	fmt.Println(*Martin)

	var angka int
	fmt.Print("Masukan Angka .....")
	fmt.Scanln(&angka)
	cetak_pola(angka)
}

type Person struct {
	name string
	age  int
}

func cetak_pola(jumlah_baris int) {
	for i := 1; i <= jumlah_baris; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		for k := 1; k <= 2*(jumlah_baris-i); k++ {
			fmt.Print("1")
		}
		for l := 1; l <= i; l++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
