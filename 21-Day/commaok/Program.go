package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	read := "please read all user input "
	fmt.Println(read)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating now")

	//comma ok || err
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ,", input)
	fmt.Printf("Type of rating %T ,", input)

}
