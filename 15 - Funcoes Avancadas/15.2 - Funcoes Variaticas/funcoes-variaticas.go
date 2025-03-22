package main

import "fmt"

func sum(numbers ...int) int {
	amount := 0
	for _, number := range numbers {
		amount += number
	}

	return amount
}

func main() {
	amountSum := sum(1, 3, 5, 9)
	fmt.Println(amountSum)
}