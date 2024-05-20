package main

import "fmt"

func main() {
	var angka int
	var bantuan int = 1
	var slice []int

	fmt.Print("Masukkan Angka: ")
	fmt.Scanln(&angka)

	fmt.Println("Faktor dari angka", angka, "adalah:")
	for bantuan <= angka {
		if angka%bantuan == 0 {
			slice = append(slice, bantuan)
		}
		bantuan++
	}
	for index := len(slice) - 1; index >= 0; index-- {
		fmt.Println(slice[index])
	}
}
