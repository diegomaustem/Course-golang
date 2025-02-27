package main

import "fmt"

func main() {
	var name string = "Theo"
	lastname := "Grejoy"

	var ( 
		age int = 28
		color string = "Blue"
 	)

	address, street := "Kobrasol-SC", "Know Black Street"

	const identificationSerial string = "455kht-647455kodd-83lkd"

	fmt.Println("My name is:", name)
	fmt.Println("My lastname is:", lastname)

	fmt.Println("My age today is:", age)
	fmt.Println("My favorite color is:", color)

	fmt.Println("My address is:", address)
	fmt.Println("My street is:", street)

	fmt.Println("My serial:", identificationSerial)
}