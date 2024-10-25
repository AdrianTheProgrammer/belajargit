package main

import "fmt"

type Product struct {
	Quantity int
	SubTotal uint64
}

func main() {
	product := [3]Product{}
	product[0] = Product{
		Quantity: 1,
		SubTotal: 4299000,
	}
	product[1] = Product{
		Quantity: 1,
		SubTotal: 3299000,
	}
	product[2] = Product{
		Quantity: 2,
		SubTotal: 21999000,
	}

	var grandTotal uint64

	for i, val := range product {
		product[i].SubTotal = val.SubTotal * uint64(val.Quantity)
		grandTotal += product[i].SubTotal
	}

	fmt.Println(product)
	fmt.Println(grandTotal)

	// var number Product
	// fmt.Println(number)
}
