package main

import "fmt"

//Global variables==> Which are declared Outside the main Function

//Local variables ==>Which are Declared in any function Inside the function

var a int = 60

func main() {
	Firstname := "test"
	Lastname := "me"
	fmt.Printf("%v", Firstname, Lastname)
	read()

}

func read() {
	fmt.Println(a)
}
