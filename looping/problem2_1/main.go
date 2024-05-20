package main

import "fmt"

func main() {
	var angka int
	var bantuan int = 1

	fmt.Print("Masukkan Angka: ")
	fmt.Scanln(&angka)

	fmt.Println("Faktor dari angka", angka, "adalah:")
	for bantuan <= angka {
		if angka%bantuan == 0 {
			fmt.Println(bantuan)
		}
		bantuan++
	}
}
