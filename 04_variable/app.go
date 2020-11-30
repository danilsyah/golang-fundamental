package main

import "fmt"

func main(){
	var name string
	var umur int
	
	name = "Danil syah arihardjo"
	umur = 26
	fmt.Println(name,"berumur",umur)

	name = "Haykal Dafiansyah"
	fmt.Println(name)

	// deklarasi variabel tanpa type data
	var firstName = "Danil"
	var lastName = "Syah"
	var age = 26
	fmt.Println(firstName, lastName, "berumur ",age)

	// deklarasi variable tanpa penulisan var menggunakan :=
	country := "Indonesia"
	fmt.Println(country)
	
	country = "Arab"
	fmt.Println(country)


	// deklarasi multiple variable
	var(
		namaAwal = "Haykal"
		namaAkhir = "Dafiansyah"
		berumur = 2
		isMarried = false
	)

	fmt.Println(namaAwal,namaAkhir,berumur,isMarried)

}