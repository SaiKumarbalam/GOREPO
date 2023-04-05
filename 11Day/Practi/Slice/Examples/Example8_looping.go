package main

import "fmt"

func main() {
	number := []int{2, 4, 6, 8, 10}

	for i := 0; i < len(number); i++ {
		fmt.Println(number[i])
	}
}
