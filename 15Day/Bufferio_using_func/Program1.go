package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	reader := bu`fio.NewReader(os.Stdin)
	num1, _ := reader.ReadString('\n')

	num2, _ := reader.ReadString('\n') // strconv(int)// strconv(int)

	n1, _ := strconv.ParseInt(num1, 6, 12)
	n2, _ := strconv.ParseInt(num2, 6, 12)
	// fmt.Println(strconv.Atoi(num2)+strconv.Atoi(num2))
	fmt.Println(n1)

	// strconv.Atoi(num2)
	//function call
	add(n1, n2)
	sub(n1, n2)
	mul(n1, n2)

	fmt.Println(n1)
	// div(n1, n2)

	fmt.Println("Addition:", add(n1, n2))
	fmt.Println("Subtraction:", sub(n1, n2))
	fmt.Println("Multiplication:", mul(n1, n2))
	// fmt.Println("Division:", div(n1, n2))

}

func add(num1 int64, num2 int64) int64 { //Fun

	return num1 + num2

}
func sub(num1 int64, num2 int64) int64 {
	return num1 - num2

}
func mul(num1 int64, num2 int64) int64 {
	return num1 * num2

}

// func div(num1 int, num2 int) int {
// 	if num1/num2 == 0 {
// 		return 0
// 	}
// 	return num1 / num2

// }
