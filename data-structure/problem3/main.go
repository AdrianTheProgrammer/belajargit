package main

import (
	"fmt"
)

func main() {
	unique_array([]int{1, 2, 3, 4}, []int{1, 3, 5, 10, 16})
	unique_array([]int{10, 20, 30, 40}, []int{5, 10, 15, 59})
	unique_array([]int{1, 3, 7}, []int{1, 3, 5})
	unique_array([]int{3, 8}, []int{2, 8})
	unique_array([]int{1, 2, 3}, []int{3, 2, 1})
}

func unique_array(slice1, slice2 []int) {
	hasil := []int{}
	unique := true

	for i := 0; i < len(slice1); i++ {
		for j := 0; j < len(slice2); j++ {
			if i < len(slice1) {
				if slice1[i] != slice2[j] {
					unique = true
				} else {
					unique = false
					i++
				}
			}
		}

		if unique {
			hasil = append(hasil, slice1[i])
		}
	}

	fmt.Println("Unique Array:", hasil)
}
