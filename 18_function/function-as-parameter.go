package main

import "fmt"

// function type declarations / alias deklarasi function
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter){
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string)string{
	if name == "anjing"{
		return "..."
	}else{
		return name
	}
}

func main() {
	sayHelloWithFilter("danil", spamFilter)
}