package main

import "fmt"

func main() {
	name := [3]string{"test", "rest", "must"}

	name[1] = "kest"

	fmt.Println("Name:", name)
}
