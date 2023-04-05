package main

import "fmt"

func main() {
	a := 12.36
	b := 69
	// var c int
	// c = int(a) + int(b)
	var c float32
	c = float32(a) + float32(b)

	fmt.Println(c)
}
