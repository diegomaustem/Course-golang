package main

import "fmt"

func main() {
	// Arrays :::
	var names [3] string
	names[0] = "Mariana"
	names[1] = "Carlos"
	names[2] = "Pedro"

	var ages = [...]int{18, 19, 21, 35, 49}

	fmt.Println(names)
	fmt.Println(ages[2])

	// Slices ::: 
	prices := []float32{1.25, 2,95, 759.99, 4.99, 6.77}
	fmt.Println(prices)
	prices = append(prices, 5.88)
	fmt.Println(prices)
}