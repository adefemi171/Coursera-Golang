package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type user struct {
	firstName string
	lastName  string
}

// readFile function for readin the file
func readFile() string {
	var fileName string
	fmt.Println("*** Enter absolute file path or place file in same directory of program execution if using just file name (example file.txt) *** ")
	fmt.Printf("\n" + "Please enter the name of the file: ")
	fmt.Scan(&fileName)
	return fileName
}

// openFile function for readin the file
func openFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return file
	// defer file.Close()
}

// checkLength function to know the Length of strings in the text file and
// set maximum size to 20
func checkLength(s string) string {
	maxSize := 20
	if len(s) > maxSize {
		return s[:maxSize]
	}
	return s
}

func readDetails(file *os.File) []user {
	var users []user
	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		details := strings.Fields(line)
		fname := checkLength(details[0])
		lname := checkLength(details[1])

		user := user{firstName: fname, lastName: lname}
		users = append(users, user)
	}

	return users
}

func print(users []user) {
	fmt.Printf("\n")
	for _, user := range users {
		fmt.Printf("First Name: %s Last Name: %s \n", user.firstName, user.lastName)
	}
}

func main() {
	fileName := readFile()

	file := openFile(fileName)
	defer file.Close()

	users := readDetails(file)

	print(users)
}
