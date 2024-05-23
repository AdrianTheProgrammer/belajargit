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
	fmt.Print("Masukkan nilai pergeseran (Pos=R, Neg=L): ")
	fmt.Scanln(&geser)

	enkripsi := []string{}
	kata_bytes := kata.Bytes()

	for i := 0; i < len(kata_bytes); i++ {
		if kata_bytes[i] != 32 {
			if kata_bytes[i] >= 65 && kata_bytes[i] <= 90 {
				if kata_bytes[i]+byte(geser) > 90 {
					enkripsi = append(enkripsi, string(kata_bytes[i]+byte(geser)-26))
				} else if kata_bytes[i]+byte(geser) < 65 {
					enkripsi = append(enkripsi, string(kata_bytes[i]+byte(geser)+26))
				} else {
					enkripsi = append(enkripsi, string(kata_bytes[i]+byte(geser)))
				}
			} else if kata_bytes[i] >= 97 && kata_bytes[i] <= 122 {
				if kata_bytes[i]+byte(geser) > 122 {
					enkripsi = append(enkripsi, string(kata_bytes[i]+byte(geser)-26))
				} else if kata_bytes[i]+byte(geser) < 97 {
					enkripsi = append(enkripsi, string(kata_bytes[i]+byte(geser)+26))
				} else {
					enkripsi = append(enkripsi, string(kata_bytes[i]+byte(geser)))
				}
			}
		} else {
			enkripsi = append(enkripsi, string(kata_bytes[i]))
		}
	}

	fmt.Println(strings.Join(enkripsi, ""))
}
