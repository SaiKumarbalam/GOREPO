//create a new slice from the existing slice

package main

import "fmt"

func main() {
	num := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	new := num[1:7]
	arr := new[2:5]

	fmt.Println("New:", new)
	fmt.Println("Array", arr)

}
