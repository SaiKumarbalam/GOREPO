package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Employee struct {
	Name      string
	ID        int
	Salary    float64
	Techstack []string
	Project   string
}

func main() {
	Employes := map[string]Employee{}

	var emp Employee

	for i := 0; i < 5; i++ {
		Employes["emp"+strconv.Itoa(emp.ID)] = UserInput(emp) //why emp.ID
		emp.ID += 1
	}
	fmt.Println(Employes)

}

func UserInput(em Employee) Employee {
	reader := bufio.NewReader(os.Stdin)
	//Name
	fmt.Println("Enter the Name: ")
	name, err := reader.ReadString('\n')
	//salary
	fmt.Println("Enter the salary: ")
	salValue, _ := reader.ReadString('\n')
	//technologies
	fmt.Println("Enter the technologies: ")
	techstack, err := reader.ReadSlice('\n')
	for i, v := range strings.Split(string(techstack), " ") {
		em.Techstack = append(em.Techstack, string(v))
		em.Techstack[i] = strings.TrimSpace(string(v))
	}

	fmt.Println("Enter the project: ")
	project, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	em.Name = strings.TrimSpace(name)                                  //why we use trim space
	em.Salary, _ = strconv.ParseFloat(strings.TrimSpace(salValue), 64) //why we use float
	em.Project = strings.TrimSpace(project)
	return em

}

// UpdateSalary Method to update the salary
func (e *Employee) UpdateSalary(sal float64) {
	e.Salary = sal
}

// UpdateProject Method to update the Project
func (e *Employee) UpdateProject(project string) {
	e.Project = project
}
