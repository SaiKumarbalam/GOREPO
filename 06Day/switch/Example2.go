package main

import "fmt"

func main() {
	dayOfweek := 3

	switch dayOfweek {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednsday")
		fallthrough
	case 4:
		fmt.Println("thrusday")
		fallthrough
	case 5:
		fmt.Println("friday")
	default:
		fmt.Println("No day")
	}

}
