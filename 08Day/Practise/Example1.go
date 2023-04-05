package main

import "fmt"

func main() {
	for count := 1; count <= 10; count++ {
		if count == 5 {
			continue
		}
		fmt.Println(count)
	}
}
