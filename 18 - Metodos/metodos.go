package main

import "fmt"

type user struct {
	name  string
	age uint8
}

func (u user) userSave() {
	fmt.Printf("Saved user: %s\n", u.name)
}

func (u *user) haveABirthday() {
	u.age++
}

func (u user) ofAge() bool{
	if u.age >= 18 {
		return true
	}
	return false
}

func main() {
	newUser := user{
		name: "Lucas",
		age: 12,
	}
	newUser.userSave()
	newUser.haveABirthday()
	ofAge := newUser.ofAge()
	fmt.Println(newUser)	
	fmt.Println(ofAge)
}