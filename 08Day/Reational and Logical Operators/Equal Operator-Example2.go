package main

import "fmt"

func main() {
	num1 := 20
	num2 := 12

	var result bool

	// Equal Operator
	result = (num1 == num2)
	fmt.Printf("%d == %d returns %t \n", num1, num2, result)

}
