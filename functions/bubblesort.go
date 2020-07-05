package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("Please enter ten integers, each number followed by Enter.")

	var arrNum []int
	var input string

	for {
		fmt.Print(": ")
		fmt.Scanln(&input)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("- Input must be an integer number")
			continue
		}
		arrNum = append(arrNum, num)
		if len(arrNum) == 10 {
			break
		}
	}
	fmt.Println("Values before getting Sorted", arrNum)
	BubbleSort(arrNum)
	fmt.Println("Values after getting sorted", arrNum)
}

// BubbleSort function
func BubbleSort(list []int) {
	end := len(list) - 1
	for {
		if end == 0 {
			break
		}
		for i := 0; i < len(list)-1; i++ {
			if list[i] > list[i+1] {
				Swap(list, i)
			}
		}
		end--
	}
}

// Swap function for bubble sort
func Swap(list []int, i int) {
	list[i], list[i+1] = list[i+1], list[i]
}