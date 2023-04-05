package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name     string
	RollNo   int
	Subjects []string
	School   string
	Gender   string
}

func main() {
	Students := map[string]Student{}
	var std Student

	for i := 0; i < 2; i++ {
		Students["std"+strconv.Itoa(std.RollNo)] = UserInput(std)
		std.RollNo += 1
	}
	fmt.Println(Students)

}
func UserInput(sr Student) Student {
	reader := bufio.NewReader(os.Stdin)
	//Name
	fmt.Println("Enter the Name: ")
	name, err := reader.ReadString('\n')
	//School
	fmt.Println("Enter the School: ")
	schValue, _ := reader.ReadString('\n')
	//Subjects

	fmt.Println("Enter the realsub: ")
	Subjects, err := reader.ReadSlice('\n')

	for i, v := range strings.Split(string(Subjects), " ") {
		sr.Subjects = append(sr.Subjects, string(v))
		sr.Subjects[i] = strings.TrimSpace(string(v))
	}

	fmt.Println("Enter the gender: ")
	Gender, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	sr.Name = strings.TrimSpace(name)
	sr.School = strings.TrimSpace(schValue)
	sr.Gender = strings.TrimSpace(Gender)
	return sr

}
func (s *Student) UpdateSchool(sch string) {
	s.School = sch
}
