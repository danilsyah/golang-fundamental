package main

import "fmt"

func main(){
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8) //akan menghasilkan minus karena int8 batas nilai max  127

	var name = "danil"
	var d = name[1]
	var dString = string(d)

	fmt.Println(dString)
}