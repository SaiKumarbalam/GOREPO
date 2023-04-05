package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Id        int
}

func main() {
	p := Person{
		FirstName: "bala",
		LastName:  "danny",
		Id:        21,
	}
	fmt.Println(p.FirstName)
	fmt.Println(p.LastName)
	fmt.Println(p.Id)

}
