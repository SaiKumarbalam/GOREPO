package main

import "fmt"

type Person struct {
	FirstName, LastName string
	Age                 int
	Weight              float32
}

func (e Person) displayDeatils() {
	fmt.Println("FirstName:", e.FirstName)
	fmt.Println("LastName:", e.LastName)
	fmt.Println("Age:", e.Age)
	fmt.Println("Weight:", e.Weight)
}

func main() {
	emp := Person{
		FirstName: "Ravi",
		LastName:  "Pinky",
		Age:       63,
		Weight:    56.36,
	}
	emp.displayDeatils()

}
