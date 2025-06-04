package main

import (
	"fmt"
)

func main() {
	var hargaMakanan, hargaMinuman int
	var beriTip bool

	fmt.Scan(&hargaMakanan)

	fmt.Scan(&hargaMinuman)

	fmt.Scan(&beriTip)

	total := hargaMakanan + hargaMinuman
	if beriTip {
		total += 5000
	}

	fmt.Println(total)
}