package main

import (
	"fmt"
	"os"
)

func main() {
	// Membuat direktori baru.
	err := os.Chmod("SitiGhumaisa", 0131)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Izin 'SitiGhumaisa' telah diubah menjadi 0131.")
}
