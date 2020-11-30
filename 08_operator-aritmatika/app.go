package main

import "fmt"

func main(){
	// operator matematika 
	a := 15
	b := 2
	
	var (
		penjumlahan int
		pengurangan int
		perkalian int32
		pembagian float32
		sisaBagi int
	)

	penjumlahan = a + b
	fmt.Println(penjumlahan)

	pengurangan = a - b
	fmt.Println(pengurangan)

	perkalian  = int32(a * b)
	fmt.Println(perkalian)

	pembagian = float32(a / b)
	fmt.Println(pembagian)

	sisaBagi = a % b
	fmt.Println(sisaBagi)

	fmt.Println("---------------------")
	// augmented assigments
	a += 10
	fmt.Println(a)

	a -= 10
	fmt.Println(a)

	a *= 2
	fmt.Println(a)

	a /= 3
	fmt.Println(a)

	a %= 6
	fmt.Println(a)

	// unary operator
	b++ // b = b + 1
	fmt.Println(b)
	b-- // b = b - 1
	fmt.Println(b)
	
	var nilaiNegative = -200
	var nilaiPositive = +200
	fmt.Println(nilaiNegative)
	fmt.Println(nilaiPositive)
	
	var nilaiKebalikan = !true
	fmt.Println(nilaiKebalikan)
}