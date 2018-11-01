package main

import "fmt"

func main() {
	var (
		command = "on"
		thing   = "tv"
	)

	if command == "on" {
		fmt.Printf("It seems you want to turn on the %v", thing)
	} else if command == "off" {
		fmt.Printf("It seems you want to turn off the %v", thing)
	} else {
		fmt.Printf("Well, it seems you don't have any idea about what to do with the %v", thing)
	}
}
