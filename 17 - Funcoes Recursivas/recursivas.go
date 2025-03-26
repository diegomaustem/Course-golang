package main

import "fmt"

func fibonacci(position uint) int {
	if position == 0 {
		return 0
	} else if position == 1 {
		return 1
	}
	return fibonacci(position-1) + fibonacci(position-2)
}
func main() {
	position := uint(8)
	fmt.Println(fibonacci(position))
}