// Golang program to replace all specified
// substrings by given substring

package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "ABC,DEF,GHI,JKL"
	str2 := "PQR,ERT,TYU,YOU"

	fmt.Println(strings.ReplaceAll(str1, "ABC", "ABS"))
	fmt.Println(strings.ReplaceAll(str2, "PQR", "ABS"))
}
