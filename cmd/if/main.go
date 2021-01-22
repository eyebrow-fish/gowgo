package main

import (
	"fmt"
)

func main() {
	x := true
	if x {
		fmt.Println("x is true")
	} else {
		fmt.Println("x is false")
	}
	// scoped variables
	if good := "dog"; good == "dog" {
		// good can be used here
		fmt.Printf("good %s\n", good)
	} else {
		// good can still be used
		fmt.Printf("not good %s\n", good)
	}
	// good can no longer be used
	if num := 50; num < 5 {
		fmt.Println("tiny")
		// scoped variables can be used
	} else if num < 20 {
		fmt.Println("medium")
	} else if num < 50 {
		fmt.Println("large")
	} else {
		fmt.Printf("%d is big\n", num)
	}
}
