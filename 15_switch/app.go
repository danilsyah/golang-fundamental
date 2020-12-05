package main

import "fmt"

func main()  {
	bulan := 14

	switch bulan {
		case 1:
			fmt.Println("Januari")
		case 2:
			fmt.Println("Februari")
		case 3:
			fmt.Println("Maret")
		case 4:
			fmt.Println("April")
		case 5:
			fmt.Println("May")
		case 6:
			fmt.Println("Juni")
		case 7:
			fmt.Println("Juli")
		case 8:
			fmt.Println("August")
		case 9:
			fmt.Println("September")
		case 10:
			fmt.Println("October")
		case 11:
			fmt.Println("November")
		case 12:
			fmt.Println("December")
		default:
			fmt.Println("bulan tidak ditemukan")
	}

	// switch short statement
	name := "udin"
	switch length := len(name); length > 5{
		case true:
			fmt.Println("nama terlalu panjang ")
		case false:
			fmt.Println("nama sudah benar")
	}

	// switch dengan kondisi didalam case
	length := len(name)
	switch{
	case length > 5:
		fmt.Println("Nama terlalu panjang")
	case length > 3:
		fmt.Println("Nama kependekan")
	default:
		fmt.Println("nama sudah benar")
	}
}