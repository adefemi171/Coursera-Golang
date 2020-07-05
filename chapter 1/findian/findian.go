package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

// Main function
func main() {
	//Accepting users value
	fmt.Printf("Input a string of your choice: ")

	// Declared a viarable called user_input
	userInput := bufio.NewReader(os.Stdin)
	input, err := userInput.ReadString('\n')
	if err != nil {panic(err)}

	// val, err := fmt.Scan(&user_input)
	// if err != nil {panic(err)}

	// converting to lower case
	input = strings.ToLower(input)

	// trimmed white spaces
	input = strings.TrimSpace(input)

	// Using control flow to check if the string starts with i
	if strings.HasPrefix(input, "i"){
		// using another control to check if it contains a
		if strings.Contains(input, "a"){
			// using the last control flow to check if it ends with n
			if strings.HasSuffix(input, "n"){
				fmt.Println("Found!")
			} else { fmt.Println("Not Found!")}
			// returns not found for n

		}else { fmt.Println("Not Found!")}
		// returns not found for a

	}else { fmt.Println("Not Found!")}
	// returns not found for i

}

