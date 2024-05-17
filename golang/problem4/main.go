package main

import "fmt"

func main() {
	var alas float64
	var tinggi float64

	fmt.Println("Masukkan Alas:")
	fmt.Scanln(&alas)
	fmt.Println("Masukkan Tinggi")
	fmt.Scanln(&tinggi)

	var luas float64 = 0.5 * alas * tinggi

	fmt.Println("Luas Segitiganya Adalah:", luas)
}
