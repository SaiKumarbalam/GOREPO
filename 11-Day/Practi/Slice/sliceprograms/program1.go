//create a slice from an integer array

package main

import "fmt"

func main() {

	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	index := arr[2:5]

	fmt.Println("Array:", index)
}
