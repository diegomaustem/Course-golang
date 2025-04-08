package main

import "fmt"

var number int 

func init() {
	number = 10
	fmt.Println("Init function executed")
}

func main() {
	fmt.Println("Main function executed")
	fmt.Println(number)
}