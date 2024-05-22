package main

import "fmt"

func main() {
	var input int

	fmt.Print("Masukkan Angka: ")
	fmt.Scan(&input)

	drawXYZ(input)
}

func drawXYZ(input int) {
	var angka = 1

	for i := 1; i <= input; i++ {
		for j := 1; j <= input; j++ {
			if angka%3 == 0 {
				fmt.Print("X ")
			} else if angka%2 > 0 {
				fmt.Print("Y ")
			} else if angka%2 == 0 {
				fmt.Print("Z ")
			}

			angka++
		}

		fmt.Println()
	}
}
