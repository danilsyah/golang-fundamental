// slice adalah type data potongan dari sebuah array,  slice selalu mengacu atau mereferensi ke array
package main

import "fmt"

func main(){
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

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println("length = ",len(slice1))
	fmt.Println("capacity = ", cap(slice1))

	// ketika ada data array yang di ubah maka data di slice nya pun ikut terubah
	months[5] = "juni-di-ubah"
	fmt.Println(slice1)
	fmt.Println(months)

// merubah data di slice, maka data di array pun ikut ter ubah
	slice1[0] = "Mei-diganti"
	fmt.Println(slice1)
	fmt.Println(months)

	// method append slice
	days := [...] string{"senin","selasa","rabu","kamis","jumat","sabtu","minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur baru")
	daySlice2[0] = "Ups"
	fmt.Println(daySlice2)
	fmt.Println(days)

	// membuat slice tanpa array
	newSlice := make([]string, 3, 5)
	newSlice[0] = "danil"
	newSlice[1] = "syah"
	newSlice[2] = "arihardjo"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// hati-hati dalam membuat array atau slice
	iniArray := [5]int{1,2,3,4,5} //harus di tentukan panjang nilainya atau di isi [...]
	iniSlice := []int{1,2,3,4,5} // kurawal di kosongkan

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}