package main

import "fmt"

func main(){
	
	person := map[string]string{
		"name": "Danil",
		"address": "Bintan",
	}

	// menambahkan key baru
	person["title"] = "programmer"

	// mengakses tipe data map
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	// function Map
	fmt.Println(len(person))

	// mengubah data di map dengan key
	person["name"] = "danil syah"
	fmt.Println(person["name"])

	// menghapus data di map dengan key
	delete(person, "title")
	fmt.Println(person)

	// membuat Map baru
	var book map[string]string = make(map[string]string)
	book["title"] = "Ayo Belajar Golang"
	book["author"] = "danil"
	book["kategori"] = "komputer"

	fmt.Println(book)
	fmt.Println(len(book))

	// hapus data map
	delete(book,"kategori")

	fmt.Println(book)
	fmt.Println(len(book))
}

