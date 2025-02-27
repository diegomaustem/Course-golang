package main

import "fmt"

// Is normal funcion
func sum(numberOne int8, numberTwo int8) int8 {
	return numberOne + numberTwo
}

// Function with many returns 
func calculator(numberOne, numberTwo int8) (int8, int8) {
	sum := numberOne + numberTwo
	sub := numberOne - numberTwo
	return sum, sub
}

func main() {
	sum := sum(30, 28)
	fmt.Println(sum)

	var f = func(input string) string {
		fmt.Println(input)
		return input
	}

	result := f("Whats your movie?")
	fmt.Println(result)

	// Exec fun calculator 
	resultSum, resultSub := calculator(20, 30)
	fmt.Println(resultSum, resultSub)
}