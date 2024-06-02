package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B, C int

	fmt.Print("Input A: ")
	fmt.Scanln(&A)
	fmt.Print("Input B: ")
	fmt.Scanln(&B)
	fmt.Print("Input C: ")
	fmt.Scanln(&C)
	fmt.Println()

	map_pertambahan := pertambahan(A)
	map_perkalian := perkalian(B, map_pertambahan)
	map_perpangkatan := perpangkatan(C, map_perkalian)

	if len(map_pertambahan) == 0 || len(map_perkalian) == 0 || len(map_perpangkatan) == 0 {
		fmt.Println("No Output")
	} else {
		fmt.Println("Possible x, y, z values are :")
		for i := 0; i < len(map_perpangkatan); i++ {
			fmt.Println(map_perpangkatan[i])
		}
	}
}

func pertambahan(A int) map[int][]int {
	map_pertambahan := map[int][]int{}
	counter := 0

	for i := 0; i <= A; i++ {
		for j := 0; j <= A; j++ {
			for k := 0; k <= A; k++ {
				if i+j+k == A {
					map_pertambahan[counter] = append(map_pertambahan[counter], i, j, k)
					counter++
				}
			}
		}
	}

	return map_pertambahan
}

func perkalian(B int, map_pertambahan map[int][]int) map[int][]int {
	map_perkalian := map[int][]int{}
	counter := 0

	for i := 0; i < len(map_pertambahan); i++ {
		if map_pertambahan[i][0]*map_pertambahan[i][1]*map_pertambahan[i][2] == B {
			map_perkalian[counter] = append(map_perkalian[counter], map_pertambahan[i][0], map_pertambahan[i][1], map_pertambahan[i][2])
			counter++
		}
	}

	return map_perkalian
}

func perpangkatan(C int, map_perkalian map[int][]int) map[int][]int {
	map_perpangkatan := map[int][]int{}
	counter := 0

	for i := 0; i < len(map_perkalian); i++ {
		if math.Pow(float64(map_perkalian[i][0]), 2)+math.Pow(float64(map_perkalian[i][1]), 2)+math.Pow(float64(map_perkalian[i][2]), 2) == float64(C) {
			map_perpangkatan[counter] = append(map_perpangkatan[counter], map_perkalian[i][0], map_perkalian[i][1], map_perkalian[i][2])
			counter++
		}
	}

	return map_perpangkatan
}
