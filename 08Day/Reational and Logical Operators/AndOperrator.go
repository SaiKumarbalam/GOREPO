package main

import "fmt"

func main() {
	num1 := 23
	num2 := 23
	num3 := 55

	var result bool

	result = (num1 == num2) && (num3 == num1) && (num2 == num3)

	fmt.Println("RESULT of AND opertator %t", result)
}
