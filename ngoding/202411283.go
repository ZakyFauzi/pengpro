package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var coefficients []float64

	for i := 0; i < 3; i++ {
		scanner.Scan()
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Input tidak valid")
			return
		}
		coefficients = append(coefficients, value)
	}

	a := coefficients[0]

	if a > 0 {
		fmt.Println("terbuka ke atas")
	} else if a < 0 {
		fmt.Println("terbuka ke bawah")
	} else {
		fmt.Println("bukan fungsi kuadrat")
	}
}