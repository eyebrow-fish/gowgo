package main

import (
	"fmt"
	"strings"
)

// the entry point function: "main"
func main() {
	author := authorName()
	// "Println" and "ToLower" are functions
	// in "fmt" and "strings" packages
	fmt.Println(strings.ToLower(author))
	// functions can be assigned to variables
	appendBestLastName := func(name string) string {
		return name + " Johnston"
	}
	fmt.Printf("fullname: %s", appendBestLastName(author))
	// functions can also be called
	// on the spot
	func() {
		// in here we can use variables
		// from a higher scope
		// such as "author"
		age := 22
		fmt.Printf("author is %d", age)
	}()
	// like all blocks, functions are scoped
	// which means we cannot access "x" or "y"
}

func authorName() string {
	return "Alexander"
}
