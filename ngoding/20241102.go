package main

import (
	"fmt"
)

func main() {
	var u1, u2, u3 int
	fmt.Scan(&u1, &u2, &u3)

	if u1%2 != 0 && u2%2 != 0 && u3%2 != 0 {
		fmt.Println("ganjil")
	} else if u1%2 == 0 && u2%2 == 0 && u3%2 == 0 {
		fmt.Println("genap")
	} else {
		fmt.Println("bukan ganjil atau genap")
	}
}