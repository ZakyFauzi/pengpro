package main

import "fmt"

func main() {
	var hari string
	fmt.Scanln(&hari)

	if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" || hari == "Jumat" {
		fmt.Println("Hari kerja")
	} else if hari == "Sabtu" || hari == "Minggu" {
		fmt.Println("Akhir pekan")
	} else {
		fmt.Println("Masukan tidak valid")
	}
}
