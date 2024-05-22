package main

import "fmt"

func main() {
	var asterisk int

	fmt.Print("Masukkan Jumlah Asterisk: ")
	fmt.Scan(&asterisk)

	asterisk_triangle(asterisk)
}

func asterisk_triangle(asterisk int) {
	var spasi int = asterisk

	for i := 0; i < asterisk; i++ {
		for j := spasi; j > 0; j-- {
			fmt.Print(" ")
		}

		for k := 0; k <= i; k++ {
			fmt.Print("* ")
		}

		fmt.Println()
		spasi--
	}
}
