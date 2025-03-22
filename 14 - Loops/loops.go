package main

import (
	"fmt"
	"time"
)

func main() {
	// Is normal :::
	number := 0

	for number < 10 {
		number++
		fmt.Println(number)
	}

	// Increment ::: 
	for j := 0; j < 10; j++ {
	 	fmt.Println("Increment",j)
	 	time.Sleep(time.Second)
	}

	// Using range :::
	names := [5]string{"Carlos", "Marcelo", "Pedro", "Larissa", "Leticia"}
	for _, name := range names {
		fmt.Println(name)
	}

	// Iterate for letter ::: 
	for _, letter := range "Beet" {
		fmt.Println(string(letter))
	}

	user := map[string]string{
		"name": "Daniel",
		"lastName": "Dantas",
	}

	for _, value := range user {
		fmt.Println(value)
	}
}