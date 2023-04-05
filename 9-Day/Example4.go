package main

import "fmt"

func main() {
	var a = 35
	// greater than 0
	//even and odd  sastfisy
	//not divisble by 7

	if a >= 0 && a%2 == 0 || a%2 != 0 && a%7 != 0 {
		fmt.Println("test")
	} else {
		fmt.Println("rest")
	}
}
