package main

import "fmt"

func main() {

	numbers := []int{2, 3, 4}

	numbers = append(numbers, 5, 6)
	fmt.Println("Numbers:", numbers)
}
