// Golang program to convert the specified string
// into title case

package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "hello how are you "
	str2 := " i am fine"

	fmt.Println(strings.Title(str1))
	fmt.Println(strings.Title(str2))
}
