package main

import "fmt"

// Fungsi dibawah sama saja dengan
// 7 x (6 x (5 x (4 x (3 x (2 x (1 x 1))))))

func main() {
	// Hitung faktorial dari 7
	fmt.Println()
	fmt.Println("Hasil faktorial 7:")
	fmt.Printf("Hasil dari perkalian factorial 7 adalah = %d\n", angkaFaktorial(7) )
	fmt.Println()
}


func angkaFaktorial(n int) int {
	if n == 1 {
		fmt.Printf("%d x 1 = %d\n", n, n)
		return 1
	}

	hasilFaktorialSebelumnya := angkaFaktorial(n - 1)

	result := n * hasilFaktorialSebelumnya
	fmt.Printf("%d x %d = %d\n", n, hasilFaktorialSebelumnya, result)

	return result
}
