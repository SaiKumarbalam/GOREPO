package main

import "fmt"

func main() {

	primenumbers := []int{4, 5, 6}
	numbers := []int{1, 2, 3}

	copy(primenumbers, numbers)
	fmt.Println("primenumbers:", primenumbers)

}
