package main

import "fmt"

func main(){

	// deklarasi variabel array dengan menentukan jumlah index nya
	var names[3] string

	names[0] = "danil"
	names[1] = "haykal"
	names[2] = "razil"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// deklarasi array secara langsung
	var values = [5] int{
		90,
		88,
		32,
	}

	// cara meng akses array
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	// fungsi mendapatkan jumlah array
	fmt.Println(len(names))
	fmt.Println(len(values))
	// cara merubah nilai array berdasarkan index
	values[1] = 100
	fmt.Println(values[1])

	// membuat array tanpa menentukan jumlah kapasitas datanya
	var months = [...] string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	fmt.Println(months)
}