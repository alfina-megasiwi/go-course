package main

import "fmt"

type exercise4 int

var temp exercise4

func Exercise4() {
	fmt.Println("Exercise 4")
	fmt.Println("a.", temp)
	fmt.Println("b.")
	fmt.Printf("%T\n", temp)

	// c
	temp = 42
	fmt.Println("d.", temp)

}
