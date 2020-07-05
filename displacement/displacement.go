package main

import (
	"fmt"
	"math"
)

// GenDisplaceFn returns a function with time argument
func GenDisplaceFn(acc, initVelo, initDisplace float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*acc*(math.Pow(t,2)) + (initVelo*t) + initDisplace
	}
}

func main() {
	// Displacement Formula: s = 1/2 a t2 + v0t + s0
	var a, v0, s0, t float64

	// Get values from the User
	fmt.Print("Enter value for Acceleration: ")
	fmt.Scan(&a)
	fmt.Print("\n Enter value for Initial Velocity: ")
	fmt.Scan(&v0)
	fmt.Print("\n Enter value for Initial Displacement: ")
	fmt.Scan(&s0)

	// Getting a function for different Time
	fn := GenDisplaceFn(a, v0, s0)

	// Calling a displacement for different time period
	fmt.Print("\n Enter value for Time: ")
	fmt.Scan(&t)
	fmt.Println("\nDisplacement is, ", fn(t))

}