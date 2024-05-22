package main

import (
	"fmt"
)

func main() {
	var input int

	fmt.Print("Masukkan Angka: ")
	fmt.Scan(&input)

	tabel_perkalian(input)

}

func tabel_perkalian(input int) {
	var angka int = 1

	for i := 1; i <= input; i++ {
		for j := 1; j <= input; j++ {
			fmt.Print(j*angka, " ")
		}
		angka++
		fmt.Println()
	}
}
