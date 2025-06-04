package main

import (
	"fmt"
)

func isFactor(x, y int) bool {
	// Memeriksa apakah y habis dibagi x
	return y%x == 0
}

func main() {
	var x, y int

	// Input
	fmt.Print("nilai x: ")
	fmt.Scan(&x)
	fmt.Print("nilai y: ")
	fmt.Scan(&y)

	// Memeriksa apakah x adalah faktor dari y
	if isFactor(x, y) {
		fmt.Printf("%d adalah faktor dari %d\n", x, y)
	} else {
		fmt.Printf("%d bukanlah faktor dari %d\n", x, y)
	}
}
