package main

import (
	"fmt"
	"os"
)

func main() {
	// Membuat direktori baru.
	err := os.Mkdir("SitiGhumaisa", 0131)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Direktori 'SitiGhumaisa' berhasil dibuat.")
}
