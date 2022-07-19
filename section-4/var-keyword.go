package main

import "fmt"

// short declaration operator doesn't work outside the function body
// x:= 42

// declare the variable "y"
// assign the value 43
// declare & assign = initialization
var y = 43

// Declare there is a variable with the identifier "z"
// and that the variable with the identifier "z" is of type int
// assigns the zero value of type int to "z"
var z int

func VarKeyword() {
	// short declaration operator
	// Declare a variable and assign a value of a certain type

	// using the var keyword
	fmt.Println(y)

	foo123()
	fmt.Println(z)
}

func foo123() {
	fmt.Println("again: ", y)
}
