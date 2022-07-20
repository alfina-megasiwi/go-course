package main

import "fmt"

var x int
var y string
var z bool

func Exercise23() {
	fmt.Println("Exercise 2")
	fmt.Println("a. ")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(`b. These values are called the "zero value"`)

	// exercise-3
	x = 42
	y = "James Bond"
	z = true
	fmt.Println("Exercise 3")
	// a
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println("b.")
	fmt.Println(s)
}
