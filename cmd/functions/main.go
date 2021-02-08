package main

import (
	"errors"
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
	withSurname := func(name string) string {
		return name + " Johnston"
	}
	fmt.Println(withSurname(author))
	// functions can also be called
	// on the spot
	func() {
		// in here we can use variables
		// from a higher scope
		// such as "author"
		age := 22
		fmt.Printf("author is %d\n", age)
	}()
	// cannot access "x" or "y" anymore
	// a multi-return value function
	name, err := createFullName("", "Smith")
	if err != nil {
		panic(err) // panic is a built-in function
	}
	fmt.Println(name)
}

func authorName() string {
	return "Alexander"
}

// takes in two strings
// returns a string and an potential error
func createFullName(firstName, lastName string) (string, error) {
	defer fmt.Println("Happens after createFullName returns. :)")
	if firstName == "" || lastName == "" {
		return "", errors.New("missing first or last name")
	}
	return firstName + " " + lastName, nil
}
