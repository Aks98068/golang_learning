package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type student struct {
	Student_name string
	Rollno       int
	Student_ID   int
	section      string
	Email        string
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	var s student

	fmt.Print("Enter student name: ")
	nameInput, _ := reader.ReadString('\n')
	s.Student_name = strings.TrimSpace(nameInput)

	fmt.Print("Enter roll number: ")
	rollnoInput, _ := reader.ReadString('\n')
	s.Rollno, _ = strconv.Atoi(strings.TrimSpace(rollnoInput))

	fmt.Print("Enter student ID: ")
	studentIDInput, _ := reader.ReadString('\n')
	s.Student_ID, _ = strconv.Atoi(strings.TrimSpace(studentIDInput))

	fmt.Print("Enter section: ")
	sectionInput, _ := reader.ReadString('\n')
	s.section = strings.TrimSpace(sectionInput)

	fmt.Print("Enter email: ")
	emailInput, _ := reader.ReadString('\n')
	s.Email = strings.TrimSpace(emailInput)

	data := fmt.Sprintf("Name: %s\nRoll no: %d\nStudent ID: %d\nSection: %s\nEmail: %s\n", s.Student_name, s.Rollno, s.Student_ID, s.section, s.Email)

	file, _ := os.Create(s.Student_name + ".txt")
	file.WriteString(data)
	file.Close()

}
