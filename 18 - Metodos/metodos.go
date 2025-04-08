package main

import "fmt"

type user struct {
	name  string
	age uint8
}

func (u user) userSave() {
	fmt.Printf("Saved user: %s\n", u.name)
}

func main() {
	newUser := user{
		name: "Lucas",
		age: 25,
	}
	newUser.userSave()
	fmt.Println(newUser)	
}