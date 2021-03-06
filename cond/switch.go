package cond

import "fmt"

func ExecuteSwitch() {
	var (
		command = "off"
		thing   = "tv"
	)

	switch command {
	case "on":
		fmt.Printf("It seems you want to turn on the %v", thing)
	case "off":
		fmt.Printf("It seems you want to turn off the %v", thing)
	default:
		fmt.Printf("Well, it seems you don't have any idea about what to do with the %v", thing)
	}

	fmt.Print("\n")
}
