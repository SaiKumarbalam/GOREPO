package main

import "fmt"

func main() {
	var a int = 58
	var b float32 = 65.36
	var c float32
	c = float32(a) + b
	fmt.Println(c)
}
