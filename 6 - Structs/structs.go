package main

import "fmt"

type people struct {
	name string 
	age uint8
}

func main() {
	// The first mod of declaration 
	var user people
	user.name = "John"
	user.age = 31
	fmt.Println(user)

	// The seccond mod of declaration 
	employer := people{"Carlos", 25}
	fmt.Println(employer)

	// The third mod of declaration 
	client := people{name: "Rafael"}
	fmt.Println(client)
}