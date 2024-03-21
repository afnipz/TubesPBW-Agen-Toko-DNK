package main

import (
	"bufio"
	"fmt"
	"os"
)

type Mahasiswa struct {
	Nama    string
	NPM     string
	Jurusan string
}

func main() {
	var jumlahMahasiswa int
	var npmCari string

	// Membuat scanner baru untuk membaca input
	scanner := bufio.NewScanner(os.Stdin)

	// Meminta jumlah mahasiswa dari pengguna
	fmt.Print("\n Masukkan Jumlah Mahasiswa: ")
	fmt.Scanln(&jumlahMahasiswa)

	// Membuat map untuk menyimpan data mahasiswa
	mahasiswas := make(map[string]Mahasiswa)

	// Meminta data mahasiswa dari pengguna
	for i := 1; i <= jumlahMahasiswa; i++ {
		var m Mahasiswa
		fmt.Printf("\n Data Mahasiswa ke-%d\n", i)
		fmt.Print(" Masukkan nama         : ")

		// Membaca input nama mahasiswa
		scanner.Scan()
		m.Nama = scanner.Text()

		fmt.Print(" Masukkan NPM          : ")
		fmt.Scanln(&m.NPM)
		fmt.Print(" Masukkan jurusan      : ")

		// Membaca input jurusan mahasiswa
		scanner.Scan()
		m.Jurusan = scanner.Text()

		mahasiswas[m.NPM] = m
	}

	// Menampilkan daftar nama mahasiswa
	fmt.Println("\n~~~~~ Daftar Nama Mahasiswa ~~~~~")
	for _, mahasiswa := range mahasiswas {
		fmt.Println(mahasiswa.Nama)
	}

	// Mencari data mahasiswa berdasarkan NPM
	fmt.Print("\nMasukkan NPM untuk mencari data mahasiswa: ")
	fmt.Scanln(&npmCari)
	if mahasiswa, found := mahasiswas[npmCari]; found {
		fmt.Println("\n~~~~~ Informasi Mahasiswa ~~~~~")
		fmt.Println("Nama    :", mahasiswa.Nama)
		fmt.Println("NPM     :", mahasiswa.NPM)
		fmt.Println("Jurusan :", mahasiswa.Jurusan)
	} else {
		fmt.Println("\n\nMahasiswa dengan NPM", npmCari, "tidak ditemukan.\n")
	}
}
