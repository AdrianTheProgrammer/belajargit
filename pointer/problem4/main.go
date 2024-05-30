package main

import "fmt"

func main() {
	var menu int
	var a = student{}
	var c Cipher = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student's Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncoded Student's Name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student's Encoded Name : ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecoded Student's Name " + a.nameEncode + " is : " + c.Decode())
	} else {
		fmt.Println("Wrong input!")
	}
}

type student struct {
	name       string
	nameEncode string
	//score      int
}

type Cipher interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = s.name
	nameRune := []rune{}

	for i, loop := range nameEncode {
		if nameEncode[i] != 32 {
			if nameEncode[i] >= 65 && nameEncode[i] <= 90 {
				if nameEncode[i]+3 > 90 {
					nameRune = append(nameRune, loop+3-26)
				} else if nameEncode[i]+3 < 65 {
					nameRune = append(nameRune, loop+3+26)
				} else {
					nameRune = append(nameRune, loop+3)
				}
			} else if nameEncode[i] >= 97 && nameEncode[i] <= 122 {
				if nameEncode[i]+3 > 122 {
					nameRune = append(nameRune, loop+3-26)
				} else if nameEncode[i]+3 < 97 {
					nameRune = append(nameRune, loop+3+26)
				} else {
					nameRune = append(nameRune, loop+3)
				}
			}
		} else {
			nameRune = append(nameRune, loop)
		}
	}

	return string(nameRune)
}

func (s *student) Decode() string {
	var nameDecode = s.nameEncode
	nameRune := []rune{}

	for i, loop := range nameDecode {
		if nameDecode[i] != 32 {
			if nameDecode[i] >= 65 && nameDecode[i] <= 90 {
				if nameDecode[i]-3 > 90 {
					nameRune = append(nameRune, loop-3-26)
				} else if nameDecode[i]-3 < 65 {
					nameRune = append(nameRune, loop-3+26)
				} else {
					nameRune = append(nameRune, loop-3)
				}
			} else if nameDecode[i] >= 97 && nameDecode[i] <= 122 {
				if nameDecode[i]-3 > 122 {
					nameRune = append(nameRune, loop-3-26)
				} else if nameDecode[i]-3 < 97 {
					nameRune = append(nameRune, loop-3+26)
				} else {
					nameRune = append(nameRune, loop-3)
				}
			}
		} else {
			nameRune = append(nameRune, loop)
		}
	}

	return string(nameRune)
}
