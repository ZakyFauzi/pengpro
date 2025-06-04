package main

import "fmt"

func main() {
	var C float64

	fmt.Scanln(&C)

	result := (9.0/5.0)*C + 32

	fmt.Print(result)
}
