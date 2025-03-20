package main

import (
	"fmt"
)

func main() {
	number := 15 

	if number < 10 {
		fmt.Println("Smaller")
	} else {
		fmt.Println("Greatest number")
	}

	if numberReserv := number; numberReserv > 0 {
		fmt.Println("This number is greater")
	} else {
		fmt.Println("This number is smaller")
	}
}