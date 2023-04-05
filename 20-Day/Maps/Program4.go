package main

import "fmt"

func main() {
	students := make(map[string]int)
	students["bandi"] = 23
	students["gundi"] = 23
	students["mundi"] = 23
	students["tundi"] = 23
	fmt.Println(students)
}
