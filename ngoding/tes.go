package main

import "fmt"

func f(x int, y int) float64 {
	// Menghitung nilai f(x, y)
	result := 1.0/(3*float64(x)*float64(x)+10) + 10*float64(y) + 7
	return result
}

func main() {
	// Input dari pengguna
	var x, y int
	fmt.Print("Inputan untuk nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Inputan untuk nilai y: ")
	fmt.Scan(&y)

	// Menghitung hasil
	result := f(x, y)

	//Menampilkan hasil
	fmt.Printf("nilai f(%d, %d) = %.6f\n", x, y, result)
}
