package main

import "fmt"

func main() {
	var harga float64
	var diskon float64

	fmt.Println("Masukkan Harga:")
	fmt.Scanln(&harga)
	fmt.Println("Masukkan Diskon:")
	fmt.Scanln(&diskon)

	diskon = diskon / 100 * harga
	var total float64 = harga - diskon

	fmt.Println("Harga yang harus dibayar adalah Rp", total)
}
