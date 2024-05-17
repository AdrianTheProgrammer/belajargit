package main

import "fmt"

func main() {
	var nama string

	fmt.Println("Masukkan Nama Anda:")
	fmt.Scanln(&nama)

	fmt.Println("Hello", nama, ", Saya Golang, bahasa yang sangat menyenangkan.")
}
