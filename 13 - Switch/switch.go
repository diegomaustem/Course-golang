package main

import "fmt"

func dayOfWeek(number int)string {
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Inválid number"
	}
}

func main() {
	day := dayOfWeek(6);
	fmt.Println("Today is", day)
}