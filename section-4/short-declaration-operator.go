package main

import "fmt"

func DeclarationDemo() {
	x := 42
	fmt.Println("From declaration demo:", x)
	x = 99
	fmt.Println("Assigned a new value:", x)

	y := 100 + 24
	fmt.Println(y)
}
