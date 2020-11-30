package main

import "fmt"

func main(){
	// variabel const tidak dapat di ubah value nya (tetap)
	const firstName string = "danil"
	const lastName = "syah"
	const age = 26

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)

	// multiple constant
	const(
		nik int = 3216607
		fullName = "danil syah"
	)
	println(nik)
	println(fullName)
}