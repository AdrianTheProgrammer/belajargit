package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var kata1 = bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan kata pertama: ")
	kata1.Scan()

	var kata2 = bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan kata kedua: ")
	kata2.Scan()

	compare_string(kata1, kata2)
}

func compare_string(kata1 *bufio.Scanner, kata2 *bufio.Scanner) {
	var hasil []string

	if len(kata1.Bytes()) > len(kata2.Bytes()) {
		for i := 0; i < len(kata2.Bytes()); i++ {
			if kata1.Text()[i] == kata2.Text()[i] {
				hasil = append(hasil, string(kata1.Text()[i]))
			}
		}
	} else {
		for i := 0; i < len(kata1.Bytes()); i++ {
			if kata1.Text()[i] == kata2.Text()[i] {
				hasil = append(hasil, string(kata1.Text()[i]))
			}
		}
	}

	fmt.Println("=====================================")
	fmt.Println("String yang sama adalah:", strings.Join(hasil, ""))
}
