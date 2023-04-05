package main

import "fmt"

var (
	read  string
	write string
	count int
)

var c float32

var (
	name   string
	marks  float32
	result bool
)

func main() {

	greet()
	treat()
	meet()
}

func greet() {

	//the variable name (read)
	// the short variable declaration assignment (:=)
	// the value that is being tied to the variable name (book)
	// the data type inferred by Go (string)
	read := "book"
	write := "notes"
	count = 50
	fmt.Println(read)
	fmt.Println(write)
	fmt.Println(count)

}

func treat() {
	c = 36.55
	fmt.Println(c)
}
func meet() {
	name := "randy"
	marks = 63.5
	result = false

	fmt.Println(name)
	fmt.Println(marks)
	fmt.Println(result)
}
