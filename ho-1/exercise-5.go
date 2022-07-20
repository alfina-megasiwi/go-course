package main

import "fmt"

type exercise5 int

var tempx exercise5
var tempy int

func Exercise5() {
	fmt.Println("Exercise 5")

	// 1.
	fmt.Println(tempx)
	fmt.Printf("%T\n", tempx)
	tempx = 10
	fmt.Println(tempx)

	// 2.
	tempy = int(tempx)
	fmt.Println(tempy)
	fmt.Printf("%T\n", tempy)
}
