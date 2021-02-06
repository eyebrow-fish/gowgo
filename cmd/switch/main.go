package main

import (
	"fmt"
)

func main() {
	switch name := "Alex"; name {
	case "Alex":
		fmt.Println("is author")
	case "Larissa", "Markus", "Annette":
		fmt.Println("is awesome")
	default:
		fmt.Println("who are you?")
	}
	// sometimes we want to evaluate conditions
	x := 5
	switch {
	case x < 5:
		fmt.Printf("%d is less than 5\n", x)
	default:
		fmt.Printf("%d is too big\n", x)
	}
	// fallthrough falls through to the next
	// case even if it's not met
	switch good := "dog"; good {
	case "dog":
		fmt.Println("good boy!")
		fallthrough
	default:
		// we will always reach this point
		fmt.Printf("got %s\n", good)
	}
}
