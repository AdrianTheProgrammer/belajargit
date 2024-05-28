package main

import "fmt"

func main() {
	remove_duplicates([]int{2, 3, 3, 3, 6, 9, 9})
	remove_duplicates([]int{2, 3, 4, 5, 6, 9, 9})
	remove_duplicates([]int{2, 2, 2, 11})
	remove_duplicates([]int{2, 2, 2, 11})
	remove_duplicates([]int{1, 2, 3, 11, 11})
}

func remove_duplicates(slice []int) {
	hasil := []int{}
	start := 0

	for i := 0; i < len(slice); i++ {
		for j := 1 + start; j < len(slice); j++ {
			if slice[i] == slice[j] {
				slice[j] = 0
			}
		}

		start++

		if slice[i] != 0 {
			hasil = append(hasil, slice[i])
		}
	}

	fmt.Println("Array:", hasil, "Length:", len(hasil))
}
