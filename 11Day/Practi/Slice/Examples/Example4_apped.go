package main

import "fmt"

func main() {

	evennumbers := []int{2, 4}
	oddnumbers := []int{1, 3}

	numbers := append(evennumbers, oddnumbers...)

	fmt.Println("Numbers:", numbers)
}
