package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int
	fmt.Scan(&number)

	if number < 1000 || number > 9999 {
		fmt.Println("Bilangan harus antara 1000 hingga 9999.")
		return
	}

	numStr := strconv.Itoa(number)
	firstDigit := int(numStr[0] - '0')
	lastDigit := int(numStr[3] - '0')

	if firstDigit%2 != 0 && lastDigit%2 == 0 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
