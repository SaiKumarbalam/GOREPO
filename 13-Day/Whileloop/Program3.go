package main

import "fmt"

func main() {
	colours := []string{"red", "yello", "green", "pink", "blue", "grey", "black"}
	fmt.Println("Colours to Examine :", colours)

	//Process all colours untill the slice is empty

	for len(colours) > 0 {
		if colours[0] == "blue" {
			fmt.Println("Found my Favourite colour,blue!")
		}
		colours = colours[1:]
	}
	fmt.Println("Colours left to process:", colours)
}
