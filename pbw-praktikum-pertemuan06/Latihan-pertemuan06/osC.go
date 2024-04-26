package main

import (
	"fmt"
	"os"
)

func main() {
	var err error
	fileInfo, err := os.Stat("SitiGhumaisa")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if fileInfo.IsDir() {
		fmt.Println("SitiGhumaisa adalah sebuah direktori.")
	} else {
		fmt.Println("SitiGhumaisa adalah sebuah file.")
	}
}

