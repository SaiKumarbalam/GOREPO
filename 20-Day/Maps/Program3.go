package main

import "fmt"

func main() {
	employees := make(map[string]int, 5) //Initializing a map using the make() function
	employees["saikumar"] = 63
	employees["kumar"] = 33

	fmt.Println(employees)

}
