package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "martin"
	names[1] = "rivaldo"
	names[2] = "pardamean"
	names[3] = "manurung"
	names[0] = "hello"
	fmt.Println(
		names[0],
		names[1],
		names[2],
		names[3],
	)

	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	fruits[0] = "mango"
	fmt.Println("Jumlah element => \t\t", len(fruits))
	fmt.Println("Isi semua element => \t", fruits)

	// inisialisasi nilai awal array tanpa jumlah elemen
	var numbers = [...]int{2, 3, 2, 4, 3}
	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	//array multidimensi
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}

	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	for _, fruit := range fruits {
		fmt.Printf("nama buah : %s\n", fruit)
	}

	var buah = make([]string, 2)
	buah[0] = "apple"
	buah[1] = "manggo"
	fmt.Println(append(buah, "as"))
}
