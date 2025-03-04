package main

import "fmt"

type people struct {
	name string
	lastName string  
	age uint8
	height uint8
}

type student struct {
	people 
	course string 
	register string 
}

func main() {
	firstPeople := people {"Lucas", "Law", 25, 185}
	fmt.Println(firstPeople)

	firstStudent := student{firstPeople, "Engineer", "025488"}
	fmt.Println(firstStudent.lastName)
}