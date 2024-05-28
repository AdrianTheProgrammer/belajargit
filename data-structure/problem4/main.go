package main

import (
	"fmt"
	"slices"
)

func main() {
	slice_max([]int{2, 1, 5, 1, 3, 2}, 3)
	slice_max([]int{2, 3, 4, 1, 5}, 2)
	slice_max([]int{2, 1, 4, 1, 1}, 2)
	slice_max([]int{2, 1, 4, 1, 1}, 3)
	slice_max([]int{2, 1, 4, 1, 1}, 4)
}

func slice_max(slice []int, jarak int) {
	start := 0
	end := jarak
	hasil := 0
	max := []int{}

	for i := 0; i <= len(slice)-jarak; i++ {
		hasil = 0

		for j := 0 + start; j < end; j++ {
			hasil = hasil + slice[j]
		}

		max = append(max, hasil)
		start++
		end++
	}

	fmt.Println("Sum tertinggi adalah:", slices.Max(max))
}
