package main

import "fmt"

func main() {
	for i := 1; i < 5; i++ {
		for j := 5; j > 0; j-- {
			if j == i {
				fmt.Printf("Outer loop %d : Inner loop :%d \n", i, j)
				break
			}
		}
	}
	fmt.Println("Done")
}
