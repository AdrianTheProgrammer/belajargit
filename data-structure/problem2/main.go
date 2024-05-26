package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	kata := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan kata: ")
	kata.Scan()

	var geser int
	fmt.Print("Masukkan nilai pergeseran: ")
	fmt.Scanln(&geser)

	kata_bytes := kata.Bytes()
	enkripsi := []string{}

	caesar_cipher(geser, kata_bytes, enkripsi)
}

func caesar_cipher(geser int, kata_bytes []uint8, enkripsi []string) {
	var adjust int

	for i := 0; i < len(kata_bytes); i++ {
		if kata_bytes[i] != 32 {
			if kata_bytes[i] >= 65 && kata_bytes[i] <= 90 {
				if kata_bytes[i]+byte(geser%26) > 90 {
					adjust = -26
					enkripsi = append_slice(enkripsi, kata_bytes, i, geser, adjust)
				} else if kata_bytes[i]+byte(geser%26) < 65 {
					adjust = 26
					enkripsi = append_slice(enkripsi, kata_bytes, i, geser, adjust)
				} else {
					adjust = 0
					enkripsi = append_slice(enkripsi, kata_bytes, i, geser, adjust)
				}
			} else if kata_bytes[i] >= 97 && kata_bytes[i] <= 122 {
				if kata_bytes[i]+byte(geser%26) > 122 {
					adjust = -26
					enkripsi = append_slice(enkripsi, kata_bytes, i, geser, adjust)
				} else if kata_bytes[i]+byte(geser%26) < 97 {
					adjust = 26
					enkripsi = append_slice(enkripsi, kata_bytes, i, geser, adjust)
				} else {
					adjust = 0
					enkripsi = append_slice(enkripsi, kata_bytes, i, geser, adjust)
				}
			}
		} else {
			enkripsi = append(enkripsi, string(kata_bytes[i]))
		}
	}

	fmt.Println("Caesar Cipher:", strings.Join(enkripsi, ""))
}

func append_slice(enkripsi []string, kata_bytes []uint8, i int, geser int, adjust int) []string {
	enkripsi = append(enkripsi, string(kata_bytes[i]+byte(geser%26)+byte(adjust)))
	return enkripsi
}
