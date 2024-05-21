package main

import (
	"fmt"
	"math"
)

func main() {
	var angka float64
	var pangkat float64
	var hasil float64

	fmt.Print("Masukkan Angka: ")
	fmt.Scanln(&angka)
	fmt.Print("Masukkan Pangkat: ")
	fmt.Scanln(&pangkat)

	hasil = math.Pow(angka, pangkat)
	fmt.Println("Hasil dari", angka, "^", pangkat, "=", hasil)
}
