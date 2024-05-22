package main

import "fmt"

func main() {
	var angka int
	var prima bool = true
	var fullprima bool = true

	fmt.Print("Masukkan Angka: ")
	fmt.Scanln(&angka)

	var angka2 int = angka

	if angka < 2 {
		prima = false
		fullprima = false
	} else {
		for bantuan1 := 2; bantuan1 < angka; bantuan1++ {
			if angka%bantuan1 == 0 {
				prima = false
				fullprima = false
			}
		}
		for angka2 > 0 && prima {
			var digit int = angka2 % 10
			if digit < 2 {
				fullprima = false
			} else {
				for bantuan2 := 2; bantuan2 < digit; bantuan2++ {
					if digit%bantuan2 == 0 {
						fullprima = false
					}
				}
			}
			angka2 /= 10
		}
	}

	if fullprima {
		fmt.Println("Angka", angka, "ADALAH Full Prima")
	} else {
		fmt.Println("Angka", angka, "BUKANLAH Full Prima")
	}
}
