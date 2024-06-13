package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := [][]int32{{52, 80}, {34, 84}}
	arrlist := make(map[int][]int32)
	//var arrunique []int32
	// var unique int32
	arrconv := make(map[int][]string)

	for i := 0; i < len(arr); i++ {
		for j := arr[i][0]; j <= arr[i][1]; j++ {
			arrlist[i] = append(arrlist[i], j)
		}
	}

	fmt.Println(arrlist)

	for i := 0; i < len(arrlist); i++ {
		for j := 0; j < len(arrlist[i]); j++ {
			intconv := strconv.Itoa(int(arrlist[i][j]))
			arrconv[i] = append(arrconv[i], intconv)
		}
	}
	unique := true

	for i := 0; i < len(arrconv); i++ {
		for j := 0; j < len(arrconv[i]); j++ {
			// fmt.Println(arrconv[i][j])
			for k := 0; k < len(arrconv[i][j]); k++ {
				start := 1

				for l := start; l < len(arrconv[i][j]) && unique; l++ {
					fmt.Println(string(arrconv[i][j][k]), "=", string(arrconv[i][j][l]))
					if arrconv[i][j][k] == arrconv[i][j][l] {
						unique = false
					}
				}

				start++
			}
		}
	}

	// fmt.Println(arrconv)

	// bytes := "12344"
	// start := 1
	// //unique := true

	// for i := 0; i < len(bytes) && unique; i++ {

	// 	for j := start; j < len(bytes) && unique; j++ {
	// 		fmt.Println(string(bytes[i]), "=", string(bytes[j]))
	// 		if bytes[i] == bytes[j] {
	// 			unique = false
	// 		}

	// 	}
	// 	start++
	// }

	// fmt.Println(bytes, unique)
}
