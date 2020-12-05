package main

import "fmt"

func main(){
	var name = "danilsyah"

	if name == "danilsyah" {
		fmt.Println("hello danil")
	}else if name == "Udin"{
		fmt.Println("ini udin")
	}else if name == "Icih"{
		fmt.Println("ini icih")
	}else{
		fmt.Println("tidak ada yang sama")
	}

	// short statement
	if length := len(name); length > 5{
		fmt.Println("nama terlalu panjang")
	}else{
		fmt.Println("nama sudah benar")
	}
}