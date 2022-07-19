package main

import "fmt"

var varA int = 42

type HotDog int // create a type hotdog, with the underline type int
var varB HotDog

func ConversionDemo() {
	varA = 42
	varB = 43
	fmt.Println(varA)
	fmt.Printf("%T\n", varA)
	fmt.Println(varB)
	fmt.Printf("%T\n", varB)

	varA = int(varB)
	fmt.Println(varA)
	fmt.Printf("%T\n", varA)
}
