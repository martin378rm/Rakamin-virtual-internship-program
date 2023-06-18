package main 

import "fmt"

func main() {
	
	var point = 8
		if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
		} else if (point > 5) {
		fmt.Println("lulus")
		} else if point == 4 {
		fmt.Println("hampir lulus")
		} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
		}


		// variabel temporary pada if statement
		var nilai = 8840.0
		if percent := nilai / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
		} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
		} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
		}


}