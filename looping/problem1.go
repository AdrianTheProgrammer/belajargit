package main

import "fmt"

func main() {
	var nilai float64

	fmt.Print("Masukkan Nilai Anda: ")
	fmt.Scanln(&nilai)

	if nilai >= 80 && nilai <= 100 {
		fmt.Println("========================")
		fmt.Println("A")
	} else if nilai >= 65 && nilai <= 79 {
		fmt.Println("========================")
		fmt.Println("B+")
	} else if nilai >= 50 && nilai <= 64 {
		fmt.Println("========================")
		fmt.Println("B")
	} else if nilai >= 35 && nilai <= 49 {
		fmt.Println("========================")
		fmt.Println("C")
	} else if nilai >= 0 && nilai <= 34 {
		fmt.Println("========================")
		fmt.Println("D")
	} else {
		fmt.Println("========================")
		fmt.Println("Nilai Tidak Valid")
	}
}

// 80 > 100 : A
// 65 > 79 : B+
// 50 > 64 : B
// 35 > 49 : C
// 0 > 34 : D
