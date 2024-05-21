package main

import "fmt"

func main() {
	var angka int
	var bantuan int = 2
	var prima bool = true

	fmt.Print("Masukkan Angka: ")
	fmt.Scanln(&angka)

	for bantuan < angka {
		if angka%bantuan == 0 {
			prima = false
		}
		bantuan++
	}

	if angka <= 1 {
		prima = false
	}

	fmt.Println("Apakah angka", angka, "adalah bilangan Prima?", prima)
}
