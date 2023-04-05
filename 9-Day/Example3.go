package main

import "fmt"

func main() {
	a := 42

	if a%2 == 0 && a >= 10 {
		fmt.Println("hello")
	} else {
		fmt.Println("Dead")
	}
}
