// Package comment
package main

import "fmt"

// function comment
func main() {
	var (
		planet   = "mars"
		distance = 54.6
	)

	fmt.Printf("We're %v million kilometers away from %v.\n", distance, planet)

	const weight = 149.0 * 0.3783
	fmt.Printf("My weight on the surface of %v is %v lbs.\n", planet, weight)

	const years = 41 * 365 / 687
	fmt.Printf("And I would be %v years old.\n", years)
}
