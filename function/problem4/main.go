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
	for i := 0; i < len(kata_bytes); i++ {
		if kata_bytes[i] != 32 {
			if kata_bytes[i] >= 65 && kata_bytes[i] <= 90 {
				if kata_bytes[i]+byte(geser%26) > 90 {
					enkripsi = append(enkripsi, string(kata_bytes[i]+byte(geser%26)-26))
				} else if kata_bytes[i]+byte(geser%26) < 65 {
					enkripsi = append(enkripsi, string(kata_bytes[i]+byte(geser%26)+26))
				} else {
					enkripsi = append(enkripsi, string(kata_bytes[i]+byte(geser%26)))
				}
			} else if kata_bytes[i] >= 97 && kata_bytes[i] <= 122 {
				if kata_bytes[i]+byte(geser%26) > 122 {
					enkripsi = append(enkripsi, string(kata_bytes[i]+byte(geser%26)-26))
				} else if kata_bytes[i]+byte(geser%26) < 97 {
					enkripsi = append(enkripsi, string(kata_bytes[i]+byte(geser%26)+26))
				} else {
					enkripsi = append(enkripsi, string(kata_bytes[i]+byte(geser%26)))
				}
			}
		} else {
			enkripsi = append(enkripsi, string(kata_bytes[i]))
		}
	}

	fmt.Println("Caesar Cipher:", strings.Join(enkripsi, ""))
}
