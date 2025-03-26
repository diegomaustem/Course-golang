package main

import "fmt"


func main() {
	
	returnTxt := func(txt string) string {
		return (txt)
	}("This is string type")

	fmt.Println(returnTxt)
}