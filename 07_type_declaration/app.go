package main

import "fmt"

func main(){
	// membuat tipe data alias sendiri
	type noKtp string
	type married bool

	var noKtpDanil noKtp = "2100110201201"
	var isMarried married = true
	fmt.Println(noKtpDanil)
	fmt.Println(isMarried)
}