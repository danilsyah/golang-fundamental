package main

import "fmt"

// function tanpa parameter
func sayHello(){
	fmt.Println("Hello danil")
}

// function menggunakan parameter
func sayGretting(name string, umur int8) {
	fmt.Println("salam kenal", name, "berumur", umur)
}

// function parameter dan mengembalikan value
func getHello(name string)string{
	result := "hello" + name
	return result
}

func getName(name string)string{
	if name == ""{
		return "Hello Bro"
	}else{
		return "Hello" + name
	}
}

// function dengan returning multiple values
func getFullName(firstName string, middleName string, lastName string)(string, string, string){
	return firstName, middleName, lastName
}


func main(){
	sayHello()
	sayGretting("haykal", 26)
	fmt.Println(getHello("danil"))
	fmt.Println(getName(""))

	firstName, middleName, _ := getFullName("danil","syah","Ganteng")
	fmt.Println(firstName)
	fmt.Println(middleName)
}

