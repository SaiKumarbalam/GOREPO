package main

import "fmt"

func main() {

	for i := 0; i <= 100; i++ {
		if i == 5 && i == 10 {
			continue

		}
		fmt.Println(i)
	}
}
