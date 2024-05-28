package main

import (
	"fmt"
	"math"
)

func main() {
	var tinggi float64
	var radius float64

	fmt.Println("Masukkan Tinggi Tabung:")
	fmt.Scanln(&tinggi)
	fmt.Println("Masukkan Jari-Jari Tabung:")
	fmt.Scanln(&radius)

	var luas float64 = 2*3.14*math.Pow(radius, 2) + 2*3.14*radius*tinggi

	fmt.Println("Luas Permukaan Tabung Adalah:", luas)
}
