package main

import (
	"fmt"
	"os"	
)

type Animal struct {
	food, locomotion, noise string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	
	for {
		fmt.Println("Enter the name of the Animal(cow,snake,bird) and Info(move,eat,speak) you want or Enter x to exit: ")
		fmt.Print("> ")

		var animal, action string
		fmt.Scanln(&animal, &action)
		fmt.Println(animal, action)

		var animalObj Animal
		switch animal {
		case "cow":
			animalObj = Animal{"grass", "walk", "moo"}
		case "bird":
			animalObj = Animal{"worms", "fly", "peep"}
		case "snake":
			animalObj = Animal{"mice", "slither", "hsss"}
		case "x":
			os.Exit(3)
		default:
			fmt.Println("Wrong animal type")
		}

		switch action {
		case "eat":
			animalObj.Eat()
		case "move":
			animalObj.Move()
		case "speak":
			animalObj.Speak()
		case "x":
			os.Exit(3)
		default:
			fmt.Println("Wrong action type")
		}
	}
}
