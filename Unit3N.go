// Author: William Provost
// Version: 1.0.0
// Date: 2025-11-17
// Fileoverview: This program uses nested if statements to determine a user's life stage based on age and student status.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// variable declaration
	var name string
	var ageString string
	var age int
	var studentString string
	var isStudent bool

	var reader = bufio.NewReader(os.Stdin)

	// get name from user
	fmt.Print("Enter your name: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// get age from user
	fmt.Print("Enter your age: ")
	ageString, _ = reader.ReadString('\n')
	ageString = strings.TrimSpace(ageString)
	age, _ = strconv.Atoi(ageString)

	// get student status
	fmt.Print("Are you a student (true/false): ")
	studentString, _ = reader.ReadString('\n')
	studentString = strings.TrimSpace(studentString)
	isStudent, _ = strconv.ParseBool(studentString)

	// nested if checks
	if isStudent {
		if age >= 13 && age <= 19 {
			fmt.Printf("Student %s is a teenager.\n", name)
		} else {
			if age >= 5 && age <= 12 {
				fmt.Printf("Student %s is a child.\n", name)
			} else {
				fmt.Printf("Student %s is in a different life stage.\n", name)
			}
		}
	} else {
		if age >= 20 && age <= 30 {
			fmt.Printf("%s is a young adult.\n", name)
		} else {
			fmt.Printf("%s is in a different life stage.\n", name)
		}
	}

	fmt.Println("\nDone.")
}
