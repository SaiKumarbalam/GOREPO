package main

import "fmt"

func main() {
	var num int = 0
	fmt.Print("Enter a Number:")
	fmt.Scanf("%d", &num)

	for count := 1; count <= 10; count++ {
		fmt.Printf(" %d * %d =%d \n", num, count, num*count)
	}
}
