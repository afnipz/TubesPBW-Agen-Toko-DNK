package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Membuat scanner untuk input dari konsol
	scanner := bufio.NewScanner(os.Stdin)

	// Input nama
	fmt.Print("\n Masukkan nama Anda : ")
	scanner.Scan()
	nama := scanner.Text()

	// Input usia
	fmt.Print(" Masukkan usia Anda : ")
	scanner.Scan()
	var usia int
	fmt.Sscanf(scanner.Text(), "%d", &usia)

	// Menentukan kategori usia
	var kategori string
	if usia < 18 {
		kategori = "Anak-anak"
	} else {
		kategori = "Dewasa"
	}

	// Menampilkan pesan selamat datang
	fmt.Printf("\n Selamat datang, %s! Anda termasuk kategori %s.\n\n", nama, kategori)
}
