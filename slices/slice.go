package main

import (
	"fmt"
	"sort"
	"strconv"
)
func main(){
	empInt := make([]int, 3)
	// sortInt := make([]int, 3)
	var val string
	
	for true{
		fmt.Printf("Enter a number of any choice or press X to quit: ")
		fmt.Scanln(&val)
		if (val == "X" || val == "x"){
			break
		}
		uns, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Values inputted worngly, try again!")
			continue
		}

		empInt = append(empInt, uns)
		// sortInts := sort.Ints(empInt)
		sort.Ints(empInt)
		fmt.Println("Your values after sorting are: ", empInt)
}
}