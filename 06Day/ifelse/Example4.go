package main

import "fmt"

func main() {
	num1 := 20
	num2 := 13
	if num1 == num2 { // Equal to Operator
		fmt.Println("this  are eqaul")
	} else if num1 > num2 { // greater
		fmt.Println("this is  greater")
	} else if num1 < num2 { // lesser than
		fmt.Println("This is  no more needed")
	} else if num1 <= num2 {
		fmt.Println("This is  no more needed")
	} else if num1 >= num2 {
		fmt.Println("This is  no more needed")
	}
}
