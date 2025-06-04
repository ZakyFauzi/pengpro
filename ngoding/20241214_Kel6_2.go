package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n > 0 {
		fmt.Print("Positif, ")
		pembagiTerbesar := 1
		for i := 2; i <= n/2; i++ {
			if n%i == 0 {
				pembagiTerbesar = i
			}
		}
		fmt.Print(pembagiTerbesar)
	} else if n < 0 {
		fmt.Print("Negatif, ")
		n := -n
		pembagiTerbesar := 1
		for i := 2; i <= n/2; i++ {
			if n%i == 0 {
				pembagiTerbesar = i
			}
		}
		fmt.Print(pembagiTerbesar)
	} else {
		fmt.Print("Nol, Tidak ada pembagi")
	}
}
