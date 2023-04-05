package main

import "fmt"

func main() {

	a := 25
	b := 30
	c := 65
	fmt.Println(a < b && b > c) //25 <30 =t //30 >60
	fmt.Println(a < b || b > c)
	fmt.Println(!(a == b && a > c))

	/*
		Truth Table AND
		true * true =true
		true * false =false
		false * false =false
		false * true =false

		Truth Table OR
		true || true =true
		true || flase =true
		flase || true =true
		false || false =false

		Truth Table NOT
		 !(TRUE)=FALSE
		 !(FALSE)=TRUE
	*/
}
