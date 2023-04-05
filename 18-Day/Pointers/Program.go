package main

import (
	"fmt"
)

func main() {
	b := 66
	var a *int = &b
	fmt.Println("type of a is %T \n", a)
	fmt.Println("address of b", a)
}
