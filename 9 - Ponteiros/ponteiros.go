package main

import "fmt"

func main() {
	var value = 100 

	// &  reference for address memory of variable
	var pointer = &value

	fmt.Println(value)
	fmt.Println(pointer)
	
	// Access the value of variable 
	fmt.Println(*pointer)
}