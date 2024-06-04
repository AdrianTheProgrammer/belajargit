package main

import (
	"fmt"
)

func main() {
	fmt.Println("Index:", binary_search([]int{1, 1, 3, 5, 5, 6, 7}, 3))
	fmt.Println("Index:", binary_search([]int{1, 2, 3, 5, 6, 8, 10}, 5))
	fmt.Println("Index:", binary_search([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53))
	fmt.Println("Index:", binary_search([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100))
}

func binary_search(data []int, target int) int {
	beg := 0
	end := len(data) - 1

	for beg <= end {
		mid := (beg + end) / 2

		if data[mid] == target {
			return mid
		} else if data[beg] == target {
			return beg
		} else if data[end] == target {
			return end
		} else if target < data[0] || target > data[len(data)-1] {
			return -1
		} else if data[mid] < target {
			beg = mid + 1
			end--
		} else {
			end = mid - 1
			beg++
		}
	}

	return -1
}
