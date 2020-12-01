package main

import "fmt"

func main(){
	var ujian = 90
	var absensi = 81

	var lulusNilaiAkhir bool = ujian > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi


	if (lulus) {
		fmt.Println("selamat anda lulus")
	}else{
		fmt.Println("anda tidak lulus")
	}

	var nilaiTugas = 75
	var nilaiPemrograman = 60

	if nilaiTugas >= 75 || nilaiPemrograman >= 75 {
		fmt.Println("anda bisa menambah sks")
	}else{
		fmt.Println("anda tidak bisa menambah sks")
	}
}