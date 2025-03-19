package main

import "fmt"

func main() {
	// Arrays Internos :::
	slice := make([]float32, 10, 11)
	slice = append(slice, 8)
	slice = append(slice, 9)
	slice = append(slice, 10)

	// No slice caso extrapole o limite do array é criado mais posições ::: 
	fmt.Println(slice)
	fmt.Println(cap(slice))
}