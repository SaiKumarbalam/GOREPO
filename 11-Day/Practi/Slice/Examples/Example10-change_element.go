package main

import "fmt"

func main() {
	country := []string{"TX", "SAN", "TEST"}
	country[1] = "CAL"

	fmt.Println(country)
}
