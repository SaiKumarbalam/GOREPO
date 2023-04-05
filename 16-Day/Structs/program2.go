package main

import "fmt"

type Car struct {
	Name, Model string
	Color       string
	MadeYear    int
	Weight      float64
}

func main() {

	// Declarning variable of a struc type
	// var c Car
	// fmt.Println(c)

	c1 := Car{
		Name:     "telsa",
		Model:    "X",
		Color:    "white",
		MadeYear: 2023,
		Weight:   63.23,
	}
	// accessing struct
	fmt.Println("Car", c1)
	fmt.Println("Car -", c1.Name)
	fmt.Println("Car MadeYear -", c1.MadeYear)
	// fmt.Println("Car Color -", c1.Color)

	// Assing new value to existing struct
	c1.Color = "Black"
	fmt.Println("Car Color -", c1.Color)

}
