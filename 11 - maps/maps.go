package main

import "fmt"

func main() {
	// Maps :::
	ages := map[int]int{
		1: 25,
		2: 34,
	}

	names := map[string]string{
		"First name": "Diego",
		"Last name": "Librand",
	}

	// Aninhamento de maps ::: 
	developers := map[string]map[string]string{
		"name": {
			"First name": "Rafael",
			"Last name":  "Corn",
		},
		"address": {
			"Street":     "B Street",
			"Complement": "Street house",
		},
	}

	// Deletando chave ::: 
	delete(developers, "address")

	// Adicionando chave ::: 
	developers["sign"] = map[string]string {
		"sign": "Librand",
	}

	fmt.Println(ages)
	fmt.Println(names)
	fmt.Println(developers)
}