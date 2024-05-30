package main

import "fmt"

func main() {
	var a1, a2, a3, a4, a5, a6 int

	fmt.Print("Masukkan Nilai A1: ")
	fmt.Scanln(&a1)
	fmt.Print("Masukkan Nilai A2: ")
	fmt.Scanln(&a2)
	fmt.Print("Masukkan Nilai A3: ")
	fmt.Scanln(&a3)
	fmt.Print("Masukkan Nilai A4: ")
	fmt.Scanln(&a4)
	fmt.Print("Masukkan Nilai A5: ")
	fmt.Scanln(&a5)
	fmt.Print("Masukkan Nilai A6: ")
	fmt.Scanln(&a6)

	min, max := min_max(&a1, &a2, &a3, &a4, &a5, &a6)

	fmt.Println("Min:", min, "Max:", max)
}

func min_max(p_a1, p_a2, p_a3, p_a4, p_a5, p_a6 *int) (min, max int) {
	pointer_slice := []*int{p_a1, p_a2, p_a3, p_a4, p_a5, p_a6}
	min = *pointer_slice[0]

	for _, angka := range pointer_slice {
		value := *angka

		if value < min {
			min = value
		} else if value > max {
			max = value
		}
	}

	return min, max
}
