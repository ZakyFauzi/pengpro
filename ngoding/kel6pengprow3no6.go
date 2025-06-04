package main

import "fmt"

func dapatCashback(A int) bool {
	// memeriksa digit pertama dengan membagi A dengan 1000
	digitPertama := A / 1000

	// memeriksa digit terakhir dengan modulus 10
	digitTerakhir := A % 10

	// memeriksa apakah digit pertama sama dengan digit terakhir
	if digitPertama == digitTerakhir {
		return true
	}
	return false
}

func main() {
	// input
	var A int
	fmt.Println("masukkan nomor undian (4 digit):")
	fmt.Scan(&A)

	// memeriksa cashback
	if dapatCashback(A) {
		fmt.Println("kamu mendapat cashback")
	} else {
		fmt.Println("kamu tidak mendapat cashback")
	}
}
