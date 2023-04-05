package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 	a := strconv.Itoa(12)
	// 	fmt.Println(a)

	// var name = "34"
	// test, err := strconv.Atoi(name)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%T", test)
	// }

	var firstName = "65"
	var lastName = "56"
	test, err := strconv.Atoi(firstName) // comma ok syntax
	if err != nil {
		fmt.Println(err)
	}
	rest, err := strconv.Atoi(lastName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(test)
	fmt.Println(rest)

}
