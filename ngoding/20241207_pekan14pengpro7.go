package main

import "fmt"

func main() {
	var tipe string
	var A, B, C int

	fmt.Scan(&tipe)
	A = 0
	B = 0
	C = 0
	for _, char := range tipe {
		if char == 'A' {
			A++
		} else if char == 'B' {
			B++
		} else if char == 'C' {
			C++
		}
	}

	fmt.Println("Tipe A = ", A)
	fmt.Println("Tipe B = ", B)
	fmt.Println("Tipe C = ", C)
}
