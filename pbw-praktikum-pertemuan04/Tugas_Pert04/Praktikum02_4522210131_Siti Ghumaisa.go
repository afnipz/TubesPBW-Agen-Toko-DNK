package main

import "fmt"

// mengurutkan array dalam satu kali perulangan
func main() {
	fmt.Println("\n========Bubble Sort========")
	var arrayNumber [20]int
	data := [20]int{12, 8, 1, 15, 9, 11, 20, 17, 14, 10, 16, 13, 19, 2, 18, 9, 3, 4, 7, 6} 

	for i := 0; i < len(data); i++ {
		arrayNumber[i] = data[i]
	}

	// Menampilkan sebelum pengurutan
	fmt.Println("\nSebelum dilakukan pengurutan:")
	for _, val := range arrayNumber {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	// Algoritma Bubble Sort
	for i := 0; i < len(arrayNumber); {
		if i != (len(arrayNumber) - 1) {
			if arrayNumber[i] > arrayNumber[i+1] {
				x := arrayNumber[i]
				arrayNumber[i] = arrayNumber[i+1]
				arrayNumber[i+1] = x
				i--
			}
		}

		if i > 0 {
			if arrayNumber[i] < arrayNumber[i-1] {
				x := arrayNumber[i]
				arrayNumber[i] = arrayNumber[i-1]
				arrayNumber[i-1] = x
				i--
			}
		}
		i++
	}

	// Menampilkan setelah pengurutan
	fmt.Println()
	fmt.Println("Setelah dilakukan pengurutan:")
	for _, val := range arrayNumber {
		fmt.Printf("%d ", val)
	}
	fmt.Println("\n")
}
