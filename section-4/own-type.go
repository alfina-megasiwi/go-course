package main

import "fmt"

var a int = 42

type hotdog int // create a type hotdog, with the underline type int
var b hotdog

func OwnType() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// There is a way for:
	// a = b
	// fmt.Println(a)
	// fmt.Printf("%T\n", a)
	// in the next video
}
