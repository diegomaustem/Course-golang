package main

import "fmt"

func calcMath(numberOne, numberTwo int) (sum int, sub int) {
	sum = numberOne + numberTwo
	sub = numberOne - numberTwo
	return 
}

func main() {
	sum, sub := calcMath(10, 20)
	fmt.Println(sum, sub)
}