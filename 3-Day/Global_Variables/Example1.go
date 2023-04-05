package main

import (
	"fmt"
)

var (
	firstName string
	lastName  string
	rollNo    int
	gender    string
)

func main() {

	firstName = "Bala" // create
	lastName = "Stas"
	rollNo = 25
	gender = "male"

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(rollNo)
	fmt.Println(gender)

}
