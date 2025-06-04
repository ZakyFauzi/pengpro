package main

import (
	"fmt"
)

func hitungPenduduk(pendudukAwal, kelahiran, imigrasi, kematian, emigrasi int) int {
	// Rumus jumlah penduduk akhir
	return pendudukAwal + kelahiran + imigrasi - kematian - emigrasi
}

func main() {
	var pendudukAwal, kelahiran, imigrasi, kematian, emigrasi int

	// Input
	fmt.Print("Jumlah penduduk awal: ")
	fmt.Scan(&pendudukAwal)
	fmt.Print("Jumlah kelahiran: ")
	fmt.Scan(&kelahiran)
	fmt.Print("Jumlah imigrasi: ")
	fmt.Scan(&imigrasi)
	fmt.Print("Jumlah kematian: ")
	fmt.Scan(&kematian)
	fmt.Print("Jumlah emigrasi: ")
	fmt.Scan(&emigrasi)

	// Hitung jumlah penduduk akhir
	pendudukAkhir := hitungPenduduk(pendudukAwal, kelahiran, imigrasi, kematian, emigrasi)

	// hasil
	fmt.Printf("Jumlah penduduk akhir: %d\n", pendudukAkhir)
}
