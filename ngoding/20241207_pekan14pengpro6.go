package main

import "fmt"

func main() {
	var x, n int
	fmt.Scan(&x)
	fmt.Scan(&n)

	nemu := false
	for n > 0 {
		digit := n % 10
		if digit == x {
			nemu = true
			break
		}
		n /= 10
	}

	if nemu {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
