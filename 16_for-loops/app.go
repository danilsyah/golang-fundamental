package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("data ke ", counter)
	}

	// membuat sebuah data slice
	names := []string{"danil","haykal","razil","fika","icih"}
	// iterasi data slice
	for i := 0; i < len(names); i++ {
		fmt.Println("nama saya",names[i])
	}

	// for range
	for index, name := range names {
		fmt.Println("Index ke-",index+1,"=",name)
	}

	// gunakan _ untuk variable yang tidak di pakai
	for _, name := range names {
		fmt.Println(name)
	}

	// iterasi tipe data map
	person := make(map[string]string)
	person["name"] = "Danil"
	person["title"] = "Programmer"
	
	for key, value := range person {
		fmt.Println(key, "=",value)
	}
}