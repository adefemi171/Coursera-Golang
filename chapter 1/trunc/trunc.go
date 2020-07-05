package main

import (
	"fmt"
)

// Main function
func main() {
	// Declared a viarable called input
	var input float64

	fmt.Printf("Please enter a decimal value: ")
	//Accepting users value
	fmt.Scan(&input)
	// returning users value into integer
	fmt.Printf("Your inputted value is been truncated to: %d \n", int(input))
}

// package main

// import "fmt"
// import "math"

// func main(){
// 	var x float64
// 	fmt.Println("Enter float number: ")
// 	fmt.Scan(&x)
// 	fmt.Println(math.Trunc(x))
// }
