package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var palindrome bool
	kata := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan kata: ")
	kata.Scan()

	katatext := kata.Text()
	var bantuan2 int = len(katatext)

	for bantuan := 0; bantuan <= (len(katatext)/2)-1; bantuan++ {
		if katatext[bantuan] == katatext[bantuan2-1] {
			palindrome = true
		} else {
			palindrome = false
			break
		}
		bantuan2--
	}

	if palindrome {
		fmt.Println("Kata", katatext, "ADALAH sebuah Palindrome")
	} else {
		fmt.Println("Kata", katatext, "BUKANLAH sebuah Palindrome")
	}
}
