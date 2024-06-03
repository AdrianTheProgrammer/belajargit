package main

import "fmt"

func main() {
	fmt.Println("Pecahan uang kembalian:", kembalian(123))
	fmt.Println("Pecahan uang kembalian:", kembalian(432))
	fmt.Println("Pecahan uang kembalian:", kembalian(543))
	fmt.Println("Pecahan uang kembalian:", kembalian(7752))
	fmt.Println("Pecahan uang kembalian:", kembalian(15321))
}

func kembalian(uang int) []int {
	pecahan := []int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10, 1}
	hasil := []int{}

	for _, loop := range pecahan {
		for uang >= loop {
			uang -= loop
			hasil = append(hasil, loop)
		}
	}

	return hasil
}
