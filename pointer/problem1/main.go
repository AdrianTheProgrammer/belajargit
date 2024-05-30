package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Print("Masukkan A: ")
	fmt.Scanln(&a)

	fmt.Print("Masukkan B: ")
	fmt.Scanln(&b)

	swap_variable(a, b)
}

func swap_variable(a, b int) {
	var bantuan int
	var pointer_a *int = &a
	var pointer_b *int = &b

	bantuan = *pointer_a
	*pointer_a = *pointer_b
	*pointer_b = bantuan

	fmt.Println("A:", a)
	fmt.Println("B:", b)
}
