package main

import "fmt"

func main() {
	dayOfweek := 5

	switch dayOfweek {
	case 1:
		fmt.Println("sunday")

	case 2:
		fmt.Println("monday")

	case 3:
		fmt.Println("tuesday")

	case 4:
		fmt.Println("wensday")

	case 5:
		fmt.Println("Thursday")

	default:
		fmt.Println("no day")
	}
}
