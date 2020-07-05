package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal Interface
type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

// Cow type
type Cow struct {
	name, food, locomotion, noise string
}

// Bird type
type Bird struct {
	name, food, locomotion, noise string
}

// Snake type
type Snake struct {
	name, food, locomotion, noise string
}

// Eat method Cow
func (c Cow) Eat() string {
	return c.food
}

// Move method Cow
func (c Cow) Move() string {
	return c.locomotion
}

// Speak Method Cow
func (c Cow) Speak() string {
	return c.noise
}

// Eat method Bird
func (b Bird) Eat() string {
	return b.food
}

// Move method Bird
func (b Bird) Move() string {
	return b.locomotion
}

// Speak Method Bird
func (b Bird) Speak() string {
	return b.noise
}

// Eat method Snake
func (s Snake) Eat() string {
	return s.food
}

// Move method Snake
func (s Snake) Move() string {
	return s.locomotion
}

// Speak Method Snake
func (s Snake) Speak() string {
	return s.noise
}

// Create animal type
func animalFactory(animalMap map[string]Animal, animalName string, animalType string) (Animal, bool) {
	var animalInterface Animal
	a, ok := animalMap[animalName]
	if !ok {
		switch animalType {
		case "cow":
			animalMap[animalName] = Cow{animalName, "grass", "walk", "moo"}
			return animalMap[animalName], true
		case "bird":
			animalMap[animalName] = Bird{animalType, "worms", "fly", "peep"}
			return animalMap[animalName], true
		case "snake":
			animalMap[animalName] = Snake{animalType, "mice", "slither", "hsss"}
			return animalMap[animalName], true
		default:
			return animalInterface, false
		}
	} else {
		return a, true
	}
}

// Get animal action
func getAnimalAction(animalType Animal, animalAction string) (string, bool) {
	switch animalAction {
	case "eat":
		return animalType.Eat(), true
	case "move":
		return animalType.Move(), true
	case "speak":
		return animalType.Speak(), true
	default:
		return "Invalid", false
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	animalMap := make(map[string]Animal)
	fmt.Println("1. To input a new animal use >newanimal AnimalName AnimalType example: >newanimal Dog Dog")
	fmt.Println("2. To Query an animal use >query AnimalName infoRequested example: >query cow move")
	fmt.Println("3. Enter x to exit")
	for {
		fmt.Print(">")
		if scanner.Scan() {
			userInput := strings.Trim(scanner.Text(), " ")
			seq := strings.Split(userInput, " ")
			if userInput == "x" || userInput == "X" {
				os.Exit(3)
			}
			if len(seq) == 3 {
				userCommand := strings.ToLower(seq[0])
				if userCommand == "newanimal" {
					animalName := strings.ToLower(seq[1])
					animalType := strings.ToLower(seq[2])
					a, ok := animalFactory(animalMap, animalName, animalType)
					if !ok {
						fmt.Println("Invalid Animal Type ->", a)
					} else {
						fmt.Println("Created it!")
					}
				} else if userCommand == "query" {
					animalName := strings.ToLower(seq[1])
					animalAction := strings.ToLower(seq[2])
					a, ok := animalMap[animalName]
					if !ok {
						fmt.Println("Animal Type not found -> ", animalName)
					} else {
						action, ok := getAnimalAction(a, animalAction)
						if !ok {
							fmt.Println("Animal action not found -> ", animalAction)
						} else {
							fmt.Println(action)
						}
					}
				} else {
					fmt.Println("Invalid command -> ", userCommand)
					continue
				}

			} else {
				fmt.Println("Invalid input -> ", userInput)
			}
		}
	}

}
