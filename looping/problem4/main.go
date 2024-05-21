package main

import "fmt"

func main() {
	var kata string
	var palindrome bool

	fmt.Print("Masukkan kata: ")
	fmt.Scanln(&kata)

	var bantuan2 int = len(kata)

	for bantuan := 0; bantuan <= (len(kata)/2)-1; bantuan++ {
		if kata[bantuan] == kata[bantuan2-1] {
			palindrome = true
		} else {
			palindrome = false
			break
		}
		bantuan2--
	}

	if palindrome == true {
		fmt.Println("Kata", kata, "adalah sebuah Palindrome")
	} else {
		fmt.Println("Kata", kata, "bukanlah sebuah Palindrome")
	}
}
