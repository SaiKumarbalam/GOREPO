//edge case for cal

package main

import "fmt"

func main() {
	addition(6, 5)
	subtraction(89, 63)
	multiplication(5, 9)
	division(3, 30)

}

func addition(n1 int, n2 int) {
	sum := n1 + n2
	fmt.Println("sum:", sum)
}
func subtraction(n3 int, n4 int) {
	sub := n3 - n4
	fmt.Println("sub:", sub)
}
func multiplication(n5 int, n6 int) {
	mul := n5 * n6
	fmt.Println("mul:", mul)
}
func division(n7 int, n8 int) {
	div := n7 / n8
	fmt.Println("div:", div)
}
