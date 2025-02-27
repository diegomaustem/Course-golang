package main

import (
	"errors"
	"fmt"
)

func main (){
	// Type Int 
	var number int64 = -10000000000
	fmt.Println(number)

	var numberTwo uint32 = 10000
	fmt.Println(numberTwo)

	// Numbers with alias ::: 

	// INT32 = RUNE
	var numberThree rune = 126584
	fmt.Println(numberThree)

	// BYTE = UINT8
	var numberFour byte = 123
	fmt.Println(numberFour)
	
	// Type Float :::

	var numberReal float32 = 1254.55
	fmt.Println(numberReal)

	numberRealTwo := 1254.36
	fmt.Println(numberRealTwo)

	// Type String
	var name string = "Theo"
	fmt.Println(name)

	lastName := "Lill"
	fmt.Println(lastName) 

	// Type boolean 
	var value bool
	fmt.Println(value)

	// Type error 
	var erro error 
	fmt.Println(erro)

	var internalError error = errors.New("Internal Error")
	fmt.Println(internalError)
}