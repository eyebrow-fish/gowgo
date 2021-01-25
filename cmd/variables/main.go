package main

import (
	"fmt"
)

const allowDog bool = true

// supports multi-line
const (
	good string = "dog allowed"
	bad string = "not allowed"
)

func main() {
	dog := true // inferred
	var output string
	if dog && allowDog {
		output = good
	} else {
		output = bad
	}
	fmt.Println(output)
}
