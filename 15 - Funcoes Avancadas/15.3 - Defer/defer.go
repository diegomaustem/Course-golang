package main

import "fmt"

// Using defer

func ApprovedStudent(valueOne, valueTwo float32)bool{
	defer fmt.Println("Calculating the average...")
	average := (valueOne + valueTwo) / 2

	if average >= 7 {
		return true
	} 
	return false
}
func main() {
	fmt.Println(ApprovedStudent(2, 8))
}