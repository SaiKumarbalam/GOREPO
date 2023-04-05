package main

import "fmt"

func main() {
	x := 56.36
	y := 6.36

	result := avg(x, y)
	fmt.Println(result)

}
func avg(x float64, y float64) float64 {
	return (x + y) / 2
}
