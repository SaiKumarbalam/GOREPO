package main

import "fmt"

func main() {
	name("Sandy", "randy")
	// firstname = "Sandy"
	// lastname = "randy"

}

func name(firstname, lastname string) {

	fmt.Println("My name is :", firstname+lastname)
}
