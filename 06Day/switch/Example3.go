// Multiple case rear case

package main

import "fmt"

func main() {
	dayOfweek := "monday"

	switch dayOfweek {
	case "Saturday", "sunday":
		fmt.Println("Weekned")

	case "monday", "tuesday", "wedsnday", "thurdsday":
		fmt.Println("weekdays") //output

	default:
		fmt.Println("no day")

	}
}
